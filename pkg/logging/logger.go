package logging

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	logger, err := buildZapLogger(
		loadConfigs(),
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("❌  Error initializing logger: %s", err.Error()))
	}
	zap.ReplaceGlobals(logger)
}

func loadConfigs() (string, zapcore.Level) {
	encoding := os.Getenv(encodingEnvVar)
	if strings.Compare(encoding, "") == 0 {
		encoding = encodingDefault
	}
	if !isEncodingAvailable(encoding) {
		fmt.Println("❌  Error initializing logger: unknown encoding")
		os.Exit(1)
	}

	level := os.Getenv(levelEnvVar)
	if strings.Compare(level, "") == 0 {
		level = levelDefault
	}
	if !isLevelAvailable(level) {
		fmt.Println("❌  Error initializing logger: unknown level")
		os.Exit(1)
	}
	zapLevel, err := getZapLevel(level)
	if err != nil {
		fmt.Println(fmt.Sprintf("❌  Error initializing logger: %s", err.Error()))
		os.Exit(1)
	}
	return encoding, zapLevel
}
