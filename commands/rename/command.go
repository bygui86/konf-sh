package rename

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf/logger"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create RENAME command")
	return &cli.Command{
		Name:   "rename",
		Usage:  "",
		Action: rename,
	}
}
