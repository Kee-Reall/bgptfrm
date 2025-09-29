package api

import (
	"blog-api/src/api/superAdmin"
	"github.com/gorilla/mux"
	"net/http"
)

type RouterMechanic func(*mux.Router)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func InitMainRouter() *mux.Router {

	mainRouter := mux.NewRouter()
	mainRouter.StrictSlash(true)

	mainRouter.NotFoundHandler = globalHttpLogger(http.HandlerFunc(notFound))

	//получаем глобальный роутер
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	//мап механикам.
	pathPrefixes := map[string]RouterMechanic{
		"/testing": setTestingRouter,
		"/devises": SetDevicesRouter,
		"/sa":      superAdmin.SetSuperAdminRouter,
	}

	//Передаём саброутеры функциям механикам.
	//Они проинициализируют обработчики, для саброутера

	for path, mechanic := range pathPrefixes {
		mechanic(apiRouter.PathPrefix(path).Subrouter())
	}

	mainRouter.Use(getGlobalMiddlewares()...)
	return mainRouter
}
