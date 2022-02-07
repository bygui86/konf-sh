package reset

import (
	"fmt"
	"github.com/bygui86/konf-sh/commands"
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/logger"
)

func resetLocal(ctx *cli.Context) error {
	logger.Logger.Debug("")
	logger.Logger.Debug("🐛 Executing RESET-CFG-LOCAL command")
	logger.Logger.Debug("")

	logger.Logger.Info(fmt.Sprintf("unset %s", commons.KubeConfigEnvVar))
	return nil
}

func resetGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing RESET-CFG-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("♻️ Reset Kubernetes configuration '%s'", kubeConfigFilePath)
	kubeConfig.CurrentContext = ""

	valWrErr := commands.ValidateAndWrite(kubeConfig, kubeConfigFilePath)
	if valWrErr != nil {
		return valWrErr
	}

	logger.SugaredLogger.Infof("✅ Completed! Kubernetes configuration reset")
	logger.Logger.Info("")
	return nil
}
