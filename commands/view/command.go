package view

import (
	"github.com/urfave/cli"

	"bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create view command")
	return &cli.Command{
		Name:   "view",
		Usage:  "View local and global Kubernetes contexts",
		Action: view,
		Subcommands: cli.Commands{
			{
				Name:   "local",
				Usage:  "View local Kubernetes context (current shell)",
				Action: viewLocal,
			},
			{
				Name:   "global",
				Usage:  "View global Kubernetes context",
				Action: viewGlobal,
			},
		},
	}
}
