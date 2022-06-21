package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
)

func Redirect(dns string, repository data.ReadKeyRepository) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		redirect, err := repository.GetRedirect(context.Background(), id)

		if err != nil {
			respond(w, r, http.StatusNotFound, NewNotFoundError(r.URL.Path), encoder)
			return
		}

		expDate, _ := time.Parse(time.RFC3339Nano, redirect.ExpirationDate)
		if expDate.Before(time.Now().UTC()) {
			respond(w, r, http.StatusNotFound, NewExpiredLinkError(id, redirect.ExpirationDate), encoder)
			return
		}

		w.Header().Set("via", dns)
		http.Redirect(w, r, redirect.Original, http.StatusFound)
		return
	}
}
