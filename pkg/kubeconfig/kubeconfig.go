package kubeconfig

import (
	"errors"
	"strings"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func Load(kubeConfigFilePath string) *clientcmdapi.Config {
	zap.S().Debugf("🐛 Load Kubernetes configuration from file '%s'", kubeConfigFilePath)
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func Split(kubeConfig *clientcmdapi.Config, kubeConfigFilePath string) map[string]*clientcmdapi.Config {
	zap.S().Debugf("🐛 Split Kubernetes configuration from %s", kubeConfigFilePath)
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
	zap.S().Debugf("🐛 Validate Kubernetes configuration")
	err := clientcmd.Validate(*kubeConfig)
	if clientcmd.IsConfigurationInvalid(err) {
		return err
	}
	return nil
}

func Write(kubeConfig *clientcmdapi.Config, filepath string) error {
	zap.S().Debugf("🐛 Write Kubernetes configuration to file '%s'", filepath)
	return clientcmd.WriteToFile(*kubeConfig, filepath)
}

func CheckIfContextExist(kubeConfig *clientcmdapi.Config, context string) error {
	zap.S().Debugf("🐛 Check context '%s' existence in Kubernetes configuration '%s'", context, kubeConfig.CurrentContext)
	ctxValue := kubeConfig.Contexts[context]
	if ctxValue == nil {
		zap.S().Debugf("🐛 Context '%s' not found", context)
		return errors.New("context not found")
	}
	zap.S().Debugf("🐛 Context '%s' is valid", context)
	return nil
}

func CheckIfClusterExist(kubeConfig *clientcmdapi.Config, cluster string) error {
	zap.S().Debugf("🐛 Check cluster '%s' existence in Kubernetes configuration '%s'", cluster, kubeConfig.CurrentContext)
	clValue := kubeConfig.Clusters[cluster]
	if clValue == nil {
		zap.S().Debugf("🐛 Cluster '%s' not found", cluster)
		return errors.New("cluster not found")
	}
	zap.S().Debugf("🐛 Cluster '%s' is valid", cluster)
	return nil
}

func CheckIfAuthInfoExist(kubeConfig *clientcmdapi.Config, authInfo string) error {
	zap.S().Debugf("🐛 Check user '%s' existence in Kubernetes configuration '%s'", authInfo, kubeConfig.CurrentContext)
	authValue := kubeConfig.AuthInfos[authInfo]
	if authValue == nil {
		zap.S().Debugf("🐛 User '%s' not found", authInfo)
		return errors.New("user not found")
	}
	zap.S().Debugf("🐛 User '%s' is valid", authInfo)
	return nil
}

func RemoveContext(ctxMap map[string]*clientcmdapi.Context, context string) (map[string]*clientcmdapi.Context, error) {
	ctxValue := ctxMap[context]
	if ctxValue == nil {
		zap.S().Debugf("🐛 Context '%s' not found", context)
		return nil, errors.New("context not found")
	}

	newCtxMap := make(map[string]*clientcmdapi.Context)
	for ctxK, ctxV := range ctxMap {
		if ctxK != context {
			newCtxMap[ctxK] = ctxV
		}
	}
	return newCtxMap, nil
}

func RemoveCluster(clMap map[string]*clientcmdapi.Cluster, cluster string) (map[string]*clientcmdapi.Cluster, error) {
	clValue := clMap[cluster]
	if clValue == nil {
		zap.S().Debugf("🐛 Cluster '%s' not found", cluster)
		return nil, errors.New("cluster not found")
	}

	newClMap := make(map[string]*clientcmdapi.Cluster)
	for clK, clV := range clMap {
		if clK != cluster {
			newClMap[clK] = clV
		}
	}
	return newClMap, nil
}

func RemoveAuthInfo(authMap map[string]*clientcmdapi.AuthInfo, authInfo string) (map[string]*clientcmdapi.AuthInfo, error) {
	clValue := authMap[authInfo]
	if clValue == nil {
		zap.S().Debugf("🐛 User '%s' not found", authInfo)
		return nil, errors.New("user not found")
	}

	newAuthMap := make(map[string]*clientcmdapi.AuthInfo)
	for authK, authV := range authMap {
		if authK != authInfo {
			newAuthMap[authK] = authV
		}
	}
	return newAuthMap, nil
}

func GetContextsKeys(contexts map[string]*clientcmdapi.Context) []string {
	keys := make([]string, 0, len(contexts))
	for k := range contexts {
		keys = append(keys, k)
	}
	return keys
}

func GetClustersKeys(clusters map[string]*clientcmdapi.Cluster) []string {
	keys := make([]string, 0, len(clusters))
	for k := range clusters {
		keys = append(keys, k)
	}
	return keys
}

func GetAuthInfosKeys(auths map[string]*clientcmdapi.AuthInfo) []string {
	keys := make([]string, 0, len(auths))
	for k := range auths {
		keys = append(keys, k)
	}
	return keys
}

func PrintOnLogs(kubeConfig *clientcmdapi.Config, isDebug bool) {
	if isDebug {
		// logger.SugaredLogger.Debugf("Api version: %s", kubeConfig.APIVersion)
		// logger.SugaredLogger.Debugf("Kind: %s", kubeConfig.Kind)
		zap.S().Debugf("Current context: %s", kubeConfig.CurrentContext)
		zap.S().Debugf("Contexts: %s", strings.Join(GetContextsKeys(kubeConfig.Contexts), ", "))
		zap.S().Debugf("Clusters: %s", strings.Join(GetClustersKeys(kubeConfig.Clusters), ", "))
		zap.S().Debugf("Users: %s", strings.Join(GetAuthInfosKeys(kubeConfig.AuthInfos), ", "))
	} else {
		// logger.SugaredLogger.Infof("Api version: %s", kubeConfig.APIVersion)
		// logger.SugaredLogger.Infof("Kind: %s", kubeConfig.Kind)
		zap.S().Infof("Current context: %s", kubeConfig.CurrentContext)
		zap.S().Infof("Contexts: %s", strings.Join(GetContextsKeys(kubeConfig.Contexts), ", "))
		zap.S().Infof("Clusters: %s", strings.Join(GetClustersKeys(kubeConfig.Clusters), ", "))
		zap.S().Infof("Users: %s", strings.Join(GetAuthInfosKeys(kubeConfig.AuthInfos), ", "))
	}
}
