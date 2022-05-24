package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/rest"
)

func NewRouter(endpoints *rest.Handler) *httprouter.Router {
	router := httprouter.New()
	router.GET("/urls", endpoints.List)
	return router
}
