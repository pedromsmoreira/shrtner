package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"log"
	"net/http"
	"sync"
)

type server struct {
	settings *configuration.Settings
	router   *httprouter.Router
	wg       sync.WaitGroup
}

func NewServer(settings *configuration.Settings, router *httprouter.Router) *server {
	return &server{
		settings: settings,
		router:   router,
	}
}

func (s *server) Start() error {
	s.wg.Add(1)
	go func(settings *configuration.Settings, r *httprouter.Router) {
		defer s.wg.Done()
		err := http.ListenAndServe(fmt.Sprintf(":%d", settings.Server.Port), r)
		if err != nil {
			log.Fatal(err)
		}
	}(s.settings, s.router)

	return nil
}

func (s *server) Shutdown() error {
	s.wg.Wait()
	return nil
}
