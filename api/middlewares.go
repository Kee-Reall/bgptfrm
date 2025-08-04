package api

import (
	"blog-api/utils/logger"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func globalHttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	})
}

func applicationSetter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func getGlobalMiddlewares() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{globalHttpLogger, applicationSetter}
}
