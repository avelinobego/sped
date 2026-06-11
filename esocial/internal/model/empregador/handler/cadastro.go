package handler

// func all(w http.ResponseWriter, r *http.Request) {
// 	e, err := services.All()
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(e)
// }

// func byId(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	e, err := services.ByID(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(e)
// }

// func byCnpj(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	cnpj := vars["cnpj"]
// 	e, err := services.ByCNPJ(cnpj)
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(e)
// }

// func create(w http.ResponseWriter, r *http.Request) {
// 	e := &entity.Empresas{}
// 	json.NewDecoder(r.Body).Decode(e)
// 	err := services.Insert(e)
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(e)
// }
