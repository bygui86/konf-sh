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
	zap.L().Debug("🐛 Executing VIEW-LOCAL command")

	zap.S().Debugf("🐛 Get '%s' environment variable value", commons.KubeConfigEnvVar)
	kCfgEnvVarValue := kubeconfig.GetKubeConfigEnvVar()
	if kCfgEnvVarValue != "" {
		zap.S().Infof("💻 Local Kubernetes context: '%s'", kCfgEnvVarValue)
	} else {
		zap.L().Info("💻 No local Kubernetes context set")
	}

	zap.L().Info("")
	return nil
}

// INFO: viewGlobal function returns error even if always nil as used to build cli command
func viewGlobal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing VIEW-GLOBAL command")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kCfg := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Infof("🌍 Global Kubernetes context: '%s'", kCfg.CurrentContext)

	zap.L().Info("")
	return nil
}
