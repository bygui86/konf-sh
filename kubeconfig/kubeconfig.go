package kubeconfig

import (
	"errors"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/bygui86/konf/logger"
)

func Load(kubeConfigFilePath string) *clientcmdapi.Config {
	logger.SugaredLogger.Debugf("🐛 Load Kubernetes configuration from file '%s'", kubeConfigFilePath)
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func Split(kubeConfig *clientcmdapi.Config, kubeConfigFilePath string) map[string]*clientcmdapi.Config {
	logger.SugaredLogger.Debugf("🐛 Split Kubernetes configuration from %s", kubeConfigFilePath)
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
	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration")
	err := clientcmd.Validate(*kubeConfig)
	if clientcmd.IsConfigurationInvalid(err) {
		return err
	}
	return nil
}

func Write(kubeConfig *clientcmdapi.Config, filepath string) error {
	logger.SugaredLogger.Debugf("🐛 Write Kubernetes configuration to file '%s'", filepath)
	return clientcmd.WriteToFile(*kubeConfig, filepath)
}

func CheckIfContextExist(kubeConfig *clientcmdapi.Config, context string) error {
	logger.SugaredLogger.Debugf("🐛 Check context '%s' existence in Kubernetes configuration '%s'", context, kubeConfig.CurrentContext)
	ctxValue := kubeConfig.Contexts[context]
	if ctxValue == nil {
		logger.SugaredLogger.Debugf("🐛 Context '%s' not found", context)
		return errors.New("context not found")
	}
	logger.SugaredLogger.Debugf("🐛 Context '%s' is valid", context)
	return nil
}
