package list

import (
	"fmt"
	"os"
	"path/filepath"

	"bygui86/kubeconfigurator/commons"
	"bygui86/kubeconfigurator/logger"
	"bygui86/kubeconfigurator/utils"

	"github.com/urfave/cli"
)

func list(ctx *cli.Context) error {
	logger.Logger.Info("")

	logger.Logger.Debug("🐛 Get single Kubernetes configurations files path")
	singleConfigsPath := ctx.String(commons.SingleConfigsFlagName)
	logger.SugaredLogger.Debugf("🐛 Single Kubernetes configurations files path: %s", singleConfigsPath)

	checkErr := utils.CheckIfFolderExist(singleConfigsPath, false)
	if checkErr != nil {
		logger.SugaredLogger.Warn("⚠️  Single Kubernetes configurations files path not found")
		logger.SugaredLogger.Warn("ℹ️  Tip: run 'konf split' before 'konf list'")
	} else {
		logger.SugaredLogger.Infof("📚 List single Kubernetes configurations in '%s':", singleConfigsPath)
		err := filepath.Walk(singleConfigsPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				logger.SugaredLogger.Infof("\t%s", info.Name())
			}
			return nil
		})
		if err != nil {
			return cli.NewExitError(
				fmt.Sprintf("❌ Error listing single Kubernetes configurations in '%s': %s", singleConfigsPath, err.Error()),
				21)
		}
	}

	logger.Logger.Info("")
	return nil
}
