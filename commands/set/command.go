package set

import (
	"github.com/urfave/cli"

	"bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create set command")
	return &cli.Command{
		Name:  "set",
		Usage: "Set local or global Kubernetes context",
		Subcommands: cli.Commands{
			{
				Name:   "local",
				Usage:  "Set local Kubernetes context (current shell)",
				Action: setLocal,
			},
			{
				Name:   "global",
				Usage:  "Set global Kubernetes context",
				Action: setGlobal,
			},
		},
	}
}
