package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtner/configuration"
)

type server struct {
	settings *configuration.Settings
	router   *mux.Router
	server   *http.Server
	wg       sync.WaitGroup
}

func NewServer(settings *configuration.Settings, router *mux.Router) *server {
	return &server{
		settings: settings,
		router:   router,
	}
}

func (s *server) Start() error {
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.settings.Server.Host, s.settings.Server.Port),
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
