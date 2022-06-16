package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
)

func Redirect(dns string, repository data.GetById) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		url, err := repository.GetById(context.Background(), id)

		if err != nil {
			respond(w, r, http.StatusNotFound, NewNotFoundError(r.URL.Path), encoder)
			return
		}

		expDate, _ := time.Parse(time.RFC3339Nano, url.ExpirationDate)
		if expDate.Before(time.Now().UTC()) {
			respond(w, r, http.StatusNotFound, NewExpiredLinkError(id, url.ExpirationDate), encoder)
			return
		}

		w.Header().Set("via", dns)
		http.Redirect(w, r, url.Original, http.StatusFound)
		return
	}
}
