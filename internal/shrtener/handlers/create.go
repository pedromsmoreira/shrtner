package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"github.com/spf13/viper"
	"net/http"
)

func Create(repository data.Create) func(c *gin.Context) {
	return func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)
		var body UrlMetadata
		err := decoder.Decode(&body)

		b, _ := json.Marshal(&body)

		he := validateInput(body.Original == "",
			"1000001",
			"could not decode the request body",
			map[string]interface{}{
				"request": string(b),
				"error":   err,
			})
		if he != nil {
			c.JSON(http.StatusBadRequest, he)
			return
		}

		he = validateInput(body.Original == "",
			"1000002",
			"'original' property should not be empty",
			map[string]interface{}{
				"request": render.JSON{Data: body},
			})
		if he != nil {
			c.JSON(http.StatusBadRequest, he)
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
			herr := &HttpError{
				Code:    "1000003",
				Message: err.Error(),
			}
			c.JSON(http.StatusConflict, herr)
			return
		}

		rBody := &UrlMetadata{
			Original: cUrl.Original,
			Short: fmt.Sprintf(
				"%s://%s:%d/%s",
				viper.GetString("server.protocol"),
				viper.GetString("server.host"),
				viper.GetInt("server.port"),
				cUrl.Short),
			ExpirationDate: cUrl.ExpirationDate,
			DateCreated:    cUrl.DateCreated,
		}

		c.JSON(http.StatusCreated, rBody)
	}
}
