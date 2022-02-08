package app

import (
	"os"
	"sort"

	"github.com/bygui86/konf-sh/cmd/completion"
	"github.com/bygui86/konf-sh/cmd/delete"
	"github.com/bygui86/konf-sh/cmd/list"
	"github.com/bygui86/konf-sh/cmd/rename"
	"github.com/bygui86/konf-sh/cmd/reset"
	"github.com/bygui86/konf-sh/cmd/set"
	"github.com/bygui86/konf-sh/cmd/split"
	"github.com/bygui86/konf-sh/cmd/view"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const (
	appName    = "konf-sh"
	appUsage   = "The KubeConfig manager for your shell"
	appVersion = "v0.5"
)

type KubeConfiguratorApp struct {
	app *cli.App
}

func Create() *KubeConfiguratorApp {
	zap.L().Debug("🐛 Creating application")
	app := cli.NewApp()
	setGlobalConfig(app)
	addCommands(app)
	setLastConfig(app)
	return &KubeConfiguratorApp{
		app: app,
	}
}

func setGlobalConfig(app *cli.App) {
	zap.L().Debug("🐛 Setting global configurations")
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
}

func addCommands(app *cli.App) {
	zap.L().Debug("🐛 Adding commands")
	app.Commands = []*cli.Command{
		split.BuildCommand(),
		list.BuildCommand(),
		view.BuildCommand(),
		set.BuildCommand(),
		delete.BuildCommand(),
		rename.BuildCommand(),
		reset.BuildCommand(),
		completion.BuildCommand(),
	}
}

func setLastConfig(app *cli.App) {
	zap.L().Debug("🐛 Setting last configurations")
	// sort flags in help section
	sort.Sort(cli.FlagsByName(app.Flags))
	// sort commands in help section
	sort.Sort(cli.CommandsByName(app.Commands))
}

func (k *KubeConfiguratorApp) Start() {
	zap.L().Debug("🐛 Starting application")
	err := k.app.Run(os.Args)
	if err != nil {
		zap.S().Errorf("❌ Error starting application: %s", err.Error())
		os.Exit(2)
	}
}
