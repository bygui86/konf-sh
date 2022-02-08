package kubeconfig

import (
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
)

func GetKubeConfigEnvVar() string {
	return commons.GetString(commons.KubeConfigEnvVar, commons.KubeConfigEnvVarDefault)
}

func GetCustomKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.KubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleKonfigsPathDefault)
}
