package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
)

type ListResponse struct {
	Data []*UrlMetadata `json:"data"`
}

func List(dns string, repository data.List) func(c *gin.Context) {
	return func(c *gin.Context) {
		dbData, err := repository.List(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerError("an error occurred in the server"))
		}

		urls := make([]*UrlMetadata, 0, len(dbData))
		for _, d := range dbData {
			u := &UrlMetadata{
				Original:       d.Original,
				Short:          fmt.Sprintf("%s/%s", dns, d.Short),
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
