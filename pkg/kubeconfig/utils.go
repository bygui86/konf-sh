package kubeconfig

import (
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/utils"
)

func GetKubeConfigEnvVar() string {
	return utils.GetString(commons.KubeConfigEnvVar, commons.KubeConfigEnvVarDefault)
}

func GetCustomKubeConfigPathDefault(home string) string {
	return filepath.Join(home, commons.KubeConfigPathDefault)
}

func GetSingleConfigsPathDefault(home string) string {
	return filepath.Join(home, commons.SingleKonfigsPathDefault)
}
