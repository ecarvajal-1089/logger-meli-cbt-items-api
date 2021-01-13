package tracing

import "context"

// TraceContextBuilder s
type TraceContextBuilder struct {
	requestID string
	fullPath  string
	flow      string
}

// NewTraceContextBuilder s
func NewTraceContextBuilder() *TraceContextBuilder {
	return new(TraceContextBuilder)
}

// WithRequestID s
func (builder *TraceContextBuilder) WithRequestID(requestID string) *TraceContextBuilder {
	builder.requestID = requestID
	return builder
}

// WithFullPath s
func (builder *TraceContextBuilder) WithFullPath(fullPath string) *TraceContextBuilder {
	builder.fullPath = fullPath
	return builder
}

// WithFlow s
func (builder *TraceContextBuilder) WithFlow(flow string) *TraceContextBuilder {
	builder.flow = flow
	return builder
}

// BuildContext s
func (builder *TraceContextBuilder) BuildContext(ctx context.Context) context.Context {
	if builder.requestID != stringEmpty {
		ctx = withRequestID(ctx, builder.requestID)
	}

	if builder.fullPath != stringEmpty {
		ctx = withFullPath(ctx, builder.fullPath)
	}

	if builder.flow != stringEmpty {
		ctx = WithFlow(ctx, builder.flow)
	}

	return ctx
}
