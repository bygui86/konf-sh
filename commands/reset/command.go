package reset

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RESET command")
	return &cli.Command{
		Name:   "reset",
		Usage:  "",
		Action: reset,
	}
}
