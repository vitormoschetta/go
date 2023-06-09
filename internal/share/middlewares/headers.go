package middlewares

import (
	"net/http"
)

var header = map[string]string{
	"Content-Type": "application/json",
}

func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range header {
			w.Header().Set(key, value)
		}
		next.ServeHTTP(w, r)
	})
}
