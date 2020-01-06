package list

import (
	"bygui86/kubeconfigurator/commons"
	"bygui86/kubeconfigurator/kubeconfig"
	"bygui86/kubeconfigurator/utils"

	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {
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
