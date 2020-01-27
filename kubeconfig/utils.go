package kubeconfig

import (
	"path/filepath"

	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/utils"
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
