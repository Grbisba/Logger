package logger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestReplace(t *testing.T) {
	before := zap.L()
	log, err := NewProduction()
	require.NoError(t, err)
	if assert.NotNil(t, log) {
		after := zap.L()
		assert.Equal(t, after, log)
		assert.NotEqual(t, after, before)
	}
}

func TestLogger(t *testing.T) {
	cfg := Config{
		Service:    "production",
		InstanceID: "125g51ad-r52341d",
	}
	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	logger.Info("hello world")
}
