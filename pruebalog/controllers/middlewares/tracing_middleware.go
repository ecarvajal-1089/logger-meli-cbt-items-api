package middlewares

import (
	"fmt"
	trace "proof-concept/logger-meli-cbt-items-api/pruebalog/logger/tracing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/tracing"
)

// Tracing s
func Tracing(ctxGin *gin.Context) {
	ctx := ctxGin.Request.Context()
	requestID := tracing.RequestID(ctx)
	fullPath := fmt.Sprintf("%s %s", ctxGin.Request.Method, ctxGin.FullPath())
	//flow := flowNaming(ctxGin.Request.Method, ctxGin.FullPath())

	ctx = trace.NewTraceContextBuilder().
		WithRequestID(requestID).
		WithFullPath(fullPath).
		//WithFlow(flow).
		BuildContext(ctx)

	ctxGin.Request = ctxGin.Request.WithContext(ctx)
}
