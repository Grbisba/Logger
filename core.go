package logger

import (
	"go.uber.org/zap/zapcore"
)

// NewCore is created a new core for logger
func NewCore() zapcore.Core {
	cfg := Configure()
	core := zapcore.NewTee(
		zapcore.NewCore(cfg.Encoder, cfg.consoleErrors, cfg.highPriority),
		zapcore.NewCore(cfg.Encoder, cfg.consoleDebugging, cfg.lowPriority),
	)
	return core
}
