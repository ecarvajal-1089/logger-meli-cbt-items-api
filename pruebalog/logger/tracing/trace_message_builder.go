package tracing

import (
	"context"
	"fmt"
	"strings"
)

const (
	stringEmpty      = ""
	formatAppendText = "%s %s"
)

// TraceMessageBuilder s
type TraceMessageBuilder struct {
	message string
}

// NewTraceMessageBuilder s
func NewTraceMessageBuilder() *TraceMessageBuilder {
	return &TraceMessageBuilder{
		message: "no message found",
	}
}

// WithText s
func (builder *TraceMessageBuilder) WithText(text string) *TraceMessageBuilder {
	builder.message = text
	return builder
}

// BuildMessage s
func (builder *TraceMessageBuilder) BuildMessage(ctx context.Context) string {
	if ctx == nil {
		return builder.message
	}

	messageTrace := buildMessageTrace(ctx)
	if trace := strings.TrimSpace(messageTrace); trace != stringEmpty {
		return fmt.Sprintf(formatAppendText, trace, builder.message)
	}

	return builder.message
}

func buildMessageTrace(ctx context.Context) string {
	var messageTrace string
	if formattedResquestID, ok := formattedRequestIDFromContext(ctx); ok {
		messageTrace = formattedResquestID
	}

	if formattedFullPath, ok := formattedFullPathFromContext(ctx); ok {
		if messageTrace == stringEmpty {
			messageTrace = formattedFullPath
		} else {
			messageTrace = fmt.Sprintf(formatAppendText, messageTrace, formattedFullPath)
		}
	}

	if formattedFlow, ok := formattedFlowFromContext(ctx); ok {
		if messageTrace == stringEmpty {
			messageTrace = formattedFlow
		} else {
			messageTrace = fmt.Sprintf(formatAppendText, messageTrace, formattedFlow)
		}
	}

	return messageTrace
}
