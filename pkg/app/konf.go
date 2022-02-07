package app

import (
	"os"
	"sort"

	"github.com/bygui86/konf-sh/cmd/clean"
	"github.com/bygui86/konf-sh/cmd/completion"
	"github.com/bygui86/konf-sh/cmd/list"
	"github.com/bygui86/konf-sh/cmd/rename"
	"github.com/bygui86/konf-sh/cmd/reset"
	"github.com/bygui86/konf-sh/cmd/set"
	"github.com/bygui86/konf-sh/cmd/split"
	"github.com/bygui86/konf-sh/cmd/view"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/urfave/cli/v2"
)

const (
	appName    = "konf"
	appUsage   = "Kubernetes Configurator for an easy daily job"
	appVersion = "v0.5"
)

type KubeConfiguratorApp struct {
	app *cli.App
}

func Create() *KubeConfiguratorApp {
	logger.Logger.Debug("🐛 Creating application")
	app := cli.NewApp()
	setGlobalConfig(app)
	addCommands(app)
	setLastConfig(app)
	return &KubeConfiguratorApp{
		app: app,
	}
}

func setGlobalConfig(app *cli.App) {
	logger.Logger.Debug("🐛 Setting global configurations")
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
}

func addCommands(app *cli.App) {
	logger.Logger.Debug("🐛 Adding commands")
	app.Commands = []*cli.Command{
		split.BuildCommand(),
		list.BuildCommand(),
		view.BuildCommand(),
		set.BuildCommand(),
		clean.BuildCommand(),
		rename.BuildCommand(),
		reset.BuildCommand(),
		completion.BuildCommand(),
	}
}

func setLastConfig(app *cli.App) {
	logger.Logger.Debug("🐛 Setting last configurations")
	// sort flags in help section
	sort.Sort(cli.FlagsByName(app.Flags))
	// sort commands in help section
	sort.Sort(cli.CommandsByName(app.Commands))
}

func (k *KubeConfiguratorApp) Start() {
	logger.Logger.Debug("🐛 Starting application")
	err := k.app.Run(os.Args)
	if err != nil {
		logger.SugaredLogger.Errorf("❌ Error starting application: %s", err.Error())
		os.Exit(2)
	}
}
