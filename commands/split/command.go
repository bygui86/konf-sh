package split

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/utils"

	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create SPLIT-CONFIG command")
	home := utils.GetHomeDirOrExit("split-cfg")
	return &cli.Command{
		Name:   "split-cfg",
		Usage:  "Split kube-config into multiple single Kubernetes configurations based on the context",
		Action: split,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
				Usage:    commons.CustomKubeConfigFlagDescription,
				EnvVar:   commons.CustomKubeConfigPathEnvVar,
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.SingleConfigsFlagName, commons.SingleConfigsFlagShort),
				Usage:    commons.SingleConfigsFlagDescription,
				EnvVar:   commons.SingleConfigsPathEnvVar,
				Value:    kubeconfig.GetSingleConfigsPathDefault(home),
				Required: false,
			},
		},
	}
}
