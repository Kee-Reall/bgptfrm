package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func clearAll(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK: DELETING ALL")); err != nil {
		fmt.Println(err.Error())
	}
}

func setTestingRouter(r *mux.Router) {

	r.HandleFunc("/all-data", clearAll).Methods("GET")
}
