package test

import (
	myLogger "github.com/package/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

type Config struct {
	Service    *string
	InstanceID *string
}

func TestReplace(t *testing.T) {
	before := zap.L()
	log, err := myLogger.NewProduction()
	require.NoError(t, err)
	if assert.NotNil(t, log) {
		after := zap.L()
		assert.Equal(t, after, log)
		assert.NotEqual(t, after, before)
	}
}

func TestLogger(t *testing.T) {
	cfg := myLogger.Cfg{
		Service:    "",
		InstanceID: "",
		WithLayer:  "",
	}

	logger, err := myLogger.NewWithConfig(cfg)
	require.NoError(t, err)
	logger.Error("hello world", zap.String("key", "val"))
}
