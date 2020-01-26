package clean

import (
	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/utils"

	"github.com/urfave/cli"
	
	"bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create CLEAN command")
	return &cli.Command{
		Name:   "clean",
		Usage:  "",
		Action: clean,
	}
}
