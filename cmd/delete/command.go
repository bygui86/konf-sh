package delete

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create DELETE command")
	home := utils.GetHomeDirOrExit("delete")
	return &cli.Command{
		Name:      "delete",
		Usage:     "Remove all specified contexts from Kubernetes configuration and single Kubernetes konfigurations",
		ArgsUsage: "<context-list_comma-separated>",
		Action:    deleteCtx,
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
