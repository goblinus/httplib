package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	MetaStorager interface {
		GetVersion() string
		GetRelease() string
		GetBuilder() string
		GetBuildTime() string
	}

	LiveHandler struct {
		storage MetaStorager
	}

	ReadyHandler struct{}
)

func NewLiveHandler(storage MetaStorager) *LiveHandler {
	return &LiveHandler{storage}
}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (l LiveHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version":  l.storage.GetVersion(),
		"release":  l.storage.GetRelease(),
		"builder":  l.storage.GetBuilder(),
		"dateTime": l.storage.GetBuildTime(),
	})
}

func (r ReadyHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"state": "OK",
	})
}
