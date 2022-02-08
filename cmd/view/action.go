package view

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func view(ctx *cli.Context) error {
	zap.L().Debug("")
	zap.L().Debug("🐛 Executing VIEW command")

	_ = viewLocal(ctx)
	_ = viewGlobal(ctx)

	return nil
}

// INFO: viewLocal function returns error even if always nil as used to build cli command
func viewLocal(ctx *cli.Context) error {
	zap.L().Info("")
	zap.L().Debug("🐛 Executing VIEW-LOCAL command")
	zap.L().Debug("")

	zap.S().Debugf("🐛 Get '%s' environment variable", commons.KubeConfigEnvVar)
	kCfg := kubeconfig.GetKubeConfigEnvVar()
	if kCfg != "" {
		zap.S().Infof("💻 Local Kubernetes context: '%s'", kCfg)
	} else {
		zap.S().Infof("💻 No local Kubernetes context set or default to '%s/%s'", commons.GetHomeDirOrExit("view-local"), commons.KubeConfigPathDefault)
	}

	zap.L().Info("")
	return nil
}

// INFO: viewGlobal function returns error even if always nil as used to build cli command
func viewGlobal(ctx *cli.Context) error {
	zap.L().Info("")
	zap.L().Debug("🐛 Executing VIEW-GLOBAL command")
	zap.L().Debug("")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kCfg := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Infof("🌍 Global Kubernetes context: '%s'", kCfg.CurrentContext)

	zap.L().Info("")
	return nil
}
