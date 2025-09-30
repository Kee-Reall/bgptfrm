package superAdmin

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetSuperAdminRouter(subRouter *mux.Router) {
	subRouter.Use(basicAuth)
	subRouter.Methods(http.MethodGet).PathPrefix("/users").HandlerFunc(getUsers)
	subRouter.Methods(http.MethodPost).PathPrefix("/users").HandlerFunc(createUser)
}
