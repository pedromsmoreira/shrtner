package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Handler) Status(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Status")
}
