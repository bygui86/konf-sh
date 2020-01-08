package split

import (
	"fmt"
	"path/filepath"

	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/logger"
	"bygui86/konf/utils"

	"github.com/urfave/cli"
)

func split(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SPLIT command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kubeConfigFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error validating Kubernetes configuration from '%s': %s", kubeConfigFilePath, valErr.Error()),
			12)
	}

	logger.Logger.Info("✂️  Split Kubernetes configuration")
	singleConfigs := kubeconfig.Split(kubeConfig)

	logger.Logger.Info("💾 Save single Kubernetes configurations files")
	logger.Logger.Debug("🐛 Get single Kubernetes configurations files path")
	singleConfigsPath := ctx.String(commons.SingleConfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes configurations files path: '%s'", singleConfigsPath)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations files path '%s'", singleConfigsPath)
	checkErr := utils.CheckIfFolderExist(singleConfigsPath, true)
	if checkErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of Kubernetes configurations files path '%s': %s", checkErr.Error(), singleConfigsPath),
			11)
	}

	// TODO implement a mechanism to avoid complete fail if just 1 out of N configurations is not valid
	for cfgKey, cfg := range singleConfigs {
		logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration '%s'", cfgKey)
		cfgValErr := kubeconfig.Validate(cfg)
		if cfgValErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error validating single Kubernetes configuration '%s': %s", cfgKey, cfgValErr.Error()),
				12)
		}

		cfgFilePath := filepath.Join(singleConfigsPath, cfgKey)
		logger.SugaredLogger.Debugf("🐛 Write Kubernetes configuration '%s' to file '%s'", cfgKey, cfgFilePath)
		cfgWriteErr := kubeconfig.Write(cfg, cfgFilePath)
		if cfgWriteErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error writing single Kubernetes configuration '%s' to file: %s", cfgKey, cfgWriteErr.Error()),
				13)
		}
	}

	logger.SugaredLogger.Infof("✅ Completed! Single Kubernetes configurations files saved in '%s'", singleConfigsPath)
	logger.Logger.Info("")
	return nil
}
