package logger

import "go.uber.org/zap"

func NewProduction(config Cfg, opts ...zap.Option) (*zap.Logger, error) {
	logger := zap.New(NewCore(config)).WithOptions(opts...)
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func NewWithConfig(config Cfg, opts ...zap.Option) (*zap.Logger, error) {
	var fields []zap.Field

	if config.InstanceID != "" {
		fields = append(fields, WithInstanceID(config.InstanceID))
	}
	if config.Service != "" {
		fields = append(fields, WithService(config.Service))
	}
	if config.WithLayer != "" {
		fields = append(fields, WithLayer(config.WithLayer))
	}

	opts = append(opts, zap.Fields(fields...))

	logger, err := NewProduction(config, opts...)

	return logger, err
}
