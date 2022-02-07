package list

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/kubeconfig"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func list(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing LIST command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations path")
	singleKonfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKonfigsPath)

	checkErr := utils.CheckIfFolderExist(singleKonfigsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warnf("⚠️  Single Kubernetes konfigurations path not found ('%s')", singleKonfigsPath)
		logger.Logger.Warn("ℹ️  Tip: run 'konf split' before anything else")
	} else {
		validCfg := make([]string, 0)
		invalidCfg := make([]string, 0)
		err := filepath.Walk(
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
				if kubeconfig.Validate(kCfg) != nil {
					invalidCfg = append(invalidCfg, info.Name())
					return nil
				}

				validCfg = append(validCfg, info.Name())

				return nil
			},
		)
		if err != nil {
			return cli.Exit(
				fmt.Sprintf("❌  Error listing single Kubernetes konfigurations in '%s': %s", singleKonfigsPath, err.Error()),
				21)
		}

		logger.SugaredLogger.Infof("📚 Available Kubernetes konfigurations in '%s':", singleKonfigsPath)
		for _, v := range validCfg {
			logger.SugaredLogger.Infof("\t%s", v)
		}
		logger.Logger.Info("")
		logger.SugaredLogger.Infof("❓️ Invalid Kubernetes konfigurations in '%s':", singleKonfigsPath)
		for _, iv := range invalidCfg {
			logger.SugaredLogger.Infof("\t%s", iv)
		}
	}

	logger.Logger.Info("")
	return nil
}
