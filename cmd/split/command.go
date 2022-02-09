package split

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create SPLIT command")
	home := commons.GetHomeDirOrExit("split")
	return &cli.Command{
		Name:   "split",
		Usage:  "Split Kubernetes configuration into multiple single Kubernetes konfigurations based on the context",
		Action: split,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     commons.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
				Usage:    commons.KubeConfigFlagDescription,
				EnvVars:  []string{commons.KubeConfigPathEnvVar},
				Value:    kubeconfig.GetKubeConfigPathDefault(home),
				Required: false,
			},
			&cli.StringFlag{
				Name:     commons.GetUrfaveFlagName(commons.SingleKonfigsFlagName, commons.SingleKonfigsFlagShort),
				Usage:    commons.SingleKonfigsFlagDescription,
				EnvVars:  []string{commons.SingleKonfigsPathEnvVar},
				Value:    kubeconfig.GetSingleConfigsPathDefault(home),
				Required: false,
			},
		},
	}
}
