package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"net/http"
)

func Create(dns string, repository data.Create) func(c *gin.Context) {
	return func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)
		var body UrlMetadata
		err := decoder.Decode(&body)

		b, _ := json.Marshal(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestError("could not decode the request body",
				map[string]interface{}{
					"request": string(b),
					"error":   err,
				}))
			return
		}

		if body.Original == "" {
			c.JSON(http.StatusBadRequest, NewBadRequestError("[original] property should not be empty",
				map[string]interface{}{
					"request": render.JSON{Data: body},
				}))
			return
		}

		u, err := domain.NewUrl(body.Original, body.ExpirationDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		cUrl, err := repository.Create(context.Background(), u)
		// TODO: use custom error
		if err != nil {
			c.JSON(http.StatusConflict, NewConflictError(err.Error()))
			return
		}

		rBody := &UrlMetadata{
			Original:       cUrl.Original,
			Short:          fmt.Sprintf("%s/%s", dns, cUrl.Short),
			ExpirationDate: cUrl.ExpirationDate,
			DateCreated:    cUrl.DateCreated,
		}

		c.JSON(http.StatusCreated, rBody)
	}
}
