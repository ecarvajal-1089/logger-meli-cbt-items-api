package tracing

import (
	"context"
	"fmt"
)

type traceKey string

var (
	traceKeyRequestID = traceKey("x-request-id")
	traceKeyFullPath  = traceKey("full-path")
	traceKeyFlow      = traceKey("flow")
)

func (key traceKey) format(value string) string {
	return fmt.Sprintf("[%s:%s]", string(key), value)
}

// withRequestID s
func withRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, traceKeyRequestID, requestID)
}

// withFullPath s
func withFullPath(ctx context.Context, fullPath string) context.Context {
	return context.WithValue(ctx, traceKeyFullPath, fullPath)
}

// WithFlow s
func WithFlow(ctx context.Context, flow string) context.Context {
	return context.WithValue(ctx, traceKeyFlow, flow)
}

// formattedRequestIDFromContext gets the x-request-id from the context formatted for trace
func formattedRequestIDFromContext(ctx context.Context) (string, bool) {
	return formatTrace(ctx, traceKeyRequestID)
}

// formattedFullPathFromContext gets the full-path (endpoint) from the context formatted for trace
func formattedFullPathFromContext(ctx context.Context) (string, bool) {
	return formatTrace(ctx, traceKeyFullPath)
}

// formattedFlowFromContext gets the flow from the context formatted for trace
func formattedFlowFromContext(ctx context.Context) (string, bool) {
	return formatTrace(ctx, traceKeyFlow)
}

func formatTrace(ctx context.Context, key traceKey) (string, bool) {
	if value, ok := ctx.Value(key).(string); ok {
		return key.format(value), ok
	}

	return stringEmpty, false
}
