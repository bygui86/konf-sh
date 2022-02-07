package commons

import (
	"fmt"

	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/urfave/cli/v2"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	// Flags
	CustomKubeConfigFlagName        = "kube-config"
	CustomKubeConfigFlagShort       = "k"
	CustomKubeConfigFlagDescription = "Kubernetes configuration custom (`PATH`)"
	SingleConfigsFlagName           = "single-configs"
	SingleConfigsFlagShort          = "c"
	SingleConfigsFlagDescription    = "Single Kubernetes configurations files custom (`PATH`)"

	// Environment variables
	CustomKubeConfigPathEnvVar = "KONF_KUBE_CONFIG_PATH"
	SingleConfigsPathEnvVar    = "KONF_SINGLE_KUBE_CONFIGS_PATH"
	KubeConfigEnvVar           = "KUBECONFIG"

	// Defaults
	CustomKubeConfigPathDefault = ".kube/config"
	SingleConfigsPathDefault    = ".kube/configs"
	KubeConfigEnvVarDefault     = ""
)

func ValidateAndWrite(kubeConfig *clientcmdapi.Config, kubeConfigFilePath string) error {
	logger.SugaredLogger.Debugf("🐛 Validate cleaned Kubernetes configuration")
	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error validating cleaned Kubernetes configuration from '%s': %s", kubeConfigFilePath, newValErr.Error()),
			12)
	}

	logger.SugaredLogger.Debugf("🐛 Write cleaned Kubernetes configuration to file '%s'", kubeConfigFilePath)
	newWriteErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error writing cleaned Kubernetes configuration '%s' to file: %s", kubeConfigFilePath, newWriteErr.Error()),
			13)
	}

	return nil
}
