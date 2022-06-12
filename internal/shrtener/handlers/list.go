package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
	"strconv"
)

type ListResponse struct {
	Data []*UrlMetadata `json:"data"`
	Next string
}

func List(dns string, repository data.List) func(c *gin.Context) {
	return func(c *gin.Context) {
		qPage := c.DefaultQuery("page", "0")
		qSize := c.DefaultQuery("size", "10")

		p, err := strconv.Atoi(qPage)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrorWithoutDetails(fmt.Sprintf("[page] was %v. Must be an integer.", qPage)))
			return
		}

		s, err := strconv.Atoi(qSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrorWithoutDetails(fmt.Sprintf("[size] was %v. Must be an integer.", qSize)))
			return
		}

		dbData, err := repository.List(context.Background(), p, s)
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

		response := &ListResponse{
			Data: urls,
		}

		if len(urls) > 0 {
			response.Next = fmt.Sprintf("%s/urls?page=%d&size=%d", dns, p+1, s)
		}

		c.JSON(http.StatusOK, response)
	}
}
