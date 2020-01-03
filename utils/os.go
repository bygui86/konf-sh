package utils

import (
	"os/user"

	"bygui86/kubeconfigurator/envvar"
)

const (
	homeEnvVar = "HOME"
)

func GetHomeDir() (string,error) {
	home, err := GetCurrentUserHomeDir()
	if err != nil {
		return "", err
	}
	return envvar.GetString(homeEnvVar, home), nil
}

func GetCurrentUserHomeDir() (string,error) {
	usr, err := user.Current()
	if err != nil{
		return "", err
	}
	return usr.HomeDir, nil
}
