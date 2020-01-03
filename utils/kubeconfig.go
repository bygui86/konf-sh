package utils

import (
	"bygui86/kubeconfigurator/envvar"
	"os"
)

const (
	kubeConfigEnvVar = "KUBECONFIG"

	kubeConfigOutputFolderPerm = 0755
)

func GetKubeConfigEnvVar() string {
	return envvar.GetString(kubeConfigEnvVar, "")
}

func SetKubeConfigEnvVar(kubeConfigsPath string) error {
	return envvar.Set(kubeConfigEnvVar, kubeConfigsPath)
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
