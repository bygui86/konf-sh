package view

import (
	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/logger"

	"github.com/urfave/cli"
)

func view(ctx *cli.Context) error {
	logger.Logger.Debug("")
	logger.Logger.Debug("🐛 Executing VIEW command")

	viewLocal(ctx)
	viewGlobal(ctx)
	return nil
}

func viewLocal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-LOCAL command")
	logger.Logger.Debug("")

	kubeConfig := kubeconfig.GetKubeConfigEnvVar()
	if kubeConfig != "" {
		logger.SugaredLogger.Infof("💻 Local Kuberntes context: '%s'", kubeConfig)
	} else {
		logger.SugaredLogger.Infof("💻 No local Kuberntes context set")
	}

	logger.Logger.Info("")
	return nil
}

func viewGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-GLOBAL command")
	logger.Logger.Debug("")

	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Infof("🌍 Global Kuberntes context: '%s'", kubeConfig.CurrentContext)

	logger.Logger.Info("")
	return nil
}
