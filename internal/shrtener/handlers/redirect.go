package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
	"time"
)

func Redirect(dns string, repository data.GetById) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		url, err := repository.GetById(context.Background(), id)

		if err != nil {
			c.JSON(http.StatusNotFound, NewNotFoundError(c.Request.URL.Path))
			return
		}

		expDate, _ := time.Parse(time.RFC3339Nano, url.ExpirationDate)
		if expDate.Before(time.Now().UTC()) {
			c.JSON(http.StatusNotFound, NewExpiredLinkError(id, url.ExpirationDate))
			return
		}

		c.Header("Via", dns)
		c.Redirect(http.StatusFound, url.Original)
		return
	}
}
