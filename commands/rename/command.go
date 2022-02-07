package rename

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/utils"
	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RENAME-CONTEXT command")
	home := utils.GetHomeDirOrExit("rename-ctx")
	return &cli.Command{
		Name:   "rename-ctx",
		Usage:  "Rename specified context with a new one",
		ArgsUsage: "<context-to-rename> <new-context-name>",
		Action: rename,
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
