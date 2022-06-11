package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/spf13/viper"
	"net/http"
)

type ListResponse struct {
	Data []*UrlMetadata `json:"data"`
}

func List(repository data.List) func(c *gin.Context) {
	return func(c *gin.Context) {
		dbData, err := repository.List(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, &HttpError{
				Code:    "1000004",
				Message: "an error occurred in the server",
			})
		}

		urls := make([]*UrlMetadata, 0, len(dbData))
		for _, d := range dbData {
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

		c.JSON(http.StatusOK, &ListResponse{
			Data: urls,
		})
	}
}
