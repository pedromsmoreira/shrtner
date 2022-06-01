package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
)

func NewRouter(handlers *handlers.RestHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/urls", handlers.List)
	router.POST("/urls", handlers.Create)

	return router
}
