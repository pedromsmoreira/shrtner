package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Redirect")
}
