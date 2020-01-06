package utils

import (
	"os"
	"os/user"

	"bygui86/kubeconfigurator/config/envvar"
	"bygui86/kubeconfigurator/logger"
)

const (
	standardOsPerm = 0755

	homeEnvVar = "HOME"
)

func CheckIfFolderExist(path string, createIfNot bool) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if createIfNot {
				return os.Mkdir(path, standardOsPerm)
			}
		}
		return err
	}
	return nil
}

func GetHomeDirOrExit(command string) string {
	logger.Logger.Debug("🐛 Get HOME path")
	home, homeErr := GetHomeDir()
	if homeErr != nil {
		logger.SugaredLogger.Errorf("❌ Error creating '%s' command - Error getting HOME environment variable: %s", command, homeErr.Error())
		os.Exit(3)
	}
	logger.SugaredLogger.Debugf("🐛 HOME path: '%s'", home)
	return home
}

func GetHomeDir() (string, error) {
	userHome, err := getCurrentUserHomeDir()
	if err != nil {
		return "", err
	}
	return envvar.GetString(homeEnvVar, userHome), nil
}

func getCurrentUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
