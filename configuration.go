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
	Service    *string
	InstanceID *string
}

type ConfigFunc struct {
	highPriority     zapcore.LevelEnabler
	lowPriority      zapcore.LevelEnabler
	jsonEncoder      zapcore.Encoder
	consoleDebugging zapcore.WriteSyncer
	consoleErrors    zapcore.WriteSyncer
}

func Configure() *ConfigFunc {
	highPriority := zap.LevelEnablerFunc(highPriorityLevelEnableFunc)
	lowPriority := zap.LevelEnablerFunc(lowPriorityLevelEnableFunc)

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	cfg := &ConfigFunc{
		highPriority:     highPriority,
		lowPriority:      lowPriority,
		jsonEncoder:      jsonEncoder,
		consoleDebugging: consoleDebugging,
		consoleErrors:    consoleErrors,
	}

	return cfg
}

func highPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}
