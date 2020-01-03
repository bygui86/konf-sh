package envvar

import (
	"os"
	"strconv"
)

func Check(key string) bool {
	return os.Getenv(key) != ""
}

func Set(key string, value string) error {
	return os.Setenv(key, value)
}

func Unset(key string) error {
	return os.Unsetenv(key)
}

func GetString(key, defaultValue string) string {
	if !Check(key) {
		return defaultValue
	}
	return os.Getenv(key)
}

func GetBool(key string, defaultValue bool) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func GetInt(key string, defaultValue int) int {
	return int(GetInt64(key, int64(defaultValue)))
}

func GetInt64(key string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
