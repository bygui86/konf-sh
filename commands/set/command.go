package set

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/logger"
	"github.com/bygui86/konf/utils"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create SET command")
	home := utils.GetHomeDirOrExit("set")
	return &cli.Command{
		Name:  "set",
		Usage: "Set local or global Kubernetes context",
		Subcommands: cli.Commands{
			{
				Name:   "local",
				Usage:  "Set local Kubernetes context (current shell)",
				ArgsUsage: "<context>",
				Action: setLocal,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     utils.GetUrfaveFlagName(commons.SingleConfigsFlagName, commons.SingleConfigsFlagShort),
						Usage:    commons.SingleConfigsFlagDescription,
						EnvVar:   commons.SingleConfigsPathEnvVar,
						Value:    kubeconfig.GetSingleConfigsPathDefault(home),
						Required: false,
					},
				},
			},
			{
				Name:   "global",
				Usage:  "Set global Kubernetes context",
				ArgsUsage: "<context>",
				Action: setGlobal,
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
