package logger

import "go.uber.org/zap/zapcore"

// NewCore is created a new core for logger
func NewCore(config Cfg) zapcore.Core {
	productionConfig, developmentCfg := Configure(config)
	core := zapcore.NewTee(
		zapcore.NewCore(developmentCfg.consoleEncoder,
			developmentCfg.consoleErrors,
			developmentCfg.highPriority),
		zapcore.NewCore(developmentCfg.consoleEncoder,
			developmentCfg.consoleDebugging,
			developmentCfg.lowPriority),
		zapcore.NewCore(productionConfig.jsonEncoder,
			productionConfig.file,
			productionConfig.highPriority),
		zapcore.NewCore(productionConfig.jsonEncoder,
			productionConfig.file,
			productionConfig.lowPriority),
	)
	return core
}
