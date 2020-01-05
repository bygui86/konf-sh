package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"bygui86/kubeconfigurator/envvar"
)

const (
	logEncodingEnvVar = "LOG_ENCODING"
	logLevelEnvVar = "LOG_LEVEL"

	logEncodingDefault = "console"
	logLevelDefault = "info"
)

var Logger *zap.Logger

func init() {
	encoding := envvar.GetString(logEncodingEnvVar, logEncodingDefault)
	levelString := envvar.GetString(logLevelEnvVar, logLevelDefault)
	level := zapcore.InfoLevel
	err := level.Set(levelString)
	if err != nil {
		panic(err)
	}
	buildLogger(encoding, level)
}

func buildLogger(encoding string, level zapcore.Level) {
	Logger, _ = zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildEncoderConfig(level),
	}.Build()
}

func buildEncoderConfig(level zapcore.Level) zapcore.EncoderConfig{
	if level == zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			//EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeTime: zapcore.RFC3339TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	} else {
		return zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			//EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeTime: zapcore.RFC3339TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		}
	}
}
