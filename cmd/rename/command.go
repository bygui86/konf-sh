package rename

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func BuildCommand() *cli.Command {
	zap.L().Debug("🐛 Create RENAME command")
	home := commons.GetHomeDirOrExit("rename")
	return &cli.Command{
		Name:      "rename",
		Usage:     "Rename specified context in Kubernetes configuration and single Kubernetes konfiguration",
		ArgsUsage: "<context-to-rename> <new-context-name>",
		Action:    rename,
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
