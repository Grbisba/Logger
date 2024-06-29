package logger

import "go.uber.org/zap/zapcore"

// NewCore is created a new core for logger
func NewCore() zapcore.Core {
	cfg := Configure()
	core := zapcore.NewTee(
		zapcore.NewCore(cfg.jsonEncoder, cfg.consoleErrors, cfg.highPriority),
		zapcore.NewCore(cfg.jsonEncoder, cfg.consoleDebugging, cfg.lowPriority),
	)
	return core
}
