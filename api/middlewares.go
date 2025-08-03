package api

import (
	"blog-api/utils/ink"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func getColorizedMoment(t time.Time) string {
	return ink.Cyan(t.UTC().Format("2006-01-02T15:04:05.000Z"))
}

func globalHttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(
			"%s [HTTP | request] %s %s\n",
			getColorizedMoment(time.Now()),
			ink.Magenta(r.Method),
			r.RequestURI,
		)
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
