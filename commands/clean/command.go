package clean

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/logger"
	"github.com/bygui86/konf-sh/utils"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create CLEAN-CONTEXT command")
	home := utils.GetHomeDirOrExit("clean-ctx")
	return &cli.Command{
		Name:      "clean-ctx",
		Usage:     "Remove specified context (and relative user and cluster) from Kubernetes configuration",
		ArgsUsage: "<context-list_comma-separated>",
		Action:    clean,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
				Usage:    commons.CustomKubeConfigFlagDescription,
				EnvVar:   commons.CustomKubeConfigPathEnvVar,
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
		},
	}
}
