package list

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bygui86/konf-sh/pkg/commons"
	"github.com/bygui86/konf-sh/pkg/logger"
	"github.com/bygui86/konf-sh/pkg/utils"
	"github.com/urfave/cli/v2"
)

func list(ctx *cli.Context) error {
	logger.Logger.Info("")
	logger.Logger.Debug("🐛 Executing LIST command")
	logger.Logger.Debug("")

	logger.Logger.Debug("🐛 Get single Kubernetes konfigurations files path")
	singleConfigsPath := ctx.String(commons.SingleKonfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes konfigurations files path: '%s'", singleConfigsPath)

	checkErr := utils.CheckIfFolderExist(singleConfigsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warn("⚠️  Single Kubernetes konfigurations files path not found")
		logger.SugaredLogger.Warn("ℹ️  Tip: run 'konf split' before 'konf list'")
	} else {
		logger.SugaredLogger.Infof("📚 List single Kubernetes konfigurations in '%s':", singleConfigsPath)
		err := filepath.Walk(singleConfigsPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				logger.SugaredLogger.Infof("\t%s", info.Name())
			}
			return nil
		})
		if err != nil {
			return cli.Exit(
				fmt.Sprintf("❌ Error listing single Kubernetes konfigurations in '%s': %s", singleConfigsPath, err.Error()),
				21)
		}
	}

	logger.Logger.Info("")
	return nil
}
