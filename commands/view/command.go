package view

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/logger"
	"github.com/bygui86/konf-sh/utils"

	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create VIEW-CONFIG command")
	home := utils.GetHomeDirOrExit("view-cfg")
	return &cli.Command{
		Name:   "view-cfg",
		Usage:  "View local and global Kubernetes contexts",
		Action: view,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
				Usage:    commons.CustomKubeConfigFlagDescription,
				EnvVar:   commons.CustomKubeConfigPathEnvVar,
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
		},
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
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
						Usage:    commons.CustomKubeConfigFlagDescription,
						EnvVar:   commons.CustomKubeConfigPathEnvVar,
						Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
						Required: false,
					},
				},
			},
		},
	}
}
