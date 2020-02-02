package rename

import (
	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/utils"
	"github.com/urfave/cli"

	"github.com/bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RENAME-CONTEXT command")
	home := utils.GetHomeDirOrExit("rename-ctx")
	return &cli.Command{
		Name:   "rename",
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
