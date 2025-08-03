package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func placeHolder(w http.ResponseWriter, r *http.Request) {

}

func SetDevicesRouter(router *mux.Router) {
	router.Handle("/", http.HandlerFunc(placeHolder)).Methods("GET")
	router.Handle("/", http.HandlerFunc(placeHolder)).Methods("DELETE")
	router.Handle("{deviceId}", http.HandlerFunc(placeHolder)).Methods("DELETE")
}
