package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
)

func NewRouter(dns string, repository data.Repository) *gin.Engine {
	router := gin.Default()

	// TODO: add http metrics middleware - prometheus
	// TODO: handle http error cleanly
	// https://golangexample.com/gin-error-handling-middleware-is-a-middleware-for-the-popular-gin-framework

	router.GET("/urls", handlers.List(dns, repository))
	router.POST("/urls", handlers.Create(dns, repository))

	return router
}
