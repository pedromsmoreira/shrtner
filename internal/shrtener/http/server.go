package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"net/http"
	"sync"
)

type server struct {
	settings *configuration.Settings
	router   *gin.Engine
	server   *http.Server
	wg       sync.WaitGroup
}

func NewServer(settings *configuration.Settings, router *gin.Engine) *server {
	return &server{
		settings: settings,
		router:   router,
	}
}

func (s *server) Start() error {
	err := s.router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.settings.Server.Host, s.settings.Server.Port),
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}

func (s *server) Shutdown() error {
	s.wg.Wait()
	return nil
}
