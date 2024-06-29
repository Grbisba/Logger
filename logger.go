package logger

import "go.uber.org/zap"

func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	logger := zap.New(NewCore(), opts...)
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func NewWithConfig(cfg Config, opts ...zap.Option) (*zap.Logger, error) {
	var fields []zap.Field
	if cfg.InstanceID != nil {
		fields = append(fields, WithInstanceID(*cfg.InstanceID))
	}
	if cfg.Service != nil {
		fields = append(fields, WithService(*cfg.Service))
	}
	opts = append(opts, zap.Fields(fields...))

	logger, err := NewProduction(opts...)

	zap.L().With()
	return logger, err
}
