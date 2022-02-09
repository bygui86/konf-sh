package reset

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create RESET command")
	home := commons.GetHomeDirOrExit("reset")
	return &cli.Command{
		Name:  "reset",
		Usage: "Reset Kubernetes configuration",
		Subcommands: cli.Commands{
			{
				Name:      "local",
				Usage:     "Reset local Kubernetes configuration (current shell)",
				ArgsUsage: "<context>",
				Action:    resetLocal,
			},
			{
				Name:      "global",
				Usage:     "Reset global Kubernetes configuration",
				ArgsUsage: "<context>",
				Action:    resetGlobal,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     commons.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
						Usage:    commons.KubeConfigFlagDescription,
						EnvVars:  []string{commons.KubeConfigPathEnvVar},
						Value:    kubeconfig.GetKubeConfigPathDefault(home),
						Required: false,
					},
				},
			},
		},
	}
}
