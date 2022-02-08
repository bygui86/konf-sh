package commons

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

func Check(key string) (string, bool) {
	zap.S().Debugf("🐛 Check '%s' environment variable", key)
	return os.LookupEnv(key)
}

func GetString(key, defaultValue string) string {
	zap.S().Debugf("🐛 Get '%s' environment variable as string, default '%s'", key, defaultValue)
	value, exist := Check(key)
	if !exist {
		return defaultValue
	}
	return value
}

func GetBool(key string, defaultValue bool) bool {
	zap.S().Debugf("🐛 Get '%s' environment variable as bool, default '%t'", key, defaultValue)
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func GetInt(key string, defaultValue int) int {
	zap.S().Debugf("🐛 Get '%s' environment variable as int, default '%d'", key, defaultValue)
	return int(GetInt64(key, int64(defaultValue)))
}

func GetInt64(key string, defaultValue int64) int64 {
	zap.S().Debugf("🐛 Get '%s' environment variable as int64, default '%d'", key, defaultValue)
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
