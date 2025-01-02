package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLoggingMiddleware(excludePaths []string) gin.HandlerFunc {
	var skipPaths map[string]string = make(map[string]string)
	_ = skipPaths

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)

		var event *zerolog.Event

		status := ctx.Writer.Status()
		if len(ctx.Errors) > 0 {
			event = log.Error()
		} else if status > 499 {
			event = log.Error()
		} else if status > 399 {
			event = log.Warn()
		} else {
			event = log.Info()
		}

		event.Str("latency", stopTime.String()).
			Str("status", fmt.Sprintf("%d", status)).
			Str("path", ctx.Request.URL.Path).
			Str("method", ctx.Request.Method).
			Msg("")
	}
}
