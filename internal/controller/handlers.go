package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/oi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	return r
}
