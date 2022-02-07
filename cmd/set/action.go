package set

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

// INFO: it seems that is not possible to run a command like "source ./set-local-script.sh" :(
func setLocal(ctx *cli.Context) error {
	logger.Logger.Debug("🐛 Executing SET-LOCAL command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations path")
	singleConfigsPath := ctx.String(commons.SingleKonfigsFlagName)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes konfigurations path '%s'", singleConfigsPath)
	checkFolderErr := utils.CheckIfFolderExist(singleConfigsPath, true)
	if checkFolderErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes konfigurations path '%s': %s", singleConfigsPath, checkFolderErr.Error()),
			31)
	}
	logger.SugaredLogger.Debugf("📚 Single Kubernetes konfigurations path: '%s'", singleConfigsPath)

	logger.Logger.Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit(
			"❌  Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args.Get(0)

	logger.SugaredLogger.Debugf("🐛 Check existence of single Kubernetes konfigurations file for context '%s'", context)
	localKubeConfig := filepath.Join(singleConfigsPath, context)
	checkFileErr := utils.CheckIfFileExist(localKubeConfig)
	if checkFileErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes context '%s' configuration file: %s", localKubeConfig, checkFileErr.Error()),
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
	kubeConfigFilePath := ctx.String(commons.KubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.Logger.Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit(
			"❌  Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args.Get(0)

	logger.SugaredLogger.Debugf("🐛 Check existence of context '%s' in Kubernetes configuration '%s'", context, kubeConfigFilePath)
	checkCtxErr := kubeconfig.CheckIfContextExist(kubeConfig, context)
	if checkCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of context '%s' in Kubernetes configuration '%s': %s", context, kubeConfigFilePath, checkCtxErr.Error()),
			34)
	}
	logger.SugaredLogger.Infof("🧩 Selected Kubernetes context: '%s'", context)

	logger.SugaredLogger.Debugf("🐛 Set new context '%s' in Kubernetes configuration '%s'", context, kubeConfigFilePath)
	kubeConfig.CurrentContext = context

	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s", kubeConfigFilePath, newValErr.Error()),
			12)
	}

	newWriteErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error writing Kubernetes configuration '%s' to file: %s", kubeConfigFilePath, newWriteErr.Error()),
			13)
	}

	logger.SugaredLogger.Infof("✅ Completed! Kubernetes global configuration '%s' successfully updated with current context '%s'", kubeConfigFilePath, context)
	logger.Logger.Info("")
	return nil
}
