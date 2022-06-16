package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
)

func Delete(repository data.ReadDelete) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		_, err := repository.GetById(context.Background(), id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err = encoder.Encode(w, r, NewNotFoundError(r.URL.Path)); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		err = repository.Delete(context.Background(), id)

		if err != nil {
			fmt.Print("unsucessful deletion for id " + id)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
