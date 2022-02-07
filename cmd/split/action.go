package split

import (
	"fmt"
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/tools/clientcmd/api"
)

func split(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SPLIT-CONFIG command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kubeConfigFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error validating Kubernetes configuration from '%s': %s", kubeConfigFilePath, valErr.Error()),
			12)
	}

	logger.SugaredLogger.Infof("✂️  Split Kubernetes configuration from %s", kubeConfigFilePath)
	singleConfigs := kubeconfig.Split(kubeConfig, kubeConfigFilePath)

	logger.Logger.Info("💾 Save single Kubernetes configurations files")
	logger.Logger.Debug("🐛 Get single Kubernetes configurations files path")
	singleConfigsPath := ctx.String(commons.SingleConfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes configurations files path: '%s'", singleConfigsPath)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations files path '%s'", singleConfigsPath)
	checkErr := utils.CheckIfFolderExist(singleConfigsPath, true)
	if checkErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error checking existence of Kubernetes configurations files path '%s': %s", checkErr.Error(), singleConfigsPath),
			11)
	}

	valWrErr := validateAndWrite(singleConfigs, singleConfigsPath)
	if valWrErr != nil {
		return valWrErr
	}

	logger.SugaredLogger.Infof("✅ Completed! Single Kubernetes configurations files saved in '%s'", singleConfigsPath)
	logger.Logger.Info("")
	return nil
}

func validateAndWrite(singleConfigs map[string]*api.Config, singleConfigsPath string) error {
	// TODO implement a mechanism to avoid complete fail if just 1 out of N configurations is not valid
	for cfgKey, cfg := range singleConfigs {
		cfgFilePath := filepath.Join(singleConfigsPath, cfgKey)
		valWrErr := commons.ValidateAndWrite(cfg, cfgFilePath)
		if valWrErr != nil {
			return valWrErr
		}
	}

	return nil
}
