package handlers

import (
	"encoding/json"
	"net/http"
	"sped/esocial/internal/model/empregador"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()
}

func GetRouter() *mux.Router {
	allEmpresas()
	getEmpresaByID()
	getEmpresasByCNPJ()
	return r
}

func allEmpresas() {
	r.HandleFunc("/empresas", func(w http.ResponseWriter, r *http.Request) {
		e, err := empregador.AllEmpresas()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
}

func getEmpresaByID() {
	r.HandleFunc("/empresa/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		e, err := empregador.GetEmpresaByID(id)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
}

func getEmpresasByCNPJ() {
	r.HandleFunc("/empresa/cnpj/{cnpj}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cnpj := vars["cnpj"]
		e, err := empregador.GetEmpresaByCNPJ(cnpj)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
}
