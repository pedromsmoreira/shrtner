package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
)

func NewRouter(handlers *handlers.RestHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/urls", handlers.List)

	return router
}
