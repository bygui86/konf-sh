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
	singleKfgsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations path: '%s'", singleKfgsPath)

	checkErr := utils.CheckIfFolderExist(singleKfgsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warnf("🚨 Single Kubernetes konfigurations path not found ('%s')", singleKfgsPath)
		logger.Logger.Warn("💬 Tip: run 'konf split' before anything else")
	} else {
		validCfgs := make([]string, 0)
		invalidCfgs := make([]string, 0)
		err := filepath.Walk(
			singleKfgsPath,
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
					invalidCfgs = append(invalidCfgs, info.Name())
					return nil
				}

				validCfgs = append(validCfgs, info.Name())

				return nil
			},
		)
		if err != nil {
			return cli.Exit(
				fmt.Sprintf("❌  Error listing single Kubernetes konfigurations in '%s': %s", singleKfgsPath, err.Error()),
				21)
		}

		if len(validCfgs) > 0 {
			logger.SugaredLogger.Infof("📚 Available Kubernetes konfigurations in '%s':", singleKfgsPath)
			for _, v := range validCfgs {
				logger.SugaredLogger.Infof("\t%s", v)
			}
		} else {
			logger.SugaredLogger.Warnf("🚨 No available Kubernetes konfigurations in '%s'", singleKfgsPath)
			logger.Logger.Warn("💬 Tip: run 'konf split' before anything else")
		}

		if len(invalidCfgs) > 0 {
			logger.Logger.Info("")
			logger.SugaredLogger.Infof("❓ Invalid Kubernetes konfigurations in '%s':", singleKfgsPath)
			for _, iv := range invalidCfgs {
				logger.SugaredLogger.Infof("\t%s", iv)
			}
		}
	}

	logger.Logger.Info("")
	return nil
}
