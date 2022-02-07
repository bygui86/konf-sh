package view

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/logger"
	"github.com/bygui86/konf-sh/utils"

	"github.com/urfave/cli"
)

func view(ctx *cli.Context) error {
	logger.Logger.Debug("")
	logger.Logger.Debug("🐛 Executing VIEW-CONFIG command")

	viewLocal(ctx)
	viewGlobal(ctx)

	return nil
}

func viewLocal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-LOCAL command")
	logger.Logger.Debug("")

	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable", commons.KubeConfigEnvVar)
	kubeConfig := kubeconfig.GetKubeConfigEnvVar()
	if kubeConfig != "" {
		logger.SugaredLogger.Infof("💻 Local Kubernetes context: '%s'", kubeConfig)
	} else {
		logger.SugaredLogger.Infof("💻 No local Kubernetes context set or default to '%s/%s'", utils.GetHomeDirOrExit("view-local"), commons.CustomKubeConfigPathDefault)
	}

	logger.Logger.Info("")
	return nil
}

func viewGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Infof("🌍 Global Kubernetes context: '%s'", kubeConfig.CurrentContext)

	logger.Logger.Info("")
	return nil
}
