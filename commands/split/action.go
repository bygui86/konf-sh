package split

import (
	"fmt"
	"github.com/urfave/cli"
	"path/filepath"

	"bygui86/kubeconfigurator/kubeconfig"
	"bygui86/kubeconfigurator/logger"
	"bygui86/kubeconfigurator/utils"
)

const (
	kubeConfigOutputFolderName = ".kube/configs"
)

func split() func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		logger.Logger.Info("")
		logger.Logger.Info("🏠 Get HOME")
		home, homeErr := utils.GetHomeDir()
		if homeErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error getting HOME environment variable: %s", homeErr.Error()),
				10)
		}

		logger.Logger.Info("📚 Get Kubernetes configuration file")
		kubeConfigFilePath, getErr := kubeconfig.GetKubeConfigFilePath(home)
		if getErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error getting Kubernetes configuration file: %s", getErr.Error()),
				20)
		}
		kubeConfig := kubeconfig.Get(kubeConfigFilePath)

		logger.Logger.Info("✂️  Split Kubernetes configuration file")
		singleConfigs := kubeconfig.Split(kubeConfig)

		logger.Logger.Info("💾 Save single Kubernetes configuration files")
		singleConfigsPath := filepath.Join(home, kubeConfigOutputFolderName)
		saveErr := kubeconfig.Save(singleConfigs, singleConfigsPath)
		if saveErr != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error saving Kubernetes configuration files: %s - Target path: %s", saveErr.Error(), singleConfigsPath),
				30)
		}

		logger.SugaredLogger.Infof("✅ Completed! Single configs files saved in '%s'", singleConfigsPath)
		logger.Logger.Info("")
		return nil
	}
}
