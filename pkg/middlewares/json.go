package middlewares

import (
	"net/http"
)

func AcceptJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
