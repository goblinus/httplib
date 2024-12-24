package healthsrv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	HealthController interface {
		Ping(ctx *gin.Context)
		HealthCheck(ctx *gin.Context)
	}

	HealthStorage interface {
		GetVersion() string
		GetBuilder() string
		GetBuildTime() string
	}

	controller struct {
		storage HealthStorage
	}
)

func buildHealthController(storage HealthStorage) *controller {
	return &controller{
		storage: storage,
	}
}

func buildGinEngine(controller HealthController) *gin.Engine {
	router := gin.New()
	router.Use(WithCustomLogger(&log.Logger))
	router.GET("/ping", controller.Ping)
	router.GET("/healthcheck", controller.HealthCheck)

	return router
}

func (c *controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (c *controller) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"builder":   c.storage.GetBuilder(),
		"version":   c.storage.GetVersion(),
		"buildTime": c.storage.GetBuildTime(),
	})
}
