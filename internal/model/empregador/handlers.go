package empregador

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Registrar(r *mux.Router) {
	sr := r.PathPrefix("/empresas").Subrouter()

	sr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		e, err := AllEmpresas()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
	// ---------------------------------------------------------------------------------
	sr.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		e, err := GetEmpresaByID(id)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
	// ---------------------------------------------------------------------------------
	sr.HandleFunc("/cnpj/{cnpj}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cnpj := vars["cnpj"]
		e, err := GetEmpresaByCNPJ(cnpj)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(e)
	}).Methods("GET")
}
