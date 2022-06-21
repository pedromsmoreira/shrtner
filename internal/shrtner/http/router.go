package http

import (
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtner/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtner/handlers"
	"github.com/pedromsmoreira/shrtener/internal/shrtner/handlers/middleware"
)

func NewRouter(dns string, repository data.ReadWriteRepository) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.LoggedHandler)
	// TODO(feat): add http metrics middleware - prometheus
	router.HandleFunc("/urls", handlers.List(dns, repository)).Methods("GET")
	router.HandleFunc("/urls", handlers.Create(dns, repository)).Methods("POST")
	router.HandleFunc("/urls/{id}", handlers.Delete(repository)).Methods("DELETE")
	router.HandleFunc("/{id}", handlers.Redirect(dns, repository)).Methods("GET")

	return router
}
