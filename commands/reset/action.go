package reset

import (
	"github.com/urfave/cli"

	"github.com/bygui86/konf/logger"
)

func reset(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing RESET command")
	logger.Logger.Debug("")

	logger.Logger.Warn("⚠️ Command not yet implemented ☢️")

	logger.Logger.Info("")
	return nil
}
