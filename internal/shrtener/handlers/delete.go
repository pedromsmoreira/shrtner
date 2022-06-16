package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/logger"
	"net/http"
)

func Delete(repository data.ReadDelete) func(w http.ResponseWriter, r *http.Request) {
	serializer := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		_, err := repository.GetById(context.Background(), id)
		if err != nil {
			respond(w, r, http.StatusNotFound, NewNotFoundError(r.URL.Path), serializer)
			return
		}

		err = repository.Delete(context.Background(), id)
		if err != nil {
			logger.Info("item not deleted", map[string]interface{}{"id": id, "error": err})
		}

		respond(w, r, http.StatusNoContent, nil, serializer)
	}
}
