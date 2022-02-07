package reset

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RESET command")
	home := utils.GetHomeDirOrExit("reset")
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
						Name:     utils.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
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
