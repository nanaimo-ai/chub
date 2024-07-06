package main

import (
	"os"

	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger != nil {
		return logger
	}
	logger = NewLogger()
	return logger
}

func NewLogger() *zap.Logger {
	var encoderCfg zapcore.EncoderConfig

	if os.Getenv("ENV") == "local" {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		return prettyconsole.NewLogger(zap.DebugLevel)
	}

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.MessageKey = "message"
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.StacktraceKey = "stacktrace"
	encoderCfg.LineEnding = zapcore.DefaultLineEnding
	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid":     os.Getpid(),
			"env":     os.Getenv("CHUB_ENV"),
			"version": os.Getenv("CHUB_VERSION"),
			"project": os.Getenv("CHUB_ID"),
		},
	}

	return zap.Must(config.Build())
}
