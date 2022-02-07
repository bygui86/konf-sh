package utils

import (
	"os"
	"strconv"

	"github.com/bygui86/konf-sh/pkg/logger"
)

func Check(key string) (string, bool) {
	logger.SugaredLogger.Debugf("🐛 Check '%s' environment variable", key)
	return os.LookupEnv(key)
}

func Set(key string, value string) error {
	logger.SugaredLogger.Debugf("🐛 Set '%s' environment variable: '%s'", key, value)
	return os.Setenv(key, value)
}

func Unset(key string) error {
	logger.SugaredLogger.Debugf("🐛 Unset '%s' environment variable", key)
	return os.Unsetenv(key)
}

func GetString(key, defaultValue string) string {
	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable as string, default '%s'", key, defaultValue)
	value, exist := Check(key)
	if !exist {
		return defaultValue
	}
	return value
}

func GetBool(key string, defaultValue bool) bool {
	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable as bool, default '%s'", key, defaultValue)
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func GetInt(key string, defaultValue int) int {
	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable as int, default '%d'", key, defaultValue)
	return int(GetInt64(key, int64(defaultValue)))
}

func GetInt64(key string, defaultValue int64) int64 {
	logger.SugaredLogger.Debugf("🐛 Get '%s' environment variable as int64, default '%d'", key, defaultValue)
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
