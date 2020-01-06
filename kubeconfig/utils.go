package kubeconfig

import (
	"path/filepath"

	"bygui86/konf/commons"
	"bygui86/konf/config/envvar"
)

func GetKubeConfigEnvVar() string {
	return envvar.GetString(commons.KubeConfigEnvVar, commons.KubeConfigEnvVarDefault)
}

func SetKubeConfigEnvVar(kubeConfigNewValue string) error {
	return envvar.Set(commons.KubeConfigEnvVar, kubeConfigNewValue)
}

func GetCustomKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.CustomKubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleConfigsPathDefault)
}
