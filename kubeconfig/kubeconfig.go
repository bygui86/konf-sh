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

func Split(kubeConfig *clientcmdapi.Config) map[string]*clientcmdapi.Config {
	singleConfigs := make(map[string]*clientcmdapi.Config, len(kubeConfig.Contexts))
	for ctxKey, ctxValue := range kubeConfig.Contexts {
		contexts := make(map[string]*clientcmdapi.Context, 1)
		contexts[ctxKey] = ctxValue
		clusters := make(map[string]*clientcmdapi.Cluster, 1)
		clusters[ctxValue.Cluster] = kubeConfig.Clusters[ctxValue.Cluster]
		authInfos := make(map[string]*clientcmdapi.AuthInfo, 1)
		authInfos[ctxValue.AuthInfo] = kubeConfig.AuthInfos[ctxValue.AuthInfo]
		singleConfigs[ctxKey] = &clientcmdapi.Config{
			APIVersion:     kubeConfig.APIVersion,
			Kind:           kubeConfig.Kind,
			CurrentContext: ctxKey,
			Contexts:       contexts,
			Clusters:       clusters,
			AuthInfos:      authInfos,
		}
	}
	return singleConfigs
}

func Save(singleConfigs map[string]*clientcmdapi.Config, singleConfigsPath string) error {
	logger.SugaredLogger.Debugf("🐛 Check single configs folder '%s'", singleConfigsPath)
	checkErr := CheckKubeConfigsFolder(singleConfigsPath)
	if checkErr != nil {
		return checkErr
	}

	for cfgKey, cfg := range singleConfigs {
		logger.SugaredLogger.Debugf("🐛 Validate config '%s'", cfgKey)
		valErr := validate(cfg)
		if valErr != nil {
			return valErr
		}

		logger.SugaredLogger.Debugf("🐛 Write config to file '%s'", cfgKey)
		writeErr := write(cfg, filepath.Join(singleConfigsPath, cfgKey))
		if writeErr != nil {
			return writeErr
		}
	}

	return nil
}

func validate(kubeConfig *clientcmdapi.Config) error {
	err := clientcmd.Validate(*kubeConfig)
	if clientcmd.IsConfigurationInvalid(err) {
		return err
	}
	return nil
}

func write(kubeConfig *clientcmdapi.Config, filepath string) error {
	return clientcmd.WriteToFile(*kubeConfig, filepath)
}
