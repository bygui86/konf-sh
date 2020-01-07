package set

import (
	"fmt"
	"path/filepath"

	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/logger"
	"bygui86/konf/utils"

	"github.com/urfave/cli"
)

func setLocal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SET-LOCAL command")
	logger.Logger.Debug("")

	logger.Logger.Warn("⚠️  Work in progress")

	logger.Logger.Debug("🐛 Get single Kubernetes configurations files path")
	singleConfigsPath := ctx.String(commons.SingleConfigsFlagName)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations files path '%s'", singleConfigsPath)
	checkFolderErr := utils.CheckIfFolderExist(singleConfigsPath, true)
	if checkFolderErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of Kubernetes configurations files path '%s': %s", checkFolderErr.Error(), singleConfigsPath),
			31)
	}
	logger.SugaredLogger.Infof("📚 Single Kubernetes configurations files path: '%s'", singleConfigsPath)

	logger.Logger.Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if len(args) == 0 || args[0] == "" {
		return cli.NewExitError(
			"❌ Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args[0]

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations files path '%s'", singleConfigsPath)
	localKubeConfig := filepath.Join(singleConfigsPath, context)
	checkFileErr := utils.CheckIfFileExist(localKubeConfig)
	if checkFileErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of Kubernetes context '%s': %s", checkFileErr.Error(), localKubeConfig),
			33)
	}
	logger.SugaredLogger.Infof("🧩 Selected Kubernetes context: '%s'", context)

	logger.Logger.Debug("🐛 Set local Kubernetes context")
	setErr := kubeconfig.SetKubeConfigEnvVar(localKubeConfig)
	if setErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error setting local Kubernetes context (env-var KUBECONFIG): %s", setErr.Error()),
			34)
	}
	// logger.SugaredLogger.Infof("⚙️  Local Kubernetes context: '%s'", localKubeConfig)
	logger.SugaredLogger.Infof("⚙️  Local Kubernetes context: '%s'", kubeconfig.GetKubeConfigEnvVar())

	logger.Logger.Info("")
	return nil
}

func setGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SET-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Warn("⚠️  Command not yet implemented")
	return nil
}
