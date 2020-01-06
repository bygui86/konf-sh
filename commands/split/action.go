package split

import (
	"fmt"
	"path/filepath"

	"bygui86/kubeconfigurator/commons"
	"bygui86/kubeconfigurator/kubeconfig"
	"bygui86/kubeconfigurator/logger"
	"bygui86/kubeconfigurator/utils"

	"github.com/urfave/cli"
)

// TODO implement flags/env-vars usage
func split(ctx *cli.Context) error {
	logger.Logger.Info("")

	logger.Logger.Info("📖 Load Kubernetes configuration")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Debugf("🐛 Kubernetes configuration file path: %s", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

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
		valErr := kubeconfig.Validate(cfg)
		if valErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error validating single Kubernetes configuration '%s': %s", cfgKey, valErr.Error()),
				12)
		}

		cfgFilePath := filepath.Join(singleConfigsPath, cfgKey)
		logger.SugaredLogger.Debugf("🐛 Write Kubernetes configuration '%s' to file '%s'", cfgKey, cfgFilePath)
		writeErr := kubeconfig.Write(cfg, cfgFilePath)
		if writeErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error writing single Kubernetes configuration '%s' to file: %s", cfgKey, writeErr.Error()),
				13)
		}
	}

	logger.SugaredLogger.Infof("✅ Completed! Single configs files saved in '%s'", singleConfigsPath)
	logger.Logger.Info("")
	return nil
}
