package rename

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func rename(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing RENAME command")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Debugf("🐛 Kubernetes configuration file path: '%s'", kCfgFilePath)

	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKonfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	zap.S().Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKonfigsPath)

	zap.L().Debug("🐛 Get context to rename")
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

	zap.S().Infof("✅  Context '%s' renamed to '%s'", contextToRename, newContextName)
	zap.L().Info("")
	return nil
}

func getContextInfo(ctx *cli.Context) (string, string, error) {
	zap.L().Debug("🐛 Get Kubernetes context to rename")
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
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
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
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s",
				kCfgFilePath, newValErr.Error()), 12)
	}

	newWriteErr := kubeconfig.Write(kCfg, kCfgFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error writing Kubernetes configuration '%s' to file: %s",
				kCfgFilePath, newWriteErr.Error()), 13)
	}

	zap.S().Infof("✅  Context renamed from '%s' to '%s' in Kubernetes configuration '%s'",
		contextToRename, newContextName, kCfgFilePath)
	return nil
}

func renameContext(kCfg *clientcmdapi.Config, kCfgFilePath, contextToRename, newContextName string) error {
	zap.S().Infof("🔀 Renaming Kubernetes context '%s' as '%s'", contextToRename, newContextName)

	if strings.Compare(kCfg.CurrentContext, contextToRename) == 0 {
		zap.S().Debugf("🐛 Set new current context '%s' in Kubernetes configuration '%s'", newContextName, kCfgFilePath)
		kCfg.CurrentContext = newContextName
	}

	zap.S().Debugf("🐛 Remove context '%s' from Kubernetes configuration '%s'", contextToRename, kCfgFilePath)
	ctx := kCfg.Contexts[contextToRename]
	newContexts, remCtxErr := kubeconfig.RemoveContext(kCfg.Contexts, contextToRename)
	if remCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error removing context '%s' from Kubernetes configuration '%s': %s",
				contextToRename, kCfgFilePath, remCtxErr.Error()), 54)
	}

	zap.S().Debugf("🐛 Insert context '%s' in Kubernetes configuration '%s'", newContextName, kCfgFilePath)
	newContexts[newContextName] = ctx
	kCfg.Contexts = newContexts

	return nil
}

func renameInKubeKonfigs(singleKonfigsPath, contextToRename, newContextName string) error {
	checkErr := commons.CheckIfFolderExist(singleKonfigsPath, false)
	if checkErr != nil {
		zap.S().Warnf("⚠️  Single Kubernetes konfigurations path not found ('%s')", singleKonfigsPath)
		zap.L().Warn("ℹ️  Tip: run 'konf split' before anything else")
	} else {
		stopWalkingErrMsg := "STOP WALKING"
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
				ctxFound := false
				for name := range kCfg.Contexts {
					if name == contextToRename {
						ctxFound = true
					}
				}

				if ctxFound {
					newPath := strings.Replace(path, contextToRename, newContextName, 1)
					zap.S().Infof("🔀 Renaming context '%s' to '%s' in single Kubernetes konfigurations",
						contextToRename, newContextName)
					renErr := os.Rename(path, newPath)
					if renErr != nil {
						return cli.Exit(
							fmt.Sprintf("❌  Error renaming '%s' single Kubernetes konfigurations from '%s': %s",
								contextToRename, singleKonfigsPath, renErr.Error()), 43)
					}

					currentCtx := kCfg.Contexts[contextToRename]
					delete(kCfg.Contexts, contextToRename)
					kCfg.Contexts[newContextName] = currentCtx
					kCfg.CurrentContext = newContextName

					newValErr := kubeconfig.Validate(kCfg)
					if newValErr != nil {
						return cli.Exit(
							fmt.Sprintf("❌  Error validating Kubernetes konfiguration from '%s': %s",
								newPath, newValErr.Error()), 12)
					}

					newWriteErr := kubeconfig.Write(kCfg, newPath)
					if newWriteErr != nil {
						return cli.Exit(
							fmt.Sprintf("❌  Error writing Kubernetes konfiguration '%s' to file: %s",
								newPath, newWriteErr.Error()), 13)
					}

					return errors.New(stopWalkingErrMsg)
				}

				return nil
			},
		)
		if walkErr != nil {
			if walkErr.Error() != stopWalkingErrMsg {
				return cli.Exit(
					fmt.Sprintf("❌  Error renaming single Kubernetes konfigurations from '%s': %s",
						singleKonfigsPath, walkErr.Error()), 43)
			}
		}

		zap.S().Infof("✅  Context '%s' renamed to '%s' in single Kubernetes konfigurations '%s'",
			contextToRename, newContextName, singleKonfigsPath)
	}

	return nil
}
