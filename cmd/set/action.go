package set

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

// INFO: it seems that is not possible to run a command like "source ./set-local-script.sh" :(
func setLocal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing SET-LOCAL command")
	zap.L().Debug("")

	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleConfigsPath := ctx.String(commons.SingleKonfigsFlagName)

	zap.S().Debugf("🐛 Check existence of single Kubernetes konfigurations path '%s'", singleConfigsPath)
	checkFolderErr := commons.CheckIfFolderExist(singleConfigsPath, true)
	if checkFolderErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes konfigurations path '%s': %s", singleConfigsPath, checkFolderErr.Error()),
			31)
	}
	zap.S().Debugf("📚 Single Kubernetes konfigurations path: '%s'", singleConfigsPath)

	zap.L().Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit(
			"❌  Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args.Get(0)

	zap.S().Debugf("🐛 Check existence of single Kubernetes konfigurations file for context '%s'", context)
	localKubeConfig := filepath.Join(singleConfigsPath, context)
	checkFileErr := commons.CheckIfFileExist(localKubeConfig)
	if checkFileErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes context '%s' configuration file: %s", localKubeConfig, checkFileErr.Error()),
			33)
	}
	zap.S().Debugf("🧩 Selected Kubernetes context: '%s'", context)

	zap.L().Info(fmt.Sprintf("export %s='%s'", commons.KubeConfigEnvVar, localKubeConfig))
	return nil
}

func setGlobal(ctx *cli.Context) error {
	zap.L().Info("")
	zap.L().Debug("🐛 Executing SET-GLOBAL command")
	zap.L().Debug("")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.L().Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit(
			"❌  Error getting Kubernetes context: context argument not specified",
			32)
	}
	context := args.Get(0)

	zap.S().Debugf("🐛 Check existence of context '%s' in Kubernetes configuration '%s'", context, kubeConfigFilePath)
	checkCtxErr := kubeconfig.CheckIfContextExist(kubeConfig, context)
	if checkCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of context '%s' in Kubernetes configuration '%s': %s", context, kubeConfigFilePath, checkCtxErr.Error()),
			34)
	}
	zap.S().Infof("🧩 Selected Kubernetes context: '%s'", context)

	zap.S().Debugf("🐛 Set new context '%s' in Kubernetes configuration '%s'", context, kubeConfigFilePath)
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

	zap.S().Infof("✅ Completed! Kubernetes global configuration '%s' successfully updated with current context '%s'", kubeConfigFilePath, context)
	zap.L().Info("")
	return nil
}
