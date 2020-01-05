package kubeconfig

import (
	"path/filepath"

	"bygui86/kubeconfigurator/logger"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func Get(kubeConfigFilePath string) *clientcmdapi.Config {
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func Split(srcConfig *clientcmdapi.Config) map[string]*clientcmdapi.Config {
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

func Save(configs map[string]*clientcmdapi.Config, configsPath string) (string, error) {
	logger.SugaredLogger.Debugf("Check configs folder '%s'", configsPath)
	checkErr := CheckKubeConfigsFolder(configsPath)
	if checkErr != nil {
		return "", checkErr
	}

	for cfgKey, cfg := range configs {
		logger.SugaredLogger.Debugf("Validate config '%s'", cfgKey)
		valErr := validate(cfg)
		if valErr != nil {
			return "", valErr
		}

		logger.SugaredLogger.Debugf("Write config to file '%s'", cfgKey)
		writeErr := write(cfg, filepath.Join(configsPath, cfgKey))
		if writeErr != nil {
			return "", writeErr
		}
	}

	return configsPath, nil
}

func validate(config *clientcmdapi.Config) error {
	validateErr := clientcmd.Validate(*config)
	if clientcmd.IsConfigurationInvalid(validateErr) {
		return validateErr
	}
	return nil
}

func write(config *clientcmdapi.Config, filepath string) error {
	return clientcmd.WriteToFile(*config, filepath)
}
