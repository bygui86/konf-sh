package utils

import (
	"errors"
	"os"
	"os/user"

	"bygui86/konf/logger"
)

const (
	standardOsPerm = 0755

	homeEnvVar = "HOME"
)

func CheckIfFolderExist(path string, createIfNot bool) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if createIfNot {
				return os.Mkdir(path, standardOsPerm)
			}
		}
		return err
	}
	if !info.IsDir() {
		return errors.New("specified path refers to a file, not a folder")
	}
	return nil
}

func CheckIfFileExist(filepath string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("specified path refers to a folder, not a file")
	}
	return nil
}

func GetHomeDirOrExit(methodCaller string) string {
	logger.Logger.Debug("🐛 Get HOME path")
	home, homeErr := GetHomeDir()
	if homeErr != nil {
		logger.SugaredLogger.Errorf("❌ Error creating '%s' methodCaller - Error getting HOME environment variable: '%s'", methodCaller, homeErr.Error())
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
	return GetString(homeEnvVar, userHome), nil
}

func getCurrentUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
