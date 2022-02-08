package view

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create VIEW command")
	home := commons.GetHomeDirOrExit("view")
	return &cli.Command{
		Name:   "view",
		Usage:  "View local and global Kubernetes contexts",
		Action: view,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     commons.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
				Usage:    commons.KubeConfigFlagDescription,
				EnvVars:  []string{commons.KubeConfigPathEnvVar},
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
					&cli.StringFlag{
						Name:     commons.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
						Usage:    commons.KubeConfigFlagDescription,
						EnvVars:  []string{commons.KubeConfigPathEnvVar},
						Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
						Required: false,
					},
				},
			},
		},
	}
}
