package commons

import (
	"errors"
	"os"
	"os/user"

	"go.uber.org/zap"
)

func CheckIfFolderExist(path string, createIfNot bool) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if createIfNot {
				return os.Mkdir(path, 0755)
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
	zap.L().Debug("🐛 Get HOME path")
	home, homeErr := GetHomeDir()
	if homeErr != nil {
		zap.S().Errorf("❌  Error creating '%s' methodCaller - Error getting HOME environment variable: '%s'", methodCaller, homeErr.Error())
		os.Exit(3)
	}
	zap.S().Debugf("🐛 HOME path: '%s'", home)
	return home
}

func GetHomeDir() (string, error) {
	userHome, err := getCurrentUserHomeDir()
	if err != nil {
		return "", err
	}
	return GetString("HOME", userHome), nil
}

func getCurrentUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
