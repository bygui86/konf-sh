package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
	fmt.Println("🏠 Get HOME")
	home, homeErr := utils.GetHomeDir()
	if homeErr != nil {
		fmt.Println("❌ ERROR: " + homeErr.Error())
		os.Exit(1)
	}

	fmt.Println("📚 Get Kubernetes configuration file")
	srcPath := filepath.Join(home, kubeConfigFolderDefault, kubeConfigFilenameDefault)
	kubeConfigFilePath, getErr := getKubeConfigFlagValue(kubeConfigFlagKey, srcPath)
	if getErr != nil {
		fmt.Println("❌ ERROR: " + getErr.Error())
		os.Exit(1)
	}
	srcConfig := getKubeConfig(kubeConfigFilePath)

	fmt.Println("✂️ Split Kubernetes configuration file")
	tgConfigs := splitKubeConfig(srcConfig)

	fmt.Println("💾 Save single Kubernetes configuration files")
	tgPath := filepath.Join(home, kubeConfigOutputFolderName)
	configsPath, saveErr := saveKubeConfigs(tgConfigs, tgPath)
	if saveErr != nil {
		fmt.Println("❌ ERROR: " + saveErr.Error())
		os.Exit(1)
	}
	fmt.Printf("✅ Completed, files stored in '%s'\n", configsPath)
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
	fmt.Printf("Check '%s' configs folder\n", configsPath)
	checkErr := utils.CheckKubeConfigsFolder(configsPath)
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
