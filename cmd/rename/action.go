package rename

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	stopWalkingErrMsg = "STOP WALKING"
)

func rename(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing RENAME command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	logger.SugaredLogger.Debugf("🐛 Kubernetes configuration file path: '%s'", kCfgFilePath)

	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations path")
	singleKonfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKonfigsPath)

	logger.Logger.Debug("🐛 Get context to rename")
	contextToRename, newContextName, ctxErr := getContextInfo(ctx)
	if ctxErr != nil {
		return ctxErr
	}

	kCfgErr := renameInKubeConfig(kCfgFilePath, contextToRename, newContextName)
	if kCfgErr != nil {
		return kCfgErr
	}

	kfgsErr := renameInKubeKonfigs(singleKonfigsPath, contextToRename, newContextName)
	if kfgsErr != nil {
		return kfgsErr
	}

	logger.SugaredLogger.Infof("✅  Context '%s' renamed to '%s'", contextToRename, newContextName)
	logger.Logger.Info("")
	return nil
}

func getContextInfo(ctx *cli.Context) (string, string, error) {
	logger.Logger.Debug("🐛 Get Kubernetes context to rename")
	args := ctx.Args()
	if args.Len() == 0 {
		return "", "", cli.Exit(
			"❌  Error getting Kubernetes context to rename: 'context to rename' and 'new context name' arguments not specified",
			51)
	}
	if strings.Compare(args.Get(0), "") == 0 {
		return "", "", cli.Exit(
			"❌  Error getting Kubernetes context to rename: 'context to rename' argument not specified",
			52)
	}
	if strings.Compare(args.Get(1), "") == 0 {
		return "", "", cli.Exit(
			"❌  Error getting Kubernetes context to rename: 'new context name' argument not specified",
			53)
	}

	return args.Get(0), args.Get(1), nil
}

func renameInKubeConfig(kCfgFilePath string, contextToRename, newContextName string) error {
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kCfg := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	checkCtxErr := kubeconfig.CheckIfContextExist(kCfg, contextToRename)
	if checkCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of context '%s' in Kubernetes configuration '%s': %s",
				contextToRename, kCfgFilePath, checkCtxErr.Error()), 34)
	}

	renameErr := renameContext(kCfg, kCfgFilePath, contextToRename, newContextName)
	if renameErr != nil {
		return renameErr
	}

	newValErr := kubeconfig.Validate(kCfg)
	if newValErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s", kCfgFilePath, newValErr.Error()),
			12)
	}

	newWriteErr := kubeconfig.Write(kCfg, kCfgFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error writing Kubernetes configuration '%s' to file: %s", kCfgFilePath, newWriteErr.Error()),
			13)
	}

	logger.SugaredLogger.Infof("✅  Context renamed from '%s' to '%s' in Kubernetes configuration '%s'",
		contextToRename, newContextName, kCfgFilePath)
	return nil
}

func renameContext(kCfg *clientcmdapi.Config, kCfgFilePath, contextToRename, newContextName string) error {
	logger.SugaredLogger.Infof("🔀 Renaming Kubernetes context '%s' as '%s'", contextToRename, newContextName)

	if strings.Compare(kCfg.CurrentContext, contextToRename) == 0 {
		logger.SugaredLogger.Debugf("🐛 Set new current context '%s' in Kubernetes configuration '%s'", newContextName, kCfgFilePath)
		kCfg.CurrentContext = newContextName
	}

	logger.SugaredLogger.Debugf("🐛 Remove context '%s' from Kubernetes configuration '%s'", contextToRename, kCfgFilePath)
	ctx := kCfg.Contexts[contextToRename]
	newContexts, remCtxErr := kubeconfig.RemoveContext(kCfg.Contexts, contextToRename)
	if remCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error removing context '%s' from Kubernetes configuration '%s': %s", contextToRename, kCfgFilePath, remCtxErr.Error()),
			54)
	}

	logger.SugaredLogger.Debugf("🐛 Insert context '%s' in Kubernetes configuration '%s'", newContextName, kCfgFilePath)
	newContexts[newContextName] = ctx
	kCfg.Contexts = newContexts

	return nil
}

func renameInKubeKonfigs(singleKonfigsPath, contextToRename, newContextName string) error {
	checkErr := utils.CheckIfFolderExist(singleKonfigsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warnf("⚠️  Single Kubernetes konfigurations path not found ('%s')", singleKonfigsPath)
		logger.Logger.Warn("ℹ️  Tip: run 'konf split' before anything else")
	} else {
		var cfgToRename string
		walkErr := filepath.Walk(
			singleKonfigsPath,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}

				if strings.HasPrefix(info.Name(), ".") {
					return nil
				}

				kCfg := kubeconfig.Load(path)
				// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file
				for name := range kCfg.Contexts {
					if name == contextToRename {
						cfgToRename = path
						return errors.New(stopWalkingErrMsg)
					}
				}

				return nil
			},
		)
		if walkErr != nil {
			if walkErr.Error() != stopWalkingErrMsg {
				return cli.Exit(
					fmt.Sprintf("❌  Error removing single Kubernetes konfigurations from '%s': %s",
						singleKonfigsPath, walkErr.Error()), 43)
			}
		}

		logger.SugaredLogger.Infof("🔀 Renaming context '%s' to '%s' in single Kubernetes konfigurations", contextToRename, newContextName)
		renErr := os.Rename(
			cfgToRename,
			strings.Replace(cfgToRename, contextToRename, newContextName, 1),
		)
		if renErr != nil {
			return cli.Exit(
				fmt.Sprintf("❌  Error renaming '%s' single Kubernetes konfigurations from '%s': %s",
					cfgToRename, singleKonfigsPath, renErr.Error()), 43)
		}

		// TODO rename also content of file!

		logger.SugaredLogger.Infof("✅  Context '%s' renamed to '%s' in single Kubernetes konfigurations '%s'",
			contextToRename, newContextName, singleKonfigsPath)
	}

	return nil
}
