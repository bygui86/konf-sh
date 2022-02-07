package rename

import (
	"fmt"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/urfave/cli/v2"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func rename(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing RENAME-CONTEXT command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	contextInfo, ctxErr := getContextInfo(ctx)
	if ctxErr != nil {
		return ctxErr
	}
	contextToRename := contextInfo.Get(0)
	newContextName := contextInfo.Get(1)

	logger.SugaredLogger.Debugf("🐛 Check existence of context '%s' in Kubernetes configuration '%s'", contextToRename, kubeConfigFilePath)
	checkCtxErr := kubeconfig.CheckIfContextExist(kubeConfig, contextToRename)
	if checkCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error checking existence of context '%s' in Kubernetes configuration '%s': %s", contextToRename, kubeConfigFilePath, checkCtxErr.Error()),
			34)
	}

	renameErr := renameContext(kubeConfig, kubeConfigFilePath, contextToRename, newContextName)
	if renameErr != nil {
		return renameErr
	}

	valWrErr := commons.ValidateAndWrite(kubeConfig, kubeConfigFilePath)
	if valWrErr != nil {
		return valWrErr
	}

	logger.SugaredLogger.Infof("✅ Completed! Context successfully renamed from '%s' to '%s' in Kubernetes configuration '%s'",
		contextToRename, newContextName, kubeConfigFilePath)
	logger.Logger.Info("")
	return nil
}

func getContextInfo(ctx *cli.Context) (cli.Args, error) {
	logger.Logger.Debug("🐛 Get Kubernetes context to rename")
	args := ctx.Args()
	if args.Len() == 0 {
		return nil, cli.Exit(
			"❌ Error getting Kubernetes context to rename: 'context to rename' and 'new context name' arguments not specified",
			51)
	}
	if strings.Compare(args.Get(0), "") == 0 {
		return nil, cli.Exit(
			"❌ Error getting Kubernetes context to rename: 'context to rename' argument not specified",
			52)
	}
	if strings.Compare(args.Get(1), "") == 0 {
		return nil, cli.Exit(
			"❌ Error getting Kubernetes context to rename: 'new context name' argument not specified",
			53)
	}

	return args, nil
}

func renameContext(kubeConfig *clientcmdapi.Config, kubeConfigFilePath, contextToRename, newContextName string) error {
	logger.SugaredLogger.Infof("🔀 Renaming Kubernetes context '%s' as '%s'", contextToRename, newContextName)

	if strings.Compare(kubeConfig.CurrentContext, contextToRename) == 0 {
		logger.SugaredLogger.Debugf("🐛 Set new current context '%s' in Kubernetes configuration '%s'", newContextName, kubeConfigFilePath)
		kubeConfig.CurrentContext = newContextName
	}

	logger.SugaredLogger.Debugf("🐛 Remove context '%s' from Kubernetes configuration '%s'", contextToRename, kubeConfigFilePath)
	ctx := kubeConfig.Contexts[contextToRename]
	newContexts, remCtxErr := kubeconfig.RemoveContext(kubeConfig.Contexts, contextToRename)
	if remCtxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌ Error removing context '%s' from Kubernetes configuration '%s': %s", contextToRename, kubeConfigFilePath, remCtxErr.Error()),
			54)
	}

	logger.SugaredLogger.Debugf("🐛 Insert context '%s' in Kubernetes configuration '%s'", newContextName, kubeConfigFilePath)
	newContexts[newContextName] = ctx
	kubeConfig.Contexts = newContexts

	return nil
}
