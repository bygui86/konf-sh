package app

import (
	"os"

	"github.com/bygui86/konf-sh/pkg/logging"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type kubeConfiguratorApp struct {
	app *cli.App
}

func Create() *kubeConfiguratorApp {
	logging.InitLogger()

	zap.L().Debug("🐛 Creating application")
	app := cli.NewApp()
	setGlobalConfig(app)
	addCommands(app)
	setLastConfig(app)
	return &kubeConfiguratorApp{
		app: app,
	}
}

func (k *kubeConfiguratorApp) Start() {
	zap.L().Debug("🐛 Starting application")
	err := k.app.Run(os.Args)
	if err != nil {
		zap.S().Errorf("❌  Error starting application: %s", err.Error())
		os.Exit(2)
	}
}
