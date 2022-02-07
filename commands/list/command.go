package list

import (
	"github.com/bygui86/konf-sh/commons"
	"github.com/bygui86/konf-sh/kubeconfig"
	"github.com/bygui86/konf-sh/utils"

	"github.com/urfave/cli"

	"github.com/bygui86/konf-sh/logger"
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
