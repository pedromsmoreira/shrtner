package http

import (
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
)

func NewRouter(dns string, repository data.Repository) *mux.Router {
	router := mux.NewRouter()

	// TODO(feat): add http metrics middleware - prometheus
	// TODO(improvement): handle http error cleanly
	// https://golangexample.com/gin-error-handling-middleware-is-a-middleware-for-the-popular-gin-framework

	router.HandleFunc("/urls", handlers.List(dns, repository)).Methods("GET")
	router.HandleFunc("/urls", handlers.Create(dns, repository)).Methods("POST")
	router.HandleFunc("/urls/{id}", handlers.Delete(repository)).Methods("DELETE")
	router.HandleFunc("/{id}", handlers.Redirect(dns, repository)).Methods("GET")

	return router
}
