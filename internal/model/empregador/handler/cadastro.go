package handler

import (
	"encoding/json"
	"net/http"
	"sped/esocial/internal/model/empregador/services"

	"github.com/gorilla/mux"
)

func all(w http.ResponseWriter, r *http.Request) {
	e, err := services.All()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(e)
}

func byId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	e, err := services.ByID(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(e)
}

func byCnpj(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cnpj := vars["cnpj"]
	e, err := services.ByCNPJ(cnpj)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(e)
}
