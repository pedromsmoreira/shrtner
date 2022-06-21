package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/pedromsmoreira/shrtener/internal/shrtner/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtner/domain"
)

func Create(dns string, repository data.Create) func(w http.ResponseWriter, r *http.Request) {
	serializer := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		var body UrlMetadata
		err := serializer.Decode(w, r, &body)
		b, _ := json.Marshal(&body)

		if err != nil {
			respond(w, r, http.StatusBadRequest, NewBadRequestError("could not decode the request body",
				map[string]interface{}{
					"request": string(b),
					"error":   err,
				}), serializer)
			return
		}

		if body.Original == "" {
			respond(w, r, http.StatusBadRequest, NewBadRequestError("[original] property should not be empty",
				map[string]interface{}{
					"request": string(b),
				}), serializer)
			return
		}

		u, err := domain.NewUrl(body.Original, body.ExpirationDate)
		if err != nil {
			respond(w, r, http.StatusInternalServerError, err.Error(), serializer)
			return
		}

		cUrl, err := repository.Create(context.Background(), u)
		if err != nil {
			logrus.WithField("error", err.Error()).Warning("conflict")
			respond(w, r, http.StatusConflict, NewConflictError(err.Error()), serializer)
			return
		}

		rBody := &UrlMetadata{
			Original:       cUrl.Original,
			Short:          fmt.Sprintf("%s/%s", dns, cUrl.Short),
			ExpirationDate: cUrl.ExpirationDate,
			DateCreated:    cUrl.DateCreated,
		}
		respond(w, r, http.StatusCreated, &rBody, serializer)
	}
}
