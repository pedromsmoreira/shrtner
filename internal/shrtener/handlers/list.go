package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type List struct {
	Data []*UrlMetadata `json:"data"`
}

func (h *RestHandler) List(c *gin.Context) {
	data, err := h.repository.List(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &HttpError{
			Code:    "1000004",
			Message: "an error occurred in the server",
		})
	}

	urls := make([]*UrlMetadata, 0, len(data))
	for _, d := range data {
		u := &UrlMetadata{
			Original: d.Original,
			Short: fmt.Sprintf(
				"%s://%s:%d/%s",
				viper.GetString("server.protocol"),
				viper.GetString("server.host"),
				viper.GetInt("server.port"),
				d.Short),
			ExpirationDate: d.ExpirationDate,
			DateCreated:    d.DateCreated,
		}

		urls = append(urls, u)
	}

	c.JSON(http.StatusOK, &List{
		Data: urls,
	})
}
