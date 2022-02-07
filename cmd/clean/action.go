package clean

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/urfave/cli/v2"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func clean(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing CLEAN-CONTEXT command")
	logger.Logger.Debug("")

	contextSlice, ctxErr := getContextList(ctx)
	if ctxErr != nil {
		return ctxErr
	}

	logger.Logger.Debug("🐛 Ask for user confirmation to clean context list")
	if userDeletionConfirm() {
		kubeConfigFilePath, cleanErr := cleanInternal(ctx, contextSlice)
		if cleanErr != nil {
			return cleanErr
		}
		logger.SugaredLogger.Infof("✅ Completed! Context list '%s' removed from Kubernetes configuration '%s'",
			strings.Join(contextSlice, ", "), kubeConfigFilePath)
		logger.Logger.Info("")

	} else {
		logger.Logger.Info("❌ User didn't confirm to proceed, aborting...")
		logger.Logger.Info("")
	}

	return nil
}

func getContextList(ctx *cli.Context) ([]string, error) {
	logger.Logger.Debug("🐛 Get Kubernetes context list to clean")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return nil, cli.Exit(
			"❌ Error getting Kubernetes context list: 'context list' argument not specified",
			41)
	}

	contextList := args.Get(0)
	contextSlice, ctxValErr := validateContextListArgument(contextList)
	if ctxValErr != nil {
		return nil, cli.Exit(
			"❌ Error validating Kubernetes context list: 'context list' argument not valid. Context list must be a comma-separated list.",
			42)
	}

	logger.SugaredLogger.Infof("📋 Context list to clean: '%s'", strings.Join(contextSlice, ", "))
	return contextSlice, nil
}

func validateContextListArgument(contextList string) ([]string, error) {
	contextSlice := strings.Split(contextList, ",")
	if len(contextSlice) > 0 {
		return contextSlice, nil
	}
	return nil, errors.New("error validating context list")
}

func userDeletionConfirm() bool {
	reader := bufio.NewReader(os.Stdin)
	logger.Logger.Warn("⚠️  Deleting a context from Kubernetes configuration can't be undone")
	logger.Logger.Info("❓ are you sure you want to proceed? [y | n]")
	for {
		confirm, _ := reader.ReadString('\n')
		confirm = strings.Replace(confirm, "\n", "", -1) // convert CRLF to LF
		logger.SugaredLogger.Debugf("🐛 User confirmation answer: %s", confirm)
		if strings.Compare(confirm, "y") == 0 {
			return true
		} else if strings.Compare(confirm, "n") == 0 {
			break
		} else {
			logger.Logger.Error("❌ Wrong input, please answer with 'y' or 'n'")
		}
	}
	return false
}

func cleanInternal(ctx *cli.Context, contextSlice []string) (string, error) {
	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kubeConfigFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return "", cli.Exit(
			fmt.Sprintf("❌ Error validating Kubernetes configuration from '%s': %s", kubeConfigFilePath, valErr.Error()),
			12)
	}

	logger.Logger.Info("🧹 Removing selected contexts from Kubernetes configuration")
	cleanErr := cleanContextList(kubeConfig, contextSlice)
	if cleanErr != nil {
		return "", cli.Exit(
			fmt.Sprintf("❌ Error cleaning Kubernetes context list: %s", cleanErr.Error()),
			43)
	}

	valWrErr := commons.ValidateAndWrite(kubeConfig, kubeConfigFilePath)
	if valWrErr != nil {
		return "", valWrErr
	}

	return kubeConfigFilePath, nil
}

func cleanContextList(kubeConfig *clientcmdapi.Config, contextSlice []string) error {
	for _, rmCtx := range contextSlice {
		err := kubeconfig.CheckIfContextExist(kubeConfig, rmCtx)
		if err != nil {
			logger.SugaredLogger.Infof("❓ Context '%s' to clean not found, skipping...", rmCtx)
			continue
		}

		logger.SugaredLogger.Debugf("🐛 Removing context '%s' from Kubernetes configuration", rmCtx)
		tempCtx := kubeConfig.Contexts[rmCtx]
		ctxMap, ctxErr := kubeconfig.RemoveContext(kubeConfig.Contexts, rmCtx)
		if ctxErr != nil {
			return ctxErr
		}
		kubeConfig.Contexts = ctxMap

		rmCluster := tempCtx.Cluster
		logger.SugaredLogger.Debugf("🐛 Removing cluster '%s' from Kubernetes configuration", rmCluster)
		clMap, clErr := kubeconfig.RemoveCluster(kubeConfig.Clusters, rmCluster)
		if clErr != nil {
			return clErr
		}
		kubeConfig.Clusters = clMap

		rmAuthInfo := tempCtx.AuthInfo // user
		logger.SugaredLogger.Debugf("🐛 Removing user '%s' from Kubernetes configuration", rmAuthInfo)
		authMap, authErr := kubeconfig.RemoveAuthInfo(kubeConfig.AuthInfos, rmAuthInfo)
		if authErr != nil {
			return authErr
		}
		kubeConfig.AuthInfos = authMap
	}

	kubeConfig.CurrentContext = ""
	return nil
}
