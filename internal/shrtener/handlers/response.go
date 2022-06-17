package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type UrlMetadata struct {
	Original       string `json:"original"`
	Short          string `json:"short,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}, encoder serializer) {
	w.Header().Set("Content-Type", encoder.ContentType(w, r))
	w.WriteHeader(status)
	if data == nil {
		return
	}

	if err := encoder.Encode(w, r, data); err != nil {
		logrus.WithField("error", err.Error()).Warning("error encoding value")
	}
}
