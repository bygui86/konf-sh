package kubeconfig

import (
	"path/filepath"

	"bygui86/konf/commons"
	"bygui86/konf/utils"
)

func GetKubeConfigEnvVar() string {
	return utils.GetString(commons.KubeConfigEnvVar, commons.KubeConfigEnvVarDefault)
}

func GetCustomKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.CustomKubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleConfigsPathDefault)
}
