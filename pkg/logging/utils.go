package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func isEncodingAvailable(encoding string) bool {
	for _, enc := range availableEncodings {
		if encoding == enc {
			return true
		}
	}
	return false
}

func isLevelAvailable(level string) bool {
	for _, lvl := range availableLevels {
		if level == lvl {
			return true
		}
	}
	return false
}

func getZapLevel(level string) (zapcore.Level, error) {
	zapLevel := zapcore.InfoLevel
	err := zapLevel.Set(level)
	if err != nil {
		return zapcore.InfoLevel, err
	}
	return zapLevel, nil
}

func buildZapLogger(encoding string, level zapcore.Level) (*zap.Logger, error) {
	return zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildZapEncoderConfig(level),
	}.Build()
}

func buildZapEncoderConfig(level zapcore.Level) zapcore.EncoderConfig {
	if level != zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			MessageKey: "message",
		}
	} else {
		return zapcore.EncoderConfig{
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			MessageKey:   "message",
		}
	}
}
