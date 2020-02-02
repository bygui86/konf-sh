package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logEncodingEnvVar = "KONF_LOG_ENCODING"
	logLevelEnvVar    = "KONF_LOG_LEVEL"

	logEncodingDefault = "console"
	logLevelDefault    = "info"
)

var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	encoding := os.Getenv(logEncodingEnvVar)
	if strings.Compare(encoding, "") == 0 {
		encoding = logEncodingDefault
	}
	level := os.Getenv(logLevelEnvVar)
	if strings.Compare(level, "") == 0 {
		level = logLevelDefault
	}
	zapLevel := zapcore.InfoLevel
	err := zapLevel.Set(level)
	if err != nil {
		fmt.Printf("❌ Error initializing zap logger: %s\n", err.Error())
		os.Exit(1)
	}
	buildLogger(encoding, zapLevel)
}

func buildLogger(encoding string, level zapcore.Level) {
	Logger, _ = zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildEncoderConfig(level),
	}.Build()
	SugaredLogger = Logger.Sugar()
}

func buildEncoderConfig(level zapcore.Level) zapcore.EncoderConfig {
	if level == zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	} else {
		return zapcore.EncoderConfig{
			MessageKey: "message",
		}
	}
}
