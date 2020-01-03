package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	clientcmd "k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	kubeConfigFlagKey = "kubeconfig"

	homeEnvVar = "HOME"
	homeEnvVarDefault = "~/"

	kubeConfigFolderDefault = ".kube"
	kubeConfigFilenameDefault = "config"

	kubeConfigOutputFolderName = ".kube/configs"
	kubeConfigOutputFolderPerm = 0755

	//kubeConfigEnvVar = "KUBECONFIG"
	kubeConfigEnvVar = "NEW_KUBECONFIG"
)

func main() {
	fmt.Println("📚 Get Kubernetes configuration file")
	kubeConfigFilePath := getKubeConfigFlagValue(kubeConfigFlagKey)
	srcConfig := getKubeConfig(*kubeConfigFilePath)

	fmt.Println("✂️ Split Kubernetes configuration file")
	tgConfigs := splitKubeConfig(srcConfig)

	fmt.Println(" Save single Kubernetes configuration files")
	configsPath, saveErr := saveKubeConfigs(tgConfigs)
	if saveErr != nil {
		fmt.Println("❌ ERROR: " + saveErr.Error())
		os.Exit(1)
	}
	fmt.Printf("✅ Completed, files stored in '%s'\n", configsPath)

	envErr := setKubeConfigEnvVar(configsPath)
	if envErr != nil {
		fmt.Println("❌ ERROR: " + envErr.Error())
		os.Exit(1)
	}
	fmt.Printf( "👀 New '%s' env var: '%s'\n", kubeConfigEnvVar, getKubeConfigEnvVar())
	
	fmt.Println("")
}

func getKubeConfigFlagValue(flagKey string) *string {
	home := getHomeDir()
	//if home != "" {
		return flag.String(flagKey, filepath.Join(home, kubeConfigFolderDefault, kubeConfigFilenameDefault), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	return flag.String(flagKey, "", "absolute path to the kubeconfig file")
	//}
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

func saveKubeConfigs(configs map[string]*clientcmdapi.Config) (string, error) {
	configsPath := buildKubeConfigsPath(getHomeDir(), kubeConfigOutputFolderName)

	fmt.Printf("Check '%s' configs folder\n", configsPath)
	checkErr := checkKubeConfigsFolder(configsPath, kubeConfigOutputFolderPerm)
	if checkErr != nil {
		return "", checkErr
	}

	for cfgKey, cfg := range configs {
		fmt.Printf("Validate config '%s'\n", cfgKey)
		valErr := validateKubeConfig(cfg)
		if valErr != nil {
			return "", valErr
		}
		fmt.Printf("Write config '%s' to file\n", cfgKey)
		writeErr := writeKubeConfig(cfg, filepath.Join(configsPath, cfgKey))
		if writeErr != nil {
			return "", writeErr
		}
	}

	return configsPath, nil
}

func checkKubeConfigsFolder(path string, mode os.FileMode) error {
	_, statErr := os.Stat(path)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			return os.Mkdir(path, mode)
		}
		return statErr
	}
	return nil
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

func getHomeDir() string {
	home := os.Getenv(homeEnvVar)
	if home != "" {
		return home
	}
	return homeEnvVarDefault
}

func buildKubeConfigsPath(home, path string) string {
	return filepath.Join(home, path)
}

func setKubeConfigEnvVar(kubeConfigsPath string) error {
	return os.Setenv(kubeConfigEnvVar, kubeConfigsPath)
}

func getKubeConfigEnvVar() string {
	return os.Getenv(kubeConfigEnvVar)
}
