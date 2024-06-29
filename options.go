package logger

import (
	"go.uber.org/zap"
)

// Options function
type Options func(l *zap.Logger)

func WithLayer(layer string) zap.Field {
	return zap.String(LayerField, layer)
}

func WithInstanceID(id string) zap.Field {
	return zap.String(InstanceIDField, id)
}

func WithService(service string) zap.Field {
	return zap.String(ServiceField, service)
}
