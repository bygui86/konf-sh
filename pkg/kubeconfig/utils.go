package kubeconfig

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"go.uber.org/zap"
	clientcmdApi "k8s.io/client-go/tools/clientcmd/api"
)

func GetKubeConfigEnvVar() string {
	return commons.GetString(commons.KubeConfigEnvVar, commons.KubeConfigEnvVarDefault)
}

func GetKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.KubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleKonfigsPathDefault)
}

func CheckIfContextExist(kubeConfig *clientcmdApi.Config, context string) error {
	zap.S().Debugf("🐛 Check context '%s' existence in Kubernetes configuration '%s'", context, kubeConfig.CurrentContext)
	ctxValue := kubeConfig.Contexts[context]
	if ctxValue == nil {
		zap.S().Debugf("🐛 Context '%s' not found", context)
		return errors.New("context not found")
	}
	zap.S().Debugf("🐛 Context '%s' is valid", context)
	return nil
}

func CheckIfClusterExist(kubeConfig *clientcmdApi.Config, cluster string) error {
	zap.S().Debugf("🐛 Check cluster '%s' existence in Kubernetes configuration '%s'", cluster, kubeConfig.CurrentContext)
	clValue := kubeConfig.Clusters[cluster]
	if clValue == nil {
		zap.S().Debugf("🐛 Cluster '%s' not found", cluster)
		return errors.New("cluster not found")
	}
	zap.S().Debugf("🐛 Cluster '%s' is valid", cluster)
	return nil
}

func CheckIfAuthInfoExist(kubeConfig *clientcmdApi.Config, authInfo string) error {
	zap.S().Debugf("🐛 Check user '%s' existence in Kubernetes configuration '%s'", authInfo, kubeConfig.CurrentContext)
	authValue := kubeConfig.AuthInfos[authInfo]
	if authValue == nil {
		zap.S().Debugf("🐛 User '%s' not found", authInfo)
		return errors.New("user not found")
	}
	zap.S().Debugf("🐛 User '%s' is valid", authInfo)
	return nil
}

func GetContextsKeys(contexts map[string]*clientcmdApi.Context) []string {
	keys := make([]string, 0, len(contexts))
	for k := range contexts {
		keys = append(keys, k)
	}
	return keys
}

func GetClustersKeys(clusters map[string]*clientcmdApi.Cluster) []string {
	keys := make([]string, 0, len(clusters))
	for k := range clusters {
		keys = append(keys, k)
	}
	return keys
}

func GetAuthInfosKeys(auths map[string]*clientcmdApi.AuthInfo) []string {
	keys := make([]string, 0, len(auths))
	for k := range auths {
		keys = append(keys, k)
	}
	return keys
}

// PrintOnLogs is for development and debugging purposes only
func PrintOnLogs(kCfg *clientcmdApi.Config, isDebug bool) {
	if isDebug {
		// zap.S().Debugf("Api version: %s", kCfg.APIVersion)
		// zap.S().Debugf("Kind: %s", kCfg.Kind)
		zap.S().Debugf("Current context: %s", kCfg.CurrentContext)
		zap.S().Debugf("Contexts: %s", strings.Join(GetContextsKeys(kCfg.Contexts), ", "))
		zap.S().Debugf("Clusters: %s", strings.Join(GetClustersKeys(kCfg.Clusters), ", "))
		zap.S().Debugf("Users: %s", strings.Join(GetAuthInfosKeys(kCfg.AuthInfos), ", "))
	} else {
		// logger.SugaredLogger.Infof("Api version: %s", kCfg.APIVersion)
		// logger.SugaredLogger.Infof("Kind: %s", kCfg.Kind)
		zap.S().Infof("Current context: %s", kCfg.CurrentContext)
		zap.S().Infof("Contexts: %s", strings.Join(GetContextsKeys(kCfg.Contexts), ", "))
		zap.S().Infof("Clusters: %s", strings.Join(GetClustersKeys(kCfg.Clusters), ", "))
		zap.S().Infof("Users: %s", strings.Join(GetAuthInfosKeys(kCfg.AuthInfos), ", "))
	}
}
