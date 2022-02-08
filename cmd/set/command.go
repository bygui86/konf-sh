package set

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create SET-CONFIG command")
	home := commons.GetHomeDirOrExit("set")
	return &cli.Command{
		Name:  "set",
		Usage: "Set local or global Kubernetes context",
		Subcommands: cli.Commands{
			{
				Name:      "local",
				Usage:     "Set local Kubernetes context (current shell)",
				ArgsUsage: "<context>",
				Action:    setLocal,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     commons.GetUrfaveFlagName(commons.SingleKonfigsFlagName, commons.SingleKonfigsFlagShort),
						Usage:    commons.SingleKonfigsFlagDescription,
						EnvVars:  []string{commons.SingleKonfigsPathEnvVar},
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
