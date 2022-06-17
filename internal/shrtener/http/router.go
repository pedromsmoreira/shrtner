package http

import (
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers/middleware"
)

func NewRouter(dns string, repository data.Repository) *mux.Router {
	router := mux.NewRouter()

	// TODO(feat): add http metrics middleware - prometheus
	router.HandleFunc("/urls", middleware.LoggedHandler(handlers.List(dns, repository))).Methods("GET")
	router.HandleFunc("/urls", middleware.LoggedHandler(handlers.Create(dns, repository))).Methods("POST")
	router.HandleFunc("/urls/{id}", middleware.LoggedHandler(handlers.Delete(repository))).Methods("DELETE")
	router.HandleFunc("/{id}", middleware.LoggedHandler(handlers.Redirect(dns, repository))).Methods("GET")

	return router
}
