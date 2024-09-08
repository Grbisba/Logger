package logger

import "go.uber.org/zap"

func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	logger := zap.New(NewCore()).WithOptions(opts...)
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func NewWithConfig(cfg Cfg, opts ...zap.Option) (*zap.Logger, error) {
	var fields []zap.Field

	if cfg.InstanceID != "" {
		fields = append(fields, WithInstanceID(cfg.InstanceID))
	}
	if cfg.Service != "" {
		fields = append(fields, WithService(cfg.Service))
	}
	if cfg.WithLayer != "" {
		fields = append(fields, WithLayer(cfg.WithLayer))
	}

	opts = append(opts, zap.Fields(fields...))

	logger, err := NewProduction(opts...)

	return logger, err
}
