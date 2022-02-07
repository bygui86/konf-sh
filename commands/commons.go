package commands

import (
	"fmt"

	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/logger"
	"github.com/urfave/cli"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func ValidateAndWrite(kubeConfig *clientcmdapi.Config, kubeConfigFilePath string) error {
	logger.SugaredLogger.Debugf("🐛 Validate cleaned Kubernetes configuration")
	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error validating cleaned Kubernetes configuration from '%s': %s", kubeConfigFilePath, newValErr.Error()),
			12)
	}

	logger.SugaredLogger.Debugf("🐛 Write cleaned Kubernetes configuration to file '%s'", kubeConfigFilePath)
	newWriteErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if newWriteErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error writing cleaned Kubernetes configuration '%s' to file: %s", kubeConfigFilePath, newWriteErr.Error()),
			13)
	}

	return nil
}
