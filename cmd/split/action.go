package split

import (
	"fmt"
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func split(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing SPLIT command")

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

	zap.S().Infof("🪚 Split Kubernetes configuration from %s", kCfgFilePath)
	singleKcfs := kubeconfig.Split(kCfg, kCfgFilePath)

	zap.L().Info("💾 Save single Kubernetes konfigurations")
	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKfgsPath := ctx.String(commons.SingleKonfigsFlagName)
	zap.S().Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKfgsPath)

	zap.S().Debugf("🐛 Check existence of single Kubernetes konfigurations path '%s'", singleKfgsPath)
	checkErr := commons.CheckIfFolderExist(singleKfgsPath, true)
	if checkErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes konfigurations path '%s': %s",
				singleKfgsPath, checkErr.Error()), 11)
	}

	validKfgs := make([]string, 0)
	invalidKfgs := make([]string, 0)
	for kfgName, kfg := range singleKcfs {
		kfgFilePath := filepath.Join(singleKfgsPath, kfgName)

		newValErr := kubeconfig.Validate(kfg)
		if newValErr != nil {
			zap.S().Errorf("❌  Error validating Kubernetes konfiguration '%s': %s - skipping",
				kfgName, newValErr.Error())
			invalidKfgs = append(invalidKfgs, kfgName)
			continue
		}

		newWriteErr := kubeconfig.Write(kfg, kfgFilePath)
		if newWriteErr != nil {
			zap.S().Errorf("❌  Error writing Kubernetes konfiguration '%s' to file '%s': %s",
				kfgName, kfgFilePath, newValErr.Error())
			invalidKfgs = append(invalidKfgs, kfgName)
			continue
		}

		validKfgs = append(validKfgs, kfgName)
	}

	if len(validKfgs) > 0 {
		zap.S().Infof("📚 Available Kubernetes konfigurations in '%s':", singleKfgsPath)
		for _, v := range validKfgs {
			zap.S().Infof("\t%s", v)
		}
	}

	if len(invalidKfgs) > 0 {
		zap.L().Info("")
		zap.S().Infof("❓️ Invalid context found in Kubernetes configuration from '%s':", kCfgFilePath)
		for _, iv := range invalidKfgs {
			zap.S().Infof("\t%s", iv)
		}
	}

	if len(validKfgs) > 0 {
		zap.S().Infof("✅  Single Kubernetes konfigurations saved to '%s'", singleKfgsPath)
	} else {
		zap.S().Infof("❌  Split Kubernetes configurations from '%s' failed: no valid context found",
			kCfgFilePath)
	}
	zap.L().Info("")
	return nil
}
