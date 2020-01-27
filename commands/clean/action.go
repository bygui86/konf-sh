package clean

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/bygui86/konf/commons"
	"github.com/bygui86/konf/kubeconfig"
	"github.com/bygui86/konf/logger"
)

func clean(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing CLEAN command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes context list to clean")
	args := ctx.Args()
	if len(args) == 0 || args[0] == "" {
		return cli.NewExitError(
			"❌ Error getting Kubernetes context list: context list argument not specified",
			41)
	}
	contextList := args[0]
	contextSlice, ctxValErr := validateContextList(contextList)
	if ctxValErr != nil {
		return cli.NewExitError(
			"❌ Error validating Kubernetes context list: context list argument not valid. Context list must be a comma-separated list.",
			42)
	}
	logger.SugaredLogger.Infof("📋 Context list to clean: %s", strings.Join(contextSlice, ", "))

	// TODO question to the user
	// Deleting contextList from Kubernetes configuration can't be undone, are you sure you want to proceed?

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kubeConfigFilePath := ctx.String(commons.CustomKubeConfigFlagName)
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kubeConfigFilePath)
	kubeConfig := kubeconfig.Load(kubeConfigFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kubeConfigFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error validating Kubernetes configuration from '%s': %s", kubeConfigFilePath, valErr.Error()),
			12)
	}

	logger.Logger.Info("🧹 Removing selected contexts from Kubernetes configuration")
	cleanErr := cleanContextList(kubeConfig, contextSlice)
	if cleanErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error cleaning Kubernetes context list: %s", cleanErr.Error()),
			43)
	}

	logger.SugaredLogger.Debugf("🐛 Validate cleaned Kubernetes configuration")
	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error validating cleaned Kubernetes configuration from '%s': %s", kubeConfigFilePath, newValErr.Error()),
			12)
	}

	logger.SugaredLogger.Debugf("🐛 Write cleaned Kubernetes configuration to file '%s'", kubeConfigFilePath)
	newWriteErr := kubeconfig.Write(kubeConfig, kubeConfigFilePath)
	if newWriteErr != nil {
		return cli.NewExitError(
			fmt.Sprintf("❌ Error writing cleaned Kubernetes configuration '%s' to file: %s", kubeConfigFilePath, newWriteErr.Error()),
			13)
	}

	logger.SugaredLogger.Infof("✅ Completed! Context list %s removed from Kubernetes configuration '%s'", strings.Join(contextSlice, ", "), kubeConfigFilePath)
	logger.Logger.Info("")
	return nil
}

func validateContextList(contextList string) ([]string, error) {
	contextSlice := strings.Split(contextList, ",")
	if len(contextSlice) > 0 {
		return contextSlice, nil
	}
	return nil, errors.New("error validating context list")
}

func cleanContextList(kubeConfig *clientcmdapi.Config, contextSlice []string) error {
	for _, rmCtx := range contextSlice {
		err := kubeconfig.CheckIfContextExist(kubeConfig, rmCtx)
		if err != nil {
			logger.SugaredLogger.Debugf("🐛 Context to clean %s not found, skipping...", rmCtx)
			continue
		}

		logger.SugaredLogger.Debugf("🐛 Removing context %s from Kubernetes configuration", rmCtx)
		tempCtx := kubeConfig.Contexts[rmCtx]
		ctxMap, ctxErr := kubeconfig.RemoveContext(kubeConfig.Contexts, rmCtx)
		if ctxErr != nil {
			return ctxErr
		}
		kubeConfig.Contexts = ctxMap

		rmCluster := tempCtx.Cluster
		logger.SugaredLogger.Debugf("🐛 Removing cluster %s from Kubernetes configuration", rmCluster)
		clMap, clErr := kubeconfig.RemoveCluster(kubeConfig.Clusters, rmCluster)
		if clErr != nil {
			return clErr
		}
		kubeConfig.Clusters = clMap

		rmAuthInfo := tempCtx.AuthInfo // user
		logger.SugaredLogger.Debugf("🐛 Removing user %s from Kubernetes configuration", rmAuthInfo)
		authMap, authErr := kubeconfig.RemoveAuthInfo(kubeConfig.AuthInfos, rmAuthInfo)
		if authErr != nil {
			return authErr
		}
		kubeConfig.AuthInfos = authMap
	}

	kubeConfig.CurrentContext = ""
	return nil
}
