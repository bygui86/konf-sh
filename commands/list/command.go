package list

import (
	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/utils"

	"github.com/urfave/cli"
	
	"bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create LIST command")
	home := utils.GetHomeDirOrExit("list")
	return &cli.Command{
		Name:   "list",
		Usage:  "List all single Kubernetes configurations",
		Action: list,
		Flags: []cli.Flag{
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
