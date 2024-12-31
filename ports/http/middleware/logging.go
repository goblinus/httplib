package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/goblinus/httplib/ports/http/config"
)

func NewLoggingMiddleware(cfg *config.HTTPConfig) gin.HandlerFunc {
	var skipPaths map[string]string = make(map[string]string)
	_ = skipPaths

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)

		var event *zerolog.Event
		var fields map[string]string = make(map[string]string)

		fields["latency"] = stopTime.String()
		fields["status"] = ctx.Request.Response.Status
		fields["path"] = ctx.Request.URL.Path
		fields["method"] = ctx.Request.Method

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

		for name, value := range fields {
			event.Str(name, value)
		}

		event.Msg("")
	}
}
