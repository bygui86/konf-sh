package clean

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/logger"
	"github.com/bygui86/konf/utils"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create CLEAN command")
	home := utils.GetHomeDirOrExit("clean")
	return &cli.Command{
		Name:      "clean",
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
