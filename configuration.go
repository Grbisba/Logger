package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	ServiceField    = "service"
	InstanceIDField = "instance"
	LayerField      = "layer"
)

type Config struct {
	Service    string
	InstanceID string
}

type ConfigFunc struct {
	highPriority     zapcore.LevelEnabler
	lowPriority      zapcore.LevelEnabler
	Encoder          zapcore.Encoder
	consoleDebugging zapcore.WriteSyncer
	consoleErrors    zapcore.WriteSyncer
}

func Configure() ConfigFunc {
	highPriority := zap.LevelEnablerFunc(highPriorityLevelEnableFunc)
	lowPriority := zap.LevelEnablerFunc(lowPriorityLevelEnableFunc)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	cfg := ConfigFunc{
		highPriority:     highPriority,
		lowPriority:      lowPriority,
		Encoder:          consoleEncoder,
		consoleDebugging: consoleDebugging,
		consoleErrors:    consoleErrors,
	}

	return cfg
}

func zapConfig() zap.Config {
	zapCfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	return zapCfg
}
func highPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}
