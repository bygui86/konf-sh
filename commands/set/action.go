package set

import (
	"github.com/urfave/cli"

	"bygui86/konf/logger"
)

func setLocal(ctx *cli.Context) error {
	logger.Logger.Warn("⚠️ Command not yet implemented")

	for _, arg := range ctx.Args() {
		logger.SugaredLogger.Infof("Argument: %s", arg)
	}

	return nil
}

func setGlobal(ctx *cli.Context) error {
	logger.Logger.Warn("⚠️ Command not yet implemented")
	return nil
}
