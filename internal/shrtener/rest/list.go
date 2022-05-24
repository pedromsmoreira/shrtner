package rest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "List")
}
