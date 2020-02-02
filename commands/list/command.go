package list

import (
	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/utils"

	"github.com/urfave/cli"

	"github.com/bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create LIST-CONFIG command")
	home := utils.GetHomeDirOrExit("list-cfg")
	return &cli.Command{
		Name:   "list-cfg",
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
