package logger

import (
	"go.uber.org/zap/zapcore"
)

// NewCore is created a new core for logger
func NewCore() zapcore.Core {
	cfgr := Configure()
	core := zapcore.NewTee(
		zapcore.NewCore(cfgr.jsonEncoder, cfgr.consoleErrors, cfgr.highPriority),
		zapcore.NewCore(cfgr.jsonEncoder, cfgr.consoleDebugging, cfgr.lowPriority),
	)
	return core
}
