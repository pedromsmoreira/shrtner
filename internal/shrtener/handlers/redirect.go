package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
	"time"
)

func Redirect(dns string, repository data.GetById) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		url, err := repository.GetById(context.Background(), id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err = encoder.Encode(w, r, NewNotFoundError(r.URL.Path)); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		expDate, _ := time.Parse(time.RFC3339Nano, url.ExpirationDate)
		if expDate.Before(time.Now().UTC()) {
			w.WriteHeader(http.StatusNotFound)
			if err = encoder.Encode(w, r, NewExpiredLinkError(id, url.ExpirationDate)); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		w.Header().Set("via", dns)
		http.Redirect(w, r, url.Original, http.StatusFound)
		return
	}
}
