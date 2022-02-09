package kubeconfig

import (
	"errors"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdApi "k8s.io/client-go/tools/clientcmd/api"
)

func Load(kubeConfigFilePath string) *clientcmdApi.Config {
	zap.S().Debugf("🐛 Load Kubernetes configuration from file '%s'", kubeConfigFilePath)
	return clientcmd.GetConfigFromFileOrDie(kubeConfigFilePath)
}

func ListContexts(kubeConfig *clientcmdApi.Config) []string {
	zap.L().Debug("🐛 List available Kubernetes contexts")
	ctxs := make([]string, len(kubeConfig.Contexts))
	idx := 0
	for ctx := range kubeConfig.Contexts {
		ctxs[idx] = ctx
		idx++
	}
	return ctxs
}

func Split(kubeConfig *clientcmdApi.Config, kubeConfigFilePath string) map[string]*clientcmdApi.Config {
	zap.S().Debugf("🐛 Split Kubernetes configuration from %s", kubeConfigFilePath)
	singleConfigs := make(map[string]*clientcmdApi.Config, len(kubeConfig.Contexts))
	for ctxKey, ctxValue := range kubeConfig.Contexts {
		contexts := make(map[string]*clientcmdApi.Context, 1)
		contexts[ctxKey] = ctxValue
		clusters := make(map[string]*clientcmdApi.Cluster, 1)
		clusters[ctxValue.Cluster] = kubeConfig.Clusters[ctxValue.Cluster]
		authInfos := make(map[string]*clientcmdApi.AuthInfo, 1)
		authInfos[ctxValue.AuthInfo] = kubeConfig.AuthInfos[ctxValue.AuthInfo]
		singleConfigs[ctxKey] = &clientcmdApi.Config{
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

func Validate(kubeConfig *clientcmdApi.Config) error {
	zap.S().Debugf("🐛 Validate Kubernetes configuration")
	err := clientcmd.Validate(*kubeConfig)
	if clientcmd.IsConfigurationInvalid(err) {
		return err
	}
	return nil
}

func Write(kubeConfig *clientcmdApi.Config, filepath string) error {
	zap.S().Debugf("🐛 Write Kubernetes configuration to file '%s'", filepath)
	return clientcmd.WriteToFile(*kubeConfig, filepath)
}

func RemoveContext(ctxMap map[string]*clientcmdApi.Context, context string) (map[string]*clientcmdApi.Context, error) {
	ctxValue := ctxMap[context]
	if ctxValue == nil {
		zap.S().Debugf("🐛 Context '%s' not found", context)
		return nil, errors.New("context not found")
	}

	newCtxMap := make(map[string]*clientcmdApi.Context)
	for ctxK, ctxV := range ctxMap {
		if ctxK != context {
			newCtxMap[ctxK] = ctxV
		}
	}
	return newCtxMap, nil
}

func RemoveCluster(clMap map[string]*clientcmdApi.Cluster, cluster string) (map[string]*clientcmdApi.Cluster, error) {
	clValue := clMap[cluster]
	if clValue == nil {
		zap.S().Debugf("🐛 Cluster '%s' not found", cluster)
		return nil, errors.New("cluster not found")
	}

	newClMap := make(map[string]*clientcmdApi.Cluster)
	for clK, clV := range clMap {
		if clK != cluster {
			newClMap[clK] = clV
		}
	}
	return newClMap, nil
}

func RemoveAuthInfo(authMap map[string]*clientcmdApi.AuthInfo, authInfo string) (map[string]*clientcmdApi.AuthInfo, error) {
	clValue := authMap[authInfo]
	if clValue == nil {
		zap.S().Debugf("🐛 User '%s' not found", authInfo)
		return nil, errors.New("user not found")
	}

	newAuthMap := make(map[string]*clientcmdApi.AuthInfo)
	for authK, authV := range authMap {
		if authK != authInfo {
			newAuthMap[authK] = authV
		}
	}
	return newAuthMap, nil
}
