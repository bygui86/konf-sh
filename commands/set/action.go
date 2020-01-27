package set

import (
	"fmt"
	"path/filepath"

	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/logger"
	"github.com/bygui86/konf/utils"

	"github.com/urfave/cli"
)

// INFO: it seems that is not possible to run a command like "source ./set-local-script.sh" :(
func setLocal(ctx *cli.Context) error {
	logger.Logger.Debug("🐛 Executing SET-LOCAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get single Kubernetes configurations files path")
	singleConfigsPath := ctx.String(commons.SingleConfigsFlagName)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations files path '%s'", singleConfigsPath)
	checkFolderErr := utils.CheckIfFolderExist(singleConfigsPath, true)
	if checkFolderErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of Kubernetes configurations files path '%s': %s", singleConfigsPath, checkFolderErr.Error()),
			31)
	}
	logger.SugaredLogger.Debugf("📚 Single Kubernetes configurations files path: '%s'", singleConfigsPath)

	logger.Logger.Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if len(args) == 0 || args[0] == "" {
		return cli.NewExitError(
			"❌ Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args[0]

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes configurations file for context '%s'", context)
	localKubeConfig := filepath.Join(singleConfigsPath, context)
	checkFileErr := utils.CheckIfFileExist(localKubeConfig)
	if checkFileErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of Kubernetes context '%s' configuration file: %s", localKubeConfig, checkFileErr.Error()),
			33)
	}
	logger.SugaredLogger.Debugf("🧩 Selected Kubernetes context: '%s'", context)

	logger.Logger.Info(fmt.Sprintf("export %s='%s'", commons.KubeConfigEnvVar, localKubeConfig))
	return nil
}

func setGlobal(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing SET-GLOBAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.Logger.Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if len(args) == 0 || args[0] == "" {
		return cli.NewExitError(
			"❌ Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args[0]

	logger.SugaredLogger.Debugf("🐛 Check context '%s' existence in Kubernetes configuration '%s'", context, kubeConfigFilePath)
	checkCtxErr := kubeconfig.CheckIfContextExist(kubeConfig, context)
	if checkCtxErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error checking existence of context '%s' in Kubernetes configuration '%s': %s", context, kubeConfigFilePath, checkCtxErr.Error()),
			34)
	}
	logger.SugaredLogger.Infof("🧩 Selected Kubernetes context: '%s'", context)

	logger.SugaredLogger.Debugf("🐛 Set new context '%s' in Kubernetes configuration '%s'", context, kubeConfigFilePath)
	kubeConfig.CurrentContext = context

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration '%s'", kubeConfigFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error validating Kubernetes configuration '%s': %s", kubeConfigFilePath, valErr.Error()),
			12)
	}

	logger.SugaredLogger.Debugf("🐛 Write Kubernetes configuration '%s' to file", kubeConfigFilePath)
	writeErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if writeErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error writing Kubernetes configuration '%s' to file: %s", kubeConfigFilePath, writeErr.Error()),
			13)
	}

	logger.SugaredLogger.Infof("✅ Completed! Kubernete global configuration '%s' successfully updated with current context '%s'", kubeConfigFilePath, context)
	logger.Logger.Info("")
	return nil
}
