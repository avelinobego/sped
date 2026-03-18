package shared

import "net/http"

type Handler func(w http.ResponseWriter, req *http.Request)
