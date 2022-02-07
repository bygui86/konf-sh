package rename

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RENAME command")
	home := utils.GetHomeDirOrExit("rename")
	return &cli.Command{
		Name:      "rename",
		Usage:     "Rename specified context with a new name",
		ArgsUsage: "<context-to-rename> <new-context-name>",
		Action:    rename,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.KubeConfigFlagName, commons.KubeConfigFlagShort),
				Usage:    commons.KubeConfigFlagDescription,
				EnvVars:  []string{commons.KubeConfigPathEnvVar},
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
		},
	}
}
