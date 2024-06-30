package glogger

import (
	"go.uber.org/zap"
	"os"
)

func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	core := NewCore().jsonCore
	glog := zap.New(core, opts...)
	zap.ReplaceGlobals(glog)
	return glog, nil
}

func NewWithConfig(cfg Config, opts ...zap.Option) (*zap.Logger, error) {
	_, err := os.Create("infra/glogger.log")
	if err != nil {
		panic("failed to create temporary file")
	}

	var fields []zap.Field
	var options []zap.Option

	if cfg.InstanceID != "" {
		fields = append(fields, WithInstanceID(cfg.InstanceID))
	}
	if cfg.Service != "" {
		fields = append(fields, WithService(cfg.Service))
	}
	if cfg.Layer != "" {
		fields = append(fields, WithLayer(cfg.Layer))
	}

	opts = append(opts, zap.Fields(fields...))

	//options = append(options, WithCaller())

	glog, err := NewProduction(options...)
	if err != nil {
		return nil, err
	}

	zap.L().With()
	return glog, nil
}
