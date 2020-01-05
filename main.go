package main

import (
	"k8s.io/client-go/tools/clientcmd/api"
	"os"
	"path/filepath"

	"bygui86/kubeconfigurator/kubeconfig"
	"bygui86/kubeconfigurator/logger"
	"bygui86/kubeconfigurator/utils"
)

const (
	kubeConfigOutputFolderName = ".kube/configs"
)

func main() {
	logger.Logger.Info("🏠 Get HOME")
	home := getHome()

	logger.Logger.Info("📚 Get Kubernetes configuration file")
	srcConfig := getSrcConfig(home)

	logger.Logger.Info("✂️  Split Kubernetes configuration file")
	tgConfigs := splitConfig(srcConfig)

	logger.Logger.Info("💾 Save single Kubernetes configuration files")
	configsPath := saveConfigs(home, tgConfigs)

	logger.SugaredLogger.Infof("✅ Completed, configs path: '%s'", configsPath)
}

func getHome() string {
	home, homeErr := utils.GetHomeDir()
	if homeErr != nil {
		logger.SugaredLogger.Errorf("❌ Error getting HOME environment variable: %s", homeErr.Error())
		os.Exit(1)
	}
	return home
}

func getSrcConfig(home string) *api.Config {
	kubeConfigFilePath, getErr := kubeconfig.GetKubeConfig(home)
	if getErr != nil {
		logger.SugaredLogger.Errorf("❌ Error getting Kubernetes configuration file: %s", getErr.Error())
		os.Exit(1)
	}
	return kubeconfig.Get(kubeConfigFilePath)
}

func splitConfig(srcConfig *api.Config) map[string]*api.Config {
	return kubeconfig.Split(srcConfig)
}

func saveConfigs(home string, tgConfigs map[string]*api.Config) string {
	tgPath := filepath.Join(home, kubeConfigOutputFolderName)
	configsPath, saveErr := kubeconfig.Save(tgConfigs, tgPath)
	if saveErr != nil {
		logger.SugaredLogger.Errorf("❌ Error saving Kubernetes configuration files: %s - Target path: %s", saveErr.Error(), tgPath)
		os.Exit(1)
	}
	return configsPath
}
