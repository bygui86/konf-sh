package set

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create SET-CONFIG command")
	home := utils.GetHomeDirOrExit("set-cfg")
	return &cli.Command{
		Name:  "set-cfg",
		Usage: "Set local or global Kubernetes context",
		Subcommands: cli.Commands{
			{
				Name:      "local",
				Usage:     "Set local Kubernetes context (current shell)",
				ArgsUsage: "<context>",
				Action:    setLocal,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     utils.GetUrfaveFlagName(commons.SingleConfigsFlagName, commons.SingleConfigsFlagShort),
						Usage:    commons.SingleConfigsFlagDescription,
						EnvVars:  []string{commons.SingleConfigsPathEnvVar},
						Value:    kubeconfig.GetSingleConfigsPathDefault(home),
						Required: false,
					},
				},
			},
			{
				Name:      "global",
				Usage:     "Set global Kubernetes context",
				ArgsUsage: "<context>",
				Action:    setGlobal,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
						Usage:    commons.CustomKubeConfigFlagDescription,
						EnvVars:  []string{commons.CustomKubeConfigPathEnvVar},
						Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
						Required: false,
					},
				},
			},
		},
	}
}
