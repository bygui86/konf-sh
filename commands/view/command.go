package view

import (
	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/logger"
	"bygui86/konf/utils"

	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create view command")
	home := utils.GetHomeDirOrExit("view")
	return &cli.Command{
		Name:   "view",
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
