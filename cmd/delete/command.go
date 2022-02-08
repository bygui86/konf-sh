package delete

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create DELETE command")
	home := commons.GetHomeDirOrExit("delete")
	return &cli.Command{
		Name:      "delete",
		Usage:     "Remove all specified contexts from Kubernetes configuration and single Kubernetes konfigurations",
		ArgsUsage: "<context-list_comma-separated>",
		Action:    deleteCtx,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     commons.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
				Usage:    commons.KubeConfigFlagDescription,
				EnvVars:  []string{commons.KubeConfigPathEnvVar},
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
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
