package delete

import (
	"bufio"
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

// INFO: deleteCtx cannot be named delete because of name collision
func deleteCtx(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing DELETE command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	logger.SugaredLogger.Debugf("🐛 Kubernetes configuration file path: '%s'", kCfgFilePath)

	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations path")
	singleKonfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKonfigsPath)

	logger.Logger.Debug("🐛 Get contexts to delete")
	contextSlice, ctxErr := getContextList(ctx)
	if ctxErr != nil {
		return ctxErr
	}
	logger.SugaredLogger.Infof("📋 Contexts to delete: '%s'", strings.Join(contextSlice, ", "))

	logger.Logger.Debug("🐛 Ask for user confirmation to delete contexts")
	if userDeletionConfirm() {
		kCfgErr := deleteFromKubeConfig(kCfgFilePath, contextSlice)
		if kCfgErr != nil {
			return kCfgErr
		}

		kfgsErr := deleteFromKubeKonfigs(singleKonfigsPath, contextSlice)
		if kfgsErr != nil {
			return kfgsErr
		}

		logger.SugaredLogger.Infof("✅  Removing contexts '%s' completed", strings.Join(contextSlice, ", "))

		logger.Logger.Info("")

	} else {
		logger.Logger.Info("❌  User didn't confirm to proceed, aborting...")
		logger.Logger.Info("")
	}

	return nil
}

func getContextList(ctx *cli.Context) ([]string, error) {
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return nil, cli.Exit(
			"❌  Error getting Kubernetes contexts: argument not specified",
			41)
	}

	contextList := args.Get(0)
	contextSlice, ctxValErr := validateContextListArgument(contextList)
	if ctxValErr != nil {
		return nil, cli.Exit(
			"❌  Error validating Kubernetes contexts: argument not valid, it must be a comma-separated list.",
			42)
	}

	return contextSlice, nil
}

func validateContextListArgument(contextList string) ([]string, error) {
	contextSlice := strings.Split(contextList, ",")
	if len(contextSlice) > 0 {
		return contextSlice, nil
	}
	return nil, errors.New("error validating contexts")
}

func userDeletionConfirm() bool {
	reader := bufio.NewReader(os.Stdin)
	logger.Logger.Warn("🚨 Deleting a context from Kubernetes configuration can't be undone")
	logger.Logger.Info("❓  are you sure you want to proceed? [y | n]")
	for {
		confirm, _ := reader.ReadString('\n')
		confirm = strings.Replace(confirm, "\n", "", -1) // convert CRLF to LF
		logger.SugaredLogger.Debugf("🐛 User confirmation answer: %s", confirm)
		if strings.Compare(confirm, "y") == 0 {
			logger.Logger.Info("")
			return true
		} else if strings.Compare(confirm, "n") == 0 {
			break
		} else {
			logger.Logger.Error("❌  Wrong input, please answer with 'y' or 'n'")
		}
	}
	return false
}

func deleteFromKubeConfig(kCfgFilePath string, contextSlice []string) error {
	logger.SugaredLogger.Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kubeConfig := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	logger.SugaredLogger.Debugf("🐛 Validate Kubernetes configuration from '%s'", kCfgFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s", kCfgFilePath, valErr.Error()),
			12)
	}

	logger.Logger.Info("🧹 Removing selected context from Kubernetes configuration")
	cleanErr := deleteContextList(kubeConfig, contextSlice)
	if cleanErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error cleaning Kubernetes contexts: %s", cleanErr.Error()),
			43)
	}

	newValErr := kubeconfig.Validate(kubeConfig)
	if newValErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s", kCfgFilePath, newValErr.Error()),
			12)
	}

	newWriteErr := kubeconfig.Write(kubeConfig, kCfgFilePath)
	if newWriteErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error writing Kubernetes configuration '%s' to file: %s", kCfgFilePath, newWriteErr.Error()),
			13)
	}

	logger.SugaredLogger.Infof("✅  Contexts '%s' removed from Kubernetes configuration '%s'",
		strings.Join(contextSlice, ", "), kCfgFilePath)
	return nil
}

func deleteContextList(kCfg *clientcmdapi.Config, contextSlice []string) error {
	for _, rmCtx := range contextSlice {
		checkErr := kubeconfig.CheckIfContextExist(kCfg, rmCtx)
		if checkErr != nil {
			logger.SugaredLogger.Debugf("🐛 Context '%s' to delete not found, skipping...", rmCtx)
			continue
		}

		logger.SugaredLogger.Debugf("🐛 Removing context '%s' from Kubernetes configuration", rmCtx)
		tempCtx := kCfg.Contexts[rmCtx]
		ctxMap, ctxErr := kubeconfig.RemoveContext(kCfg.Contexts, rmCtx)
		if ctxErr != nil {
			return ctxErr
		}
		kCfg.Contexts = ctxMap

		rmCluster := tempCtx.Cluster
		logger.SugaredLogger.Debugf("🐛 Removing cluster '%s' from Kubernetes configuration", rmCluster)
		clMap, clErr := kubeconfig.RemoveCluster(kCfg.Clusters, rmCluster)
		if clErr != nil {
			return clErr
		}
		kCfg.Clusters = clMap

		rmAuthInfo := tempCtx.AuthInfo // user
		logger.SugaredLogger.Debugf("🐛 Removing user '%s' from Kubernetes configuration", rmAuthInfo)
		authMap, authErr := kubeconfig.RemoveAuthInfo(kCfg.AuthInfos, rmAuthInfo)
		if authErr != nil {
			return authErr
		}
		kCfg.AuthInfos = authMap

		if rmCtx == kCfg.CurrentContext {
			kCfg.CurrentContext = ""
		}
	}

	return nil
}

func deleteFromKubeKonfigs(singleKonfigsPath string, contextSlice []string) error {
	checkErr := utils.CheckIfFolderExist(singleKonfigsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warnf("🚨 Single Kubernetes konfigurations path not found ('%s')", singleKonfigsPath)
		logger.Logger.Warn("💬 Tip: run 'konf split' before anything else")
	} else {
		cfgToDel := make([]string, 0)
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
					for _, ctxDel := range contextSlice {
						if name == ctxDel {
							cfgToDel = append(cfgToDel, path)
						}
					}
				}

				return nil
			},
		)
		if walkErr != nil {
			return cli.Exit(
				fmt.Sprintf("❌  Error removing single Kubernetes konfigurations from '%s': %s",
					singleKonfigsPath, walkErr.Error()), 43)
		}

		logger.Logger.Info("🧹 Removing selected contexts from single Kubernetes konfigurations")
		for _, delPath := range cfgToDel {
			delErr := os.Remove(delPath)
			if delErr != nil {
				return cli.Exit(
					fmt.Sprintf("❌  Error removing '%s' single Kubernetes konfigurations from '%s': %s",
						delPath, singleKonfigsPath, delErr.Error()), 43)
			}
		}

		logger.SugaredLogger.Infof("✅  Contexts '%s' removed from single Kubernetes konfigurations '%s'",
			strings.Join(contextSlice, ", "), singleKonfigsPath)
	}

	return nil
}
