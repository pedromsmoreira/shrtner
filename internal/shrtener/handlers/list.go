package handlers

import (
	"context"
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"net/http"
	"strconv"
)

const (
	defaultPageNumber = "0"
	defaultPageSize   = "10"
)

type ListResponse struct {
	Data []*UrlMetadata `json:"data"`
	Next string
}

func List(dns string, repository data.List) func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		qPage := defaultPageNumber
		qSize := defaultPageSize
		if r.URL.Query().Has("page") {
			qPage = r.URL.Query().Get("page")
		}
		if r.URL.Query().Has("size") {
			qSize = r.URL.Query().Get("size")
		}

		p, err := strconv.Atoi(qPage)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err = encoder.Encode(w, r, NewBadRequestErrorWithoutDetails(fmt.Sprintf("[page] was %v. Must be an integer.", qPage))); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		s, err := strconv.Atoi(qSize)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err = encoder.Encode(w, r, NewBadRequestErrorWithoutDetails(fmt.Sprintf("[size] was %v. Must be an integer.", qSize))); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		dbData, err := repository.List(context.Background(), p, s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err = encoder.Encode(w, r, NewInternalServerError("an error occurred in the server")); err != nil {
				fmt.Print("error encoding value... move to logger")
			}
			return
		}

		urls := make([]*UrlMetadata, 0, len(dbData))
		for _, d := range dbData {
			u := &UrlMetadata{
				Original:       d.Original,
				Short:          fmt.Sprintf("%s/%s", dns, d.Short),
				ExpirationDate: d.ExpirationDate,
				DateCreated:    d.DateCreated,
			}

			urls = append(urls, u)
		}

		response := &ListResponse{
			Data: urls,
		}

		if len(urls) > 0 {
			response.Next = fmt.Sprintf("%s/urls?page=%d&size=%d", dns, p+1, s)
		}

		w.WriteHeader(http.StatusOK)
		if err = encoder.Encode(w, r, response); err != nil {
			fmt.Print("error encoding value... move to logger")
		}
	}
}
