package set

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func setLocal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing SET-LOCAL command")

	zap.L().Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit("❌  Error getting Kubernetes context: context argument not specified", 32)
	}
	context := args.Get(0)

	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKcfgsPath := ctx.String(commons.SingleKonfigsFlagName)

	if context == "-" {
		return setLastLocal(singleKcfgsPath)
	}

	return setSelectedLocal(singleKcfgsPath, context)
}

func setLastLocal(singleKcfgsPath string) error {
	zap.L().Debug("🐛 Set last Kubernetes context as local")

	return setSelectedLocal(
		singleKcfgsPath,
		readLastCtx(singleKcfgsPath, "local"),
	)
}

func setSelectedLocal(singleKcfgsPath, context string) error {
	zap.S().Debugf("🐛 Set selected Kubernetes context '%s' as local", context)

	zap.S().Debugf("🐛 Check existence of single Kubernetes konfigurations path '%s'", singleKcfgsPath)
	dirErr := commons.CheckIfFolderExist(singleKcfgsPath, true)
	if dirErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes konfigurations path '%s': %s",
				singleKcfgsPath, dirErr.Error()), 31)
	}
	zap.S().Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKcfgsPath)

	zap.S().Debugf("🐛 Check existence of single Kubernetes konfigurations file for context '%s'", context)
	localKCfg := filepath.Join(singleKcfgsPath, context)
	fileErr := commons.CheckIfFileExist(localKCfg)
	if fileErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of Kubernetes context '%s' configuration file: %s",
				localKCfg, fileErr.Error()), 33)
	}
	zap.S().Debugf("🐛 Selected Kubernetes context: '%s'", context)

	zap.S().Debugf("🐛 Get '%s' environment variable value", commons.KubeConfigEnvVar)
	lastCtxFull := kubeconfig.GetKubeConfigEnvVar()
	singleKcfgsPathTemp := strings.Replace(singleKcfgsPath, "./", "", 1)
	lastCtx := strings.Replace(lastCtxFull, fmt.Sprintf("%s/", singleKcfgsPathTemp), "", 1)
	zap.S().Debugf("🐛 Last Kubernetes context: '%s'", lastCtx)
	if lastCtx != "" {
		lastErr := saveLastCtx(
			singleKcfgsPath,
			lastCtx,
			"local",
		)
		if lastErr != nil {
			zap.S().Errorf("❌  Error saving last Kubernetes context '%s': %s - 'konf set local -' might not work",
				context, lastErr.Error())
		}
	}

	zap.L().Info(fmt.Sprintf("export %s='%s'", commons.KubeConfigEnvVar, localKCfg)) // TODO to be replaced by following line
	//zap.L().Info(fmt.Sprintf("%s", localKCfg)) // TODO enable when shell wrapper is available
	return nil
}

func setGlobal(ctx *cli.Context) error {
	zap.L().Debug("🐛 Executing SET-GLOBAL command")

	zap.L().Debug("🐛 Get selected Kubernetes context")
	args := ctx.Args()
	if args.Len() == 0 || strings.Compare(args.Get(0), "") == 0 {
		return cli.Exit("❌  Error getting Kubernetes context: context argument not specified", 32)
	}
	context := args.Get(0)

	zap.L().Debug("🐛 Get single Kubernetes konfigurations path")
	singleKcfgsPath := ctx.String(commons.SingleKonfigsFlagName)

	if context == "-" {
		return setLastGlobal(ctx, singleKcfgsPath)
	}

	return setSelectedGlobal(ctx, singleKcfgsPath, context)
}

func setLastGlobal(ctx *cli.Context, singleKcfgsPath string) error {
	zap.L().Debug("🐛 Set last Kubernetes context as global")

	lastCtx := readLastCtx(singleKcfgsPath, "global")
	if lastCtx == "" {
		return cli.Exit("❌  Error retrieving last Kubernetes context: no last global found", 35)
	}

	zap.S().Infof("⏮ Set last Kubernetes context '%s' as global", lastCtx)
	return setSelectedGlobal(ctx, singleKcfgsPath, lastCtx)
}

func setSelectedGlobal(ctx *cli.Context, singleKcfgsPath, context string) error {
	zap.S().Debugf("🐛 Set selected Kubernetes context '%s' as global", context)

	zap.L().Debug("🐛 Get Kubernetes configuration file path")
	kCfgFilePath := ctx.String(commons.KubeConfigFlagName)
	zap.S().Infof("📖 Load Kubernetes configuration from '%s'", kCfgFilePath)
	kCfg := kubeconfig.Load(kCfgFilePath)
	// INFO: no need to check if kubeConfig is nil, because the inner method called will exit if it does not find the configuration file

	zap.S().Debugf("🐛 Check existence of context '%s' in Kubernetes configuration '%s'", context, kCfgFilePath)
	ctxErr := kubeconfig.CheckIfContextExist(kCfg, context)
	if ctxErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of context '%s' in Kubernetes configuration '%s': %s",
				context, kCfgFilePath, ctxErr.Error()), 34)
	}
	zap.S().Infof("🧩 Selected Kubernetes context: '%s'", context)

	zap.S().Debugf("🐛 Set new context '%s' in Kubernetes configuration '%s'", context, kCfgFilePath)
	lastCtx := kCfg.CurrentContext
	kCfg.CurrentContext = context

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

	zap.S().Debugf("🐛 Last Kubernetes context: '%s'", lastCtx)
	if lastCtx != "" {
		lastErr := saveLastCtx(
			strings.Replace(singleKcfgsPath, "/config", "", 1),
			lastCtx,
			"global",
		)
		if lastErr != nil {
			zap.S().Errorf("❌  Error saving last Kubernetes context '%s': %s - 'konf set local -' might not work",
				lastCtx, lastErr.Error())
		}
	}

	zap.S().Infof("✅  Kubernetes global configuration '%s' successfully updated with current context '%s'", kCfgFilePath, context)
	zap.L().Info("")
	return nil
}

func saveLastCtx(singleKcfgsPath, context, command string) error {
	lastDirPath := fmt.Sprintf("%s/last-ctx", singleKcfgsPath)
	lastFilePath := fmt.Sprintf("%s/%s", lastDirPath, command)
	zap.S().Debugf("🐛 Saving last %s Kubernetes context '%s' to file '%s'", command, context, lastFilePath)

	zap.S().Debugf("🐛 Check existence of last Kubernetes context path '%s'", lastDirPath)
	checkErr := commons.CheckIfFolderExist(lastDirPath, true)
	if checkErr != nil {
		return cli.Exit(
			fmt.Sprintf("❌  Error checking existence of last Kubernetes context path '%s': %s",
				lastDirPath, checkErr.Error()), 14)
	}

	file, openErr := os.Create(lastFilePath)
	if openErr != nil {
		return openErr
	}
	defer file.Close()

	_, wrErr := file.WriteString(context)
	if wrErr != nil {
		return wrErr
	}

	return nil
}

func readLastCtx(singleKcfgsPath, command string) string {
	lastDirPath := fmt.Sprintf("%s/last-ctx", singleKcfgsPath)
	lastFilePath := fmt.Sprintf("%s/%s", lastDirPath, command)
	zap.S().Debugf("🐛 Reading last %s Kubernetes context from file '%s'", command, lastFilePath)

	zap.S().Debugf("🐛 Check existence of last Kubernetes context path '%s'", lastDirPath)
	checkErr := commons.CheckIfFolderExist(lastDirPath, true)
	if checkErr != nil {
		zap.S().Debugf("❌  Error checking existence of last Kubernetes context path '%s': %s",
			lastDirPath, checkErr.Error())
		return ""
	}

	bytes, readErr := ioutil.ReadFile(lastFilePath)
	if readErr != nil {
		zap.S().Debugf("❌  Error reading last Kubernetes context from file '%s': %s",
			lastFilePath, readErr.Error())
		return ""
	}

	return string(bytes)
}
