package handlers

import (
	"encoding/json"
	"net/http"
)

type encoder interface {
	Encode(w http.ResponseWriter, r *http.Request, v interface{}) error
	ContentType(w http.ResponseWriter, r *http.Request) string
}

type jsonEncoder struct{}

var _ encoder = (*jsonEncoder)(nil)

// JSON is an Encoder for JSON.
var JSON encoder = (*jsonEncoder)(nil)

func (j *jsonEncoder) Encode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (j *jsonEncoder) ContentType(w http.ResponseWriter, r *http.Request) string {
	return "application/json; charset=utf-8"
}
