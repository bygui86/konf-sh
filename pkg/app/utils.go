package app

import (
	"sort"

	"github.com/bygui86/konf-sh/cmd/completion"
	deleteCmd "github.com/bygui86/konf-sh/cmd/delete"
	"github.com/bygui86/konf-sh/cmd/list"
	"github.com/bygui86/konf-sh/cmd/rename"
	"github.com/bygui86/konf-sh/cmd/reset"
	"github.com/bygui86/konf-sh/cmd/set"
	"github.com/bygui86/konf-sh/cmd/split"
	"github.com/bygui86/konf-sh/cmd/view"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func setGlobalConfig(app *cli.App) {
	zap.L().Debug("🐛 Setting global configurations")
	app.Name = "konf-sh"
	app.Usage = "The KubeConfig manager for your shell"
	app.Version = "v0.5.0"
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
		deleteCmd.BuildCommand(),
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
