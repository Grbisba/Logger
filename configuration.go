package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	ServiceField    = "service"
	InstanceIDField = "instance"
	LayerField      = "layer"
)

type Cfg struct {
	Service              string
	InstanceID           string
	WithLayer            string
	FilePathWithFileName string
	MaxFileAge           int
	MaxFileBackups       int
	MaxFileSize          int
}

type FileConfig struct {
	highPriority zapcore.LevelEnabler
	lowPriority  zapcore.LevelEnabler
	jsonEncoder  zapcore.Encoder
	file         zapcore.WriteSyncer
}

type ConsoleConfig struct {
	highPriority     zapcore.LevelEnabler
	lowPriority      zapcore.LevelEnabler
	consoleEncoder   zapcore.Encoder
	consoleDebugging zapcore.WriteSyncer
	consoleErrors    zapcore.WriteSyncer
}

func Configure(config Cfg) (*FileConfig, *ConsoleConfig) {
	highPriority := zap.LevelEnablerFunc(highPriorityLevelEnableFunc)
	lowPriority := zap.LevelEnablerFunc(lowPriorityLevelEnableFunc)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	logFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.FilePathWithFileName,
		MaxSize:    config.MaxFileSize,
		MaxBackups: config.MaxFileBackups,
		MaxAge:     config.MaxFileAge,
	})

	prodCfg := zap.NewProductionEncoderConfig()
	prodCfg.TimeKey = "timestamp"
	prodCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoder := zapcore.NewJSONEncoder(prodCfg)

	devCfg := zap.NewDevelopmentEncoderConfig()
	devCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(devCfg)

	pCfg := &FileConfig{
		highPriority: highPriority,
		lowPriority:  lowPriority,
		jsonEncoder:  jsonEncoder,
		file:         logFile,
	}
	cCfg := &ConsoleConfig{
		highPriority:     highPriority,
		lowPriority:      lowPriority,
		consoleEncoder:   consoleEncoder,
		consoleDebugging: consoleDebugging,
		consoleErrors:    consoleErrors,
	}
	return pCfg, cCfg
}

func highPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnableFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}
