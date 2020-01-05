package main

import (
	"flag"
	"go.uber.org/zap"
	"os"
	"path/filepath"

	"bygui86/kubeconfigurator/logger"
	"bygui86/kubeconfigurator/utils"

	clientcmd "k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	kubeConfigFlagKey = "kubeconfig"

	kubeConfigFolderDefault = ".kube"
	kubeConfigFilenameDefault = "config"

	kubeConfigOutputFolderName = ".kube/configs"
)

func main() {
	logger.Logger.Info("🏠 Get HOME")
	home, homeErr := utils.GetHomeDir()
	if homeErr != nil {
		logger.Logger.Error("❌ error getting HOME environment variable",
			zap.String("error", homeErr.Error()))
		os.Exit(1)
	}

	logger.Logger.Info("📚 Get Kubernetes configuration file")
	srcPath := filepath.Join(home, kubeConfigFolderDefault, kubeConfigFilenameDefault)
	kubeConfigFilePath, getErr := getKubeConfigFlagValue(kubeConfigFlagKey, srcPath)
	if getErr != nil {
		logger.Logger.Error("❌ error getting Kubernetes configuration file",
			zap.String("source path", srcPath),
			zap.String("error", getErr.Error()))
		os.Exit(1)
	}
	srcConfig := getKubeConfig(kubeConfigFilePath)

	logger.Logger.Info("✂️ Split Kubernetes configuration file")
	tgConfigs := splitKubeConfig(srcConfig)

	logger.Logger.Info("💾 Save single Kubernetes configuration files")
	tgPath := filepath.Join(home, kubeConfigOutputFolderName)
	configsPath, saveErr := saveKubeConfigs(tgConfigs, tgPath)
	if saveErr != nil {
		logger.Logger.Error("❌ error saving Kubernetes configuration files",
			zap.String("target path", tgPath),
			zap.String("error", saveErr.Error()))
		os.Exit(1)
	}
	logger.Logger.Info("✅ Completed",
		zap.String("configs path", configsPath))
}

func getKubeConfigFlagValue(flagKey, flagDefault string) (string,error){
	return *flag.String(flagKey, flagDefault, "(optional) absolute path to the kubeconfig file"), nil
}

func getKubeConfig(kubeConfigFilePath string) *clientcmdapi.Config {
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func splitKubeConfig(srcConfig *clientcmdapi.Config) map[string]*clientcmdapi.Config {
	tgConfigs := make(map[string]*clientcmdapi.Config, len(srcConfig.Contexts))
	for ctxKey, ctxValue := range srcConfig.Contexts {
		contexts := make(map[string]*clientcmdapi.Context, 1)
		contexts[ctxKey] = ctxValue
		clusters := make(map[string]*clientcmdapi.Cluster, 1)
		clusters[ctxValue.Cluster] = srcConfig.Clusters[ctxValue.Cluster]
		authInfos := make(map[string]*clientcmdapi.AuthInfo, 1)
		authInfos[ctxValue.AuthInfo] = srcConfig.AuthInfos[ctxValue.AuthInfo]
		tgConfigs[ctxKey] = &clientcmdapi.Config{
			APIVersion: srcConfig.APIVersion,
			Kind: srcConfig.Kind,
			CurrentContext: ctxKey,
			Contexts: contexts,
			Clusters: clusters,
			AuthInfos: authInfos,
		}
	}
	return tgConfigs
}

func saveKubeConfigs(configs map[string]*clientcmdapi.Config, configsPath string) (string, error) {
	logger.Logger.Debug("Check configs folder",
		zap.String("configs path", configsPath))
	checkErr := utils.CheckKubeConfigsFolder(configsPath)
	if checkErr != nil {
		return "", checkErr
	}

	for cfgKey, cfg := range configs {
		logger.Logger.Debug("Validate config",
			zap.String("config key", cfgKey))
		valErr := validateKubeConfig(cfg)
		if valErr != nil {
			return "", valErr
		}
		logger.Logger.Debug("Write config to file",
			zap.String("config key", cfgKey))
		writeErr := writeKubeConfig(cfg, filepath.Join(configsPath, cfgKey))
		if writeErr != nil {
			return "", writeErr
		}
	}

	return configsPath, nil
}

func validateKubeConfig(config *clientcmdapi.Config) error {
	validateErr := clientcmd.Validate(*config)
	if clientcmd.IsConfigurationInvalid(validateErr) {
		return validateErr
	}
	return nil
}

func writeKubeConfig(config *clientcmdapi.Config, filepath string) error {
	return clientcmd.WriteToFile(*config, filepath)
}
