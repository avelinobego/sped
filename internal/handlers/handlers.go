package handlers

import (
	"sped/esocial/internal/model/empregador"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	empregador.Registrar(r)
	return r
}
