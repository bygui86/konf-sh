package split

import (
	"fmt"
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func split(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SPLIT command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kCfg := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kCfgFilePath)
	valErr := kubeconfig.Validate(kCfg)
	if valErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s",
				kCfgFilePath, valErr.Error()), 12)
	}

	logger.SugaredLogger.Infof("🪚 Split Kubernetes configuration from %s", kCfgFilePath)
	singleKcfs := kubeconfig.Split(kCfg, kCfgFilePath)

	logger.Logger.Info("💾 Save single Kubernetes konfigurations")
	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations path")
	singleKfgsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKfgsPath)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes konfigurations path '%s'", singleKfgsPath)
	checkErr := utils.CheckIfFolderExist(singleKfgsPath, true)
	if checkErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes konfigurations path '%s': %s",
				checkErr.Error(), singleKfgsPath), 11)
	}

	validKfgs := make([]string, 0)
	invalidKfgs := make([]string, 0)
	for kfgName, kfg := range singleKcfs {
		kfgFilePath := filepath.Join(singleKfgsPath, kfgName)

		newValErr := kubeconfig.Validate(kfg)
		if newValErr != nil {
			logger.SugaredLogger.Errorf("❌  Error validating Kubernetes konfiguration '%s': %s - skipping",
				kfgName, newValErr.Error())
			invalidKfgs = append(invalidKfgs, kfgName)
			continue
		}

		newWriteErr := kubeconfig.Write(kfg, kfgFilePath)
		if newWriteErr != nil {
			logger.SugaredLogger.Errorf("❌  Error writing Kubernetes konfiguration '%s' to file '%s': %s",
				kfgName, kfgFilePath, newValErr.Error())
			invalidKfgs = append(invalidKfgs, kfgName)
			continue
		}

		validKfgs = append(validKfgs, kfgName)
	}

	if len(validKfgs) > 0 {
		logger.SugaredLogger.Infof("📚 Available Kubernetes konfigurations in '%s':", singleKfgsPath)
		for _, v := range validKfgs {
			logger.SugaredLogger.Infof("\t%s", v)
		}
	}

	if len(invalidKfgs) > 0 {
		logger.Logger.Info("")
		logger.SugaredLogger.Infof("❓️ Invalid context found in Kubernetes configuration from '%s':", kCfgFilePath)
		for _, iv := range invalidKfgs {
			logger.SugaredLogger.Infof("\t%s", iv)
		}
	}

	if len(validKfgs) > 0 {
		logger.SugaredLogger.Infof("✅  Single Kubernetes konfigurations saved to '%s'", singleKfgsPath)
	} else {
		logger.SugaredLogger.Infof("❌  Split Kubernetes configurations from '%s' failed: no valid context found",
			kCfgFilePath)
	}
	logger.Logger.Info("")
	return nil
}
