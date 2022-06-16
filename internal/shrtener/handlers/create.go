package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"net/http"
)

func Create(dns string, repository data.Create) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body UrlMetadata
		err := decoder.Decode(&body)

		b, _ := json.Marshal(&body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err = encoder.Encode(w, r,
				NewBadRequestError("could not decode the request body",
					map[string]interface{}{
						"request": string(b),
						"error":   err,
					})); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		if body.Original == "" {
			w.WriteHeader(http.StatusBadRequest)
			if err = encoder.Encode(w, r,
				NewBadRequestError("[original] property should not be empty",
					map[string]interface{}{
						"request": string(b),
					})); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		u, err := domain.NewUrl(body.Original, body.ExpirationDate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if err = encoder.Encode(w, r, err.Error()); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		cUrl, err := repository.Create(context.Background(), u)
		// TODO: use custom error
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			if err = encoder.Encode(w, r, NewConflictError(err.Error())); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		rBody := &UrlMetadata{
			Original:       cUrl.Original,
			Short:          fmt.Sprintf("%s/%s", dns, cUrl.Short),
			ExpirationDate: cUrl.ExpirationDate,
			DateCreated:    cUrl.DateCreated,
		}

		w.WriteHeader(http.StatusCreated)
		if err = encoder.Encode(w, r, rBody); err != nil {
			fmt.Print("error encoding value... move to logger")
		}
	}
}
