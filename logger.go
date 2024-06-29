package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	logger := zap.New(NewCore(), opts...)
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func NewWithConfig(cfg ConfigFunc, opts ...zap.Option) (*zap.Logger, error) {
	//var fields []zap.Field

	zcfg := zap.Config{
		Encoding:         "json",
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
	logger, err := zcfg.Build()
	if err != nil {
		return nil, err
	}

	//if cfg.InstanceID != "" {
	//	fields = append(fields, WithInstanceID(cfg.InstanceID))
	//}
	//if cfg.Service != "" {
	//	fields = append(fields, WithService(cfg.Service))
	//}
	//opts = append(opts, zap.Fields(fields...))

	zap.L().With()
	return logger, err
}
