package superAdmin

import (
	"blog-api/src/utils/logger"
	"encoding/base64"
	"net/http"
	"strings"
)

func basicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		header = strings.TrimSpace(header)
		authStrParts := strings.Split(header, " ")
		if len(authStrParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if authStrParts[0] != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authData, err := base64.StdEncoding.DecodeString(authStrParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		decodedParts := strings.Split(string(authData), ":")
		if len(decodedParts) != 2 || decodedParts[0] != "admin" || decodedParts[1] != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		logger.Info("ADMIN OPERATION ON " + r.RequestURI)
		h.ServeHTTP(w, r)
	})
}
