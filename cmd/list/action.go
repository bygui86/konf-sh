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

	zap.L().Info("")
	return nil
}
