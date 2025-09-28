package superAdmin

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetSuperAdminRouter(subRouter *mux.Router) {
	subRouter.Use(basicAuth)
	subRouter.Handle("/", http.HandlerFunc(getUsers)).Methods("GET")

}
