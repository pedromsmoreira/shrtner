package handlers

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"net/http"
)

func (h *RestHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var body UrlMetadata
	err := decoder.Decode(&body)

	if err != nil {
		// TODO: Add custom error to have less loc
		resp := &Error{
			Code:    "1000001",
			Message: "Error decoding body",
			Details: map[string]interface{}{
				"request": r.Body,
			},
		}

		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	u, err := domain.NewUrl(body.Original, body.ExpirationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u = domain.Shorten(u)

	cUrl, err := h.repository.Create(context.Background(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rBody := &UrlMetadata{
		Original:       cUrl.Original,
		Short:          cUrl.Short,
		ExpirationDate: cUrl.ExpirationDate.String(),
		DateCreated:    cUrl.DateCreated.String(),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rBody)
}
