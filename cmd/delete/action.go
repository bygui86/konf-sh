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
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// INFO: deleteCtx cannot be named delete because of name collision
func deleteCtx(ctx *cli.Context) error {
	zap.L().Info("")
	zap.L().Debug("🐛 Executing DELETE command")
	zap.L().Debug("")

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Debugf("🐛 Kubernetes configuration file path: '%s'", kCfgFilePath)

	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKonfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	zap.S().Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKonfigsPath)

	zap.L().Debug("🐛 Get contexts to delete")
	contextSlice, ctxErr := getContextList(ctx)
	if ctxErr != nil {
		return ctxErr
	}
	zap.S().Infof("📋 Contexts to delete: '%s'", strings.Join(contextSlice, ", "))

	zap.L().Debug("🐛 Ask for user confirmation to delete contexts")
	if userDeletionConfirm() {
		kCfgErr := deleteFromKubeConfig(kCfgFilePath, contextSlice)
		if kCfgErr != nil {
			return kCfgErr
		}

		kfgsErr := deleteFromKubeKonfigs(singleKonfigsPath, contextSlice)
		if kfgsErr != nil {
			return kfgsErr
		}

		zap.S().Infof("✅  Removing contexts '%s' completed", strings.Join(contextSlice, ", "))

		zap.L().Info("")

	} else {
		zap.L().Info("❌  User didn't confirm to proceed, aborting...")
		zap.L().Info("")
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
	zap.L().Warn("🚨 Deleting a context from Kubernetes configuration can't be undone")
	zap.L().Info("❓  are you sure you want to proceed? [y | n]")
	for {
		confirm, _ := reader.ReadString('\n')
		confirm = strings.Replace(confirm, "\n", "", -1) // convert CRLF to LF
		zap.S().Debugf("🐛 User confirmation answer: %s", confirm)
		if strings.Compare(confirm, "y") == 0 {
			zap.L().Info("")
			return true
		} else if strings.Compare(confirm, "n") == 0 {
			break
		} else {
			zap.L().Error("❌  Wrong input, please answer with 'y' or 'n'")
		}
	}
	return false
}

func deleteFromKubeConfig(kCfgFilePath string, contextSlice []string) error {
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kubeConfig := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Debugf("🐛 Validate Kubernetes configuration from '%s'", kCfgFilePath)
	valErr := kubeconfig.Validate(kubeConfig)
	if valErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error validating Kubernetes configuration from '%s': %s", kCfgFilePath, valErr.Error()),
			12)
	}

	zap.L().Info("🧹 Removing selected context from Kubernetes configuration")
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

	zap.S().Infof("✅  Contexts '%s' removed from Kubernetes configuration '%s'",
		strings.Join(contextSlice, ", "), kCfgFilePath)
	return nil
}

func deleteContextList(kCfg *clientcmdapi.Config, contextSlice []string) error {
	for _, rmCtx := range contextSlice {
		checkErr := kubeconfig.CheckIfContextExist(kCfg, rmCtx)
		if checkErr != nil {
			zap.S().Debugf("🐛 Context '%s' to delete not found, skipping...", rmCtx)
			continue
		}

		zap.S().Debugf("🐛 Removing context '%s' from Kubernetes configuration", rmCtx)
		tempCtx := kCfg.Contexts[rmCtx]
		ctxMap, ctxErr := kubeconfig.RemoveContext(kCfg.Contexts, rmCtx)
		if ctxErr != nil {
			return ctxErr
		}
		kCfg.Contexts = ctxMap

		rmCluster := tempCtx.Cluster
		zap.S().Debugf("🐛 Removing cluster '%s' from Kubernetes configuration", rmCluster)
		clMap, clErr := kubeconfig.RemoveCluster(kCfg.Clusters, rmCluster)
		if clErr != nil {
			return clErr
		}
		kCfg.Clusters = clMap

		rmAuthInfo := tempCtx.AuthInfo // user
		zap.S().Debugf("🐛 Removing user '%s' from Kubernetes configuration", rmAuthInfo)
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
	checkErr := commons.CheckIfFolderExist(singleKonfigsPath, false)
	if checkErr != nil {
		zap.S().Warnf("🚨 Single Kubernetes konfigurations path not found ('%s')", singleKonfigsPath)
		zap.L().Warn("💬 Tip: run 'konf split' before anything else")
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

		zap.L().Info("🧹 Removing selected contexts from single Kubernetes konfigurations")
		for _, delPath := range cfgToDel {
			delErr := os.Remove(delPath)
			if delErr != nil {
				return cli.Exit(
					fmt.Sprintf("❌  Error removing '%s' single Kubernetes konfigurations from '%s': %s",
						delPath, singleKonfigsPath, delErr.Error()), 43)
			}
		}

		zap.S().Infof("✅  Contexts '%s' removed from single Kubernetes konfigurations '%s'",
			strings.Join(contextSlice, ", "), singleKonfigsPath)
	}

	return nil
}
