package list

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func list(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing LIST command")

	cfgsErr := listKubeCfgs(ctx)
	if cfgsErr != nil {
		return cfgsErr
	}

	kfgsErr := listSingleKfgs(ctx)
	if kfgsErr != nil {
		return kfgsErr
	}

	zap.L().Info("")
	return nil
}

func listKubeCfgs(ctx *cli.Context) error {
	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kCfg := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Debugf("🐛 Validate Kubernetes configuration from '%s'", kCfgFilePath)
	valErr := kubeconfig.Validate(kCfg)
	if valErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s",
				kCfgFilePath, valErr.Error()), 12)
	}

	zap.S().Debugf("🐛 List contexts in Kubernetes configuration '%s'", kCfgFilePath)
	ctxs := kubeconfig.ListContexts(kCfg)
	if len(ctxs) > 0 {
		zap.S().Infof("📚 Available Kubernetes contexts in '%s':", kCfgFilePath)
		for _, v := range ctxs {
			zap.S().Infof("\t%s", v)
		}
	} else {
		zap.S().Warnf("🚨 No available Kubernetes context in '%s'", kCfgFilePath)
	}

	zap.L().Info("")
	return nil
}

func listSingleKfgs(ctx *cli.Context) error {
	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKfgsPath := ctx.String(commons.SingleKonfigsFlagName)
	zap.S().Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKfgsPath)

	checkErr := commons.CheckIfFolderExist(singleKfgsPath, false)
	if checkErr != nil {
		zap.S().Warnf("🚨 Single Kubernetes konfigurations path not found ('%s')", singleKfgsPath)
		zap.L().Warn("💬 Tip: run 'konf split' before anything else")
	} else {
		validCfgs := make([]string, 0)
		invalidCfgs := make([]string, 0)
		err := filepath.Walk(
			singleKfgsPath,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}

				if strings.HasPrefix(info.Name(), ".") {
					return nil
				}

				kCfg := kubeconfig.Load(path)
				// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file
				if kubeconfig.Validate(kCfg) != nil {
					invalidCfgs = append(invalidCfgs, info.Name())
					return nil
				}

				validCfgs = append(validCfgs, info.Name())

				return nil
			},
		)
		if err != nil {
			return cli.Exit(
				fmt.Sprintf("❌  Error listing single Kubernetes konfigurations in '%s': %s", singleKfgsPath, err.Error()),
				21)
		}

		if len(validCfgs) > 0 {
			zap.S().Infof("📚 Available Kubernetes konfigurations in '%s':", singleKfgsPath)
			for _, v := range validCfgs {
				zap.S().Infof("\t%s", v)
			}
		} else {
			zap.S().Warnf("🚨 No available Kubernetes konfigurations in '%s'", singleKfgsPath)
			zap.L().Warn("💬 Tip: run 'konf split' before anything else")
		}

		if len(invalidCfgs) > 0 {
			zap.L().Info("")
			zap.S().Infof("❓ Invalid Kubernetes konfigurations in '%s':", singleKfgsPath)
			for _, iv := range invalidCfgs {
				zap.S().Infof("\t%s", iv)
			}
		}
	}

	return nil
}
