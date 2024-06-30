package glogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Option function
type Option func(l *zap.Logger)

func Hooks(hooks ...func(zapcore.Entry) error) Option {
	entry := zapcore.Entry{
		Level:      zap.InfoLevel,
		Time:       time.Time{},
		LoggerName: "glog",
		Message:    "",
		Caller:     zapcore.EntryCaller{},
		Stack:      "",
	}

	return
}

func Hooks(hooks ...func(zapcore.Entry) error) Option {
	return optionFunc(func(glog *) {
		= zapcore.RegisterHooks(log.core, hooks...)
	})
}

func Writer()

type hooked struct {
	Core funcs []func(Entry) error
} func (h *hooked) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	// Let the wrapped Core decide whether to log this message or not. This // also gives the downstream a chance to register itself directly with the // CheckedEntry. if downstream := h.Core.Check(ent, ce); downstream ! = nil { return downstream.AddCore(ent, h) } return ce } func (h *hooked) With(fields []Field) Core { return &hooked{ Core: h.Core.With(fields), funcs: h.funcs, } } func (h *hooked) Write(ent Entry, _ []Field) error { // Since our downstream had a chance to register itself directly with the // CheckedMessage, we don't need to call it here. var err error for i := range h.funcs { err = multierr.Append(err, h.funcs[i](ent)) } return err }