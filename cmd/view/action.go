package view

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func view(ctx *cli.Context) error {
	logger.Logger.Debug("")
	logger.Logger.Debug("🐛 Executing VIEW command")

	_ = viewLocal(ctx)
	_ = viewGlobal(ctx)

	return nil
}

// INFO: viewLocal function returns error even if always nil as used to build cli command
func viewLocal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-LOCAL command")
	logger.Logger.Debug("")

	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable", commons.KubeConfigEnvVar)
	kCfg := kubeconfig.GetKubeConfigEnvVar()
	if kCfg != "" {
		logger.SugaredLogger.Infof("💻 Local Kubernetes context: '%s'", kCfg)
	} else {
		logger.SugaredLogger.Infof("💻 No local Kubernetes context set or default to '%s/%s'", utils.GetHomeDirOrExit("view-local"), commons.KubeConfigPathDefault)
	}

	logger.Logger.Info("")
	return nil
}

// INFO: viewGlobal function returns error even if always nil as used to build cli command
func viewGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing VIEW-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.KubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kCfg := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Infof("🌍 Global Kubernetes context: '%s'", kCfg.CurrentContext)

	logger.Logger.Info("")
	return nil
}
