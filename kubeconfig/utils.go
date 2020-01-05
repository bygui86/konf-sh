package kubeconfig

import (
	"os"
	"path/filepath"

	"bygui86/kubeconfigurator/config/envvar"
	"bygui86/kubeconfigurator/config/flags"
)

const (
	// ENV-VAR
	kubeConfigEnvVar = "KUBECONFIG"
	kubeConfigDefault = ""

	// FLAGS
	kubeConfigFlagKey = "kubeconfig"
	kubeConfigFolderDefault = ".kube"
	kubeConfigFilenameDefault = "config"

	// OTHERS
	kubeConfigOutputFolderPerm = 0755
)

func GetKubeConfigEnvVar() string {
	return envvar.GetString(kubeConfigEnvVar, kubeConfigDefault)
}

func SetKubeConfigEnvVar(kubeConfigsPath string) error {
	return envvar.Set(kubeConfigEnvVar, kubeConfigsPath)
}

func GetKubeConfig(home string) (string,error) {
	return flags.GetString(
		kubeConfigFlagKey,
		filepath.Join(home, kubeConfigFolderDefault, kubeConfigFilenameDefault),
		"(optional) absolute path to the kubeconfig file")
}

func CheckKubeConfigsFolder(path string) error {
	_, statErr := os.Stat(path)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			return os.Mkdir(path, kubeConfigOutputFolderPerm)
		}
		return statErr
	}
	return nil
}
