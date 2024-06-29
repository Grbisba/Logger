package logger

import "go.uber.org/zap/zapcore"

func NewCore() zapcore.Core {
	cfg := Configure()
	core := zapcore.NewTee(
		zapcore.NewCore(cfg.jsonEncoder, cfg.consoleErrors, cfg.highPriority),
		zapcore.NewCore(cfg.jsonEncoder, cfg.consoleDebugging, cfg.lowPriority),
	)
	return core
}
