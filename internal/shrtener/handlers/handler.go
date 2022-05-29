package handlers

import "github.com/pedromsmoreira/shrtener/internal/shrtener/data"

type RestHandler struct {
	repository *data.CockroachDbRepository
}

func NewRestHandler(r *data.CockroachDbRepository) *RestHandler {
	return &RestHandler{
		repository: r,
	}
}
