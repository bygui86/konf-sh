package set

import (
	"bygui86/konf/logger"

	"github.com/urfave/cli"
)

func setLocal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SET-LOCAL command")
	logger.Logger.Debug("")

	logger.Logger.Warn("⚠️ Command not yet implemented")

	// for _, arg := range ctx.Args() {
	// 	logger.SugaredLogger.Infof("Argument: %s", arg)
	// }

	return nil
}

func setGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SET-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Warn("⚠️ Command not yet implemented")
	return nil
}
