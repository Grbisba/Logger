package glogger

import (
	"go.uber.org/zap/zapcore"
)

type Core struct {
	jsonCore    zapcore.Core
	consoleCore zapcore.Core
}

// NewCore is created a new core for glogger
func NewCore() *Core {
	coreCfg := CoreConfigure()
	consoleCore := zapcore.NewTee(
		zapcore.NewCore(coreCfg.consoleEncoder, coreCfg.consoleErrors, coreCfg.highPriority),
		zapcore.NewCore(coreCfg.consoleEncoder, coreCfg.consoleDebugging, coreCfg.lowPriority),
	)
	jsonCore := zapcore.NewTee(
		zapcore.NewCore(coreCfg.jsonEncoder, coreCfg.consoleErrors, coreCfg.highPriority),
		zapcore.NewCore(coreCfg.jsonEncoder, coreCfg.consoleDebugging, coreCfg.lowPriority),
	)

	zapcore.RegisterHooks(jsonCore)

	core := &Core{
		jsonCore:    jsonCore,
		consoleCore: consoleCore,
	}
	return core
}
