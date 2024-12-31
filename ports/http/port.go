// Пакет базовой функциональности порта для http-сервиса.
// Port: структура с базовой функциональностью для поднятия сервиса
// PortOption: тип для определения опций дополнительной настройки http-сервера

package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/goblinus/httplib/v2/ports/http/config"
)

type (
	PortOption func() *Port

	Port struct {
		server *http.Server
	}
)

const shutdownTimeout = 5 * time.Second

func NewPort(c *config.HTTPConfig, router Router) *Port {
	return &Port{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", c.Addr, c.Port),
			Handler: router.Routes(),
		}}
}

func (p Port) Start() <-chan error {
	result := make(chan error)
	go func() {
		defer close(result)
		log.Info().
			Str("addr", p.server.Addr).
			Msg("starting http server")
		err := p.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			result <- err
			return
		}
	}()

	return result
}

func (p Port) Stop(ctx context.Context) <-chan error {
	result := make(chan error)
	go func() {
		defer close(result)
		ctx, cancelFn := context.WithTimeout(ctx, shutdownTimeout)
		defer cancelFn()

		err := p.server.Shutdown(ctx)
		if err != nil {
			result <- err
		}

		log.Info().Msg("http server stopped")
	}()

	return result
}
