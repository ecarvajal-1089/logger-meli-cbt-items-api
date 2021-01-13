package logger

import (
	"context"
	"proof-concept/logger-meli-cbt-items-api/pruebalog/logger/tracing"

	logtool "github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

// Info logger
func Info(ctx context.Context, message string, tags ...string) {
	message = tracing.NewTraceMessageBuilder().WithText(message).BuildMessage(ctx)
	logtool.Info(message, tags...)
}

// Infof logger
func Infof(ctx context.Context, format string, args ...interface{}) {
	format = tracing.NewTraceMessageBuilder().WithText(format).BuildMessage(ctx)
	logtool.Infof(format, args...)
}

// Warn logger
func Warn(ctx context.Context, message string, tags ...string) {
	message = tracing.NewTraceMessageBuilder().WithText(message).BuildMessage(ctx)
	logtool.Warn(message, tags...)
}

// Warnf logger
func Warnf(ctx context.Context, format string, args ...interface{}) {
	format = tracing.NewTraceMessageBuilder().WithText(format).BuildMessage(ctx)
	logtool.Warnf(format, args...)
}

// Error logger
func Error(ctx context.Context, message string, err error, tags ...string) {
	message = tracing.NewTraceMessageBuilder().WithText(message).BuildMessage(ctx)
	logtool.Error(message, err, tags...)
}

// Errorf logger
func Errorf(ctx context.Context, format string, err error, args ...interface{}) {
	format = tracing.NewTraceMessageBuilder().WithText(format).BuildMessage(ctx)
	logtool.Errorf(format, err, args...)
}
