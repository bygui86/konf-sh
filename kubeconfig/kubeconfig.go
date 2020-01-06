package kubeconfig

import (
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"bygui86/kubeconfigurator/logger"
)

func Load(kubeConfigFilePath string) *clientcmdapi.Config {
	logger.SugaredLogger.Debugf("🐛 Load Kubernetes configuration from file '%s'", kubeConfigFilePath)
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func Split(kubeConfig *clientcmdapi.Config) map[string]*clientcmdapi.Config {
	logger.Logger.Debug("🐛 Split Kubernetes configuration")
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

func Validate(kubeConfig *clientcmdapi.Config) error {
	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration '%s'", kubeConfig.CurrentContext)
	err := clientcmd.Validate(*kubeConfig)
	if clientcmd.IsConfigurationInvalid(err) {
		return err
	}
	return nil
}

func Write(kubeConfig *clientcmdapi.Config, filepath string) error {
	logger.SugaredLogger.Debugf("🐛 Write Kubernetes configuration '%s' to file '%s'", kubeConfig.CurrentContext, filepath)
	return clientcmd.WriteToFile(*kubeConfig, filepath)
}
