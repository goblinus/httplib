package healthsrv

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func WithCustomLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Errors) > 0 {
			for _, err := range ctx.Errors {
				logger.Error().Err(err).Msg("")
			}
		} else {
			logger.Info().
				Str("uri", ctx.Request.RequestURI).
				Str("method", ctx.Request.Method).
				Int("status", ctx.Writer.Status()).
				Msg("")
		}
	}
}
