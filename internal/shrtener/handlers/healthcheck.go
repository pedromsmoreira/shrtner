package handlers

import (
	"fmt"
	"net/http"
)

func Ping() func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pong := map[string]interface{}{"message": "pong"}
		if err := encoder.Encode(w, r, pong); err != nil {
			fmt.Print("error encoding value... move to logger")
		}
	}
}

func Status() func(w http.ResponseWriter, r *http.Request) {
	encoder := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pong := map[string]interface{}{"status": "up"}
		if err := encoder.Encode(w, r, pong); err != nil {
			fmt.Print("error encoding value... move to logger")
		}
	}
}
