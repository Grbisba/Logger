package glogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const (
	ServiceField    = "service"
	InstanceIDField = "instance"
	LayerField      = "layer"
)

type Config struct {
	Service    string
	InstanceID string
	Layer      string
}

type ConfigFunc struct {
	highPriority     zapcore.LevelEnabler
	lowPriority      zapcore.LevelEnabler
	consoleEncoder   zapcore.Encoder
	jsonEncoder      zapcore.Encoder
	consoleDebugging zapcore.WriteSyncer
	consoleErrors    zapcore.WriteSyncer
}

func CoreConfigure() *ConfigFunc {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",

		LevelKey:    "lvl",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.TimeEncoderOfLayout(time.RFC1123Z),

		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}

	cfg := &ConfigFunc{
		highPriority:     zap.LevelEnablerFunc(highPriorityLevelEnableFunc),
		lowPriority:      zap.LevelEnablerFunc(lowPriorityLevelEnableFunc),
		consoleEncoder:   zapcore.NewConsoleEncoder(encoderConfig),
		jsonEncoder:      zapcore.NewJSONEncoder(encoderConfig),
		consoleDebugging: zapcore.Lock(os.Stdout),
		consoleErrors:    zapcore.Lock(os.Stderr),
	}

	return cfg
}

func highPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}
