package healthsrv

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

type (
	HealthServer struct {
		httpServer *http.Server
	}
)

func BuildHealthServer(addr string, storage HealthStorage) *HealthServer {
	controller := buildHealthController(storage)
	router := buildGinEngine(controller)

	return &HealthServer{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (hs *HealthServer) Run(ctx context.Context) error {
	if err := hs.httpServer.ListenAndServe(); err == nil || !errors.Is(err, http.ErrServerClosed) {
		log.Info().Str(
			"address",
			hs.httpServer.Addr,
		).Msg("health server is running")
	}

	<-ctx.Done()

	return nil
}

func (hs *HealthServer) Shutdown(ctx context.Context) error {
	if err := hs.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
