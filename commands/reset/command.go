package reset

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/utils"
	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RESET-CONFIG command")
	home := utils.GetHomeDirOrExit("reset-cfg")
	return &cli.Command{
		Name:   "reset-cfg",
		Usage:  "Reset local or global Kubernetes configuration",
		Subcommands: cli.Commands{
			{
				Name:   "local",
				Usage:  "Reset local Kubernetes configuration (current shell)",
				ArgsUsage: "<context>",
				Action: resetLocal,
			},
			{
				Name:   "global",
				Usage:  "Reset global Kubernetes configuration",
				ArgsUsage: "<context>",
				Action: resetGlobal,
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
