package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"github.com/spf13/viper"
	"net/http"
)

func (h *RestHandler) Create(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var body UrlMetadata
	err := decoder.Decode(&body)

	he := validateInput(body.Original == "",
		"1000001",
		"could not decode the request body",
		map[string]interface{}{
			"request": json.Marshal(body),
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
			"request": json.Marshal(body),
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

	cUrl, err := h.repository.Create(context.Background(), u)
	if err != nil {
		herr := &HttpError{
			Code:    "1000003",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, herr)
		return
	}

	rBody := &UrlMetadata{
		Original: cUrl.Original,
		Short: fmt.Sprintf(
			"http://%s:%d/%s",
			viper.GetString("server.host"),
			viper.GetInt("server.port"),
			cUrl.Short),
		ExpirationDate: cUrl.ExpirationDate.String(),
		DateCreated:    cUrl.DateCreated.String(),
	}

	c.JSON(http.StatusCreated, rBody)
}
