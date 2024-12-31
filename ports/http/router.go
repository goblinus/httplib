package http

import "github.com/gin-gonic/gin"

type (
	Router interface {
		Routes() *gin.Engine
		Init(opt ...gin.OptionFunc)
	}

	HTTPRouter struct {
		router *gin.Engine
	}
)

func NewHTTPRouter() *HTTPRouter {
	router := gin.New()
	return &HTTPRouter{
		router: router,
	}
}

func (r HTTPRouter) Init(opts ...gin.OptionFunc) {
	for _, opt := range opts {
		opt(r.Routes())
	}
}

func (r HTTPRouter) Routes() *gin.Engine {
	return r.router
}
