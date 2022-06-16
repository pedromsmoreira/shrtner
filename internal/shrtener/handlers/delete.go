package handlers

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
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
			logrus.
				WithField("error", err.Error()).
				WithField("id", id).
				Info("item not deleted")
		}

		respond(w, r, http.StatusNoContent, nil, serializer)
	}
}
