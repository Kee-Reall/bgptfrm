package api

import (
	"blog-api/api/superAdmin"
	"github.com/gorilla/mux"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func InitMainRouter() *mux.Router {

	mainRouter := mux.NewRouter()
	mainRouter.NotFoundHandler = globalHttpLogger(http.HandlerFunc(notFound))

	api := mainRouter.PathPrefix("/api").Subrouter()

	setTestingRouter(api.PathPrefix("/testing").Subrouter())

	SetDevicesRouter(api.PathPrefix("/devices").Subrouter())
	superAdmin.SetSuperAdminRouter(api.PathPrefix("/sa").Subrouter())

	mainRouter.Use(getGlobalMiddlewares()...)
	return mainRouter
}
