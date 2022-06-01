package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"net/http"
)

func (h *RestHandler) Create(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var body UrlMetadata
	err := decoder.Decode(&body)

	if err != nil {
		// TODO: Add custom error to have less loc
		resp := &Error{
			Code:    "1000001",
			Message: "could not decode the request body",
			Details: map[string]interface{}{
				"request": c.Request.Body,
				"error":   err,
			},
		}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	u, err := domain.NewUrl(body.Original, body.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	cUrl, err := h.repository.Create(context.Background(), u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rBody := &UrlMetadata{
		Original:       cUrl.Original,
		Short:          cUrl.Short,
		ExpirationDate: cUrl.ExpirationDate.String(),
		DateCreated:    cUrl.DateCreated.String(),
	}

	c.JSON(http.StatusCreated, rBody)
}
