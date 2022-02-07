package reset

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RESET-CONFIG command")
	home := utils.GetHomeDirOrExit("reset-cfg")
	return &cli.Command{
		Name:  "reset-cfg",
		Usage: "Reset local or global Kubernetes configuration",
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
