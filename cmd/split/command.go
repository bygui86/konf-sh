package split

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create SPLIT command")
	home := utils.GetHomeDirOrExit("split")
	return &cli.Command{
		Name:   "split",
		Usage:  "Split Kubernetes configuration into multiple single Kubernetes konfigurations based on the context",
		Action: split,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
				Usage:    commons.KubeConfigFlagDescription,
				EnvVars:  []string{commons.KubeConfigPathEnvVar},
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
			&cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.SingleKonfigsFlagName, commons.SingleKonfigsFlagShort),
				Usage:    commons.SingleKonfigsFlagDescription,
				EnvVars:  []string{commons.SingleKonfigsPathEnvVar},
				Value:    kubeconfig.GetSingleConfigsPathDefault(home),
				Required: false,
			},
		},
	}
}
