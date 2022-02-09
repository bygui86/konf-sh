package reset

import (
	"fmt"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func resetLocal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing RESET-LOCAL command")

	zap.L().Info(fmt.Sprintf("unset %s", commons.KubeConfigEnvVar)) // TODO to be replaced by following line
	//zap.L().Info() // TODO enable when shell wrapper is available
	return nil
}

func resetGlobal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing RESET-GLOBAL command")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Debugf("🧹️ Reset Kubernetes configuration '%s'", kubeConfigFilePath)
	kubeConfig.CurrentContext = ""

	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s",
				kubeConfigFilePath, newValErr.Error()), 12)
	}

	newWriteErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error writing Kubernetes configuration '%s' to file: %s",
				kubeConfigFilePath, newWriteErr.Error()), 13)
	}

	zap.S().Infof("✅ Global Kubernetes configuration reset")
	zap.L().Info("")
	return nil
}
