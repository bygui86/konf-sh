package split

import (
	"github.com/urfave/cli"
)

// TODO add flags
func BuildCommand() *cli.Command {
	return &cli.Command{
		Name:    "split",
		Usage:   "Split kube-config into multiple single Kubernetes configurations based on the context",
		Action:  split(),
		//Flags: []cli.Flag{
		//	cli.StringFlag{
		//		Name:     "path, p",
		//		Usage:    "`PATH` to scan for files",
		//		EnvVar:   "GOCLI_LIST_PATH,LIST_PATH",
		//		Value:    ".",
		//		Required: false,
		//	},
		//	cli.BoolFlag{
		//		Name:     "show-hidden, d",
		//		Usage:    "show hidden files",
		//		EnvVar:   "GOCLI_LIST_HIDDEN,LIST_HIDDEN",
		//		Required: false,
		//		// Value field not available, but `FALSE` per default
		//	},
		//},
	}
}
