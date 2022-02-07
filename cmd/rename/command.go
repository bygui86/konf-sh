package rename

import (
	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RENAME-CONTEXT command")
	home := utils.GetHomeDirOrExit("rename-ctx")
	return &cli.Command{
		Name:      "rename-ctx",
		Usage:     "Rename specified context with a new one",
		ArgsUsage: "<context-to-rename> <new-context-name>",
		Action:    rename,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
				Usage:    commons.CustomKubeConfigFlagDescription,
				EnvVars:  []string{commons.CustomKubeConfigPathEnvVar},
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
		},
	}
}
