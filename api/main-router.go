package api

import (
	"blog-api/api/superAdmin"
	"github.com/gorilla/mux"
	"net/http"
)

type RouterMechanic func(*mux.Router)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func InitMainRouter() *mux.Router {

	mainRouter := mux.NewRouter()
	mainRouter.NotFoundHandler = globalHttpLogger(http.HandlerFunc(notFound))

	//получаем глобальный роутер
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	//мапа механикам.
	pathPrefixes := map[string]RouterMechanic{
		"/testing": setTestingRouter,
		"/devises": SetDevicesRouter,
		"/sa":      superAdmin.SetSuperAdminRouter,
	}

	//передаём саброутеры функциям механикам.
	//Они проинициализируют обработчики, на интересующие роуты, для саброутера

	for path, mechanic := range pathPrefixes {
		mechanic(apiRouter.PathPrefix(path).Subrouter())
	}

	mainRouter.Use(getGlobalMiddlewares()...)
	return mainRouter
}
