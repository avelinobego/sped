package handlers

import (
	empregador "sped/esocial/internal/model/empregador/handler"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	empregador.Register(r)
	return r
}
