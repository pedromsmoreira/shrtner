package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *RestHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "List")
}
