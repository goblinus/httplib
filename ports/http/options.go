package http

import (
	"github.com/gin-gonic/gin"
)

func WithMiddlewares(middlewares ...func(ctx *gin.Context)) gin.OptionFunc {
	return func(e *gin.Engine) {
		for _, opt := range middlewares {
			e.Use(opt)
		}
	}
}

func WithDefaultHandlers(meta MetaStorager) gin.OptionFunc {
	return func(e *gin.Engine) {
		liveHandler := NewLiveHandler(meta)
		readyHandler := NewReadyHandler()

		e.GET("/ready", readyHandler.Handle)
		e.GET("/health", liveHandler.Handle)
	}
}
