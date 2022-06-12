package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
)

func Delete(repository data.ReadDelete) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := repository.GetById(context.Background(), id)

		if err != nil {
			c.JSON(http.StatusNotFound, NewNotFoundError(c.Request.URL.Path))
			return
		}

		err = repository.Delete(context.Background(), id)

		if err != nil {
			fmt.Print("unsucessful deletion for id " + id)
		}

		c.Status(http.StatusNoContent)
	}
}
