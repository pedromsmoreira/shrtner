package handlers

import (
	"net/http"
)

func Ping() func(w http.ResponseWriter, r *http.Request) {
	serializer := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		pong := map[string]interface{}{"message": "pong"}
		respond(w, r, http.StatusOK, pong, serializer)
	}
}

func Status() func(w http.ResponseWriter, r *http.Request) {
	serializer := JSON
	return func(w http.ResponseWriter, r *http.Request) {
		pong := map[string]interface{}{"status": "up"}
		respond(w, r, http.StatusOK, pong, serializer)
	}
}
