package kubeconfig

import (
	"path/filepath"

	"bygui86/konf/commons"
	"bygui86/konf/utils"
)

const (
	kubeConfigEnvVar = "KUBECONFIG"

	kubeConfigEnvVarDefault = ""
)

func GetKubeConfigEnvVar() string {
	return utils.GetString(kubeConfigEnvVar, kubeConfigEnvVarDefault)
}

func SetKubeConfigEnvVar(kubeConfigNewValue string) error {
	return utils.Set(kubeConfigEnvVar, kubeConfigNewValue)
}

func GetCustomKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.CustomKubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleConfigsPathDefault)
}
