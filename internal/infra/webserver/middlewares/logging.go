package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		traceID := r.Header.Get("X-Trace-ID")
		log.Printf("Método: %s | Rota: %s | TraceID: %s\n", r.Method, r.URL.Path, traceID)
		next.ServeHTTP(w, r)
		timeElapsed := time.Since(timeStart)
		log.Printf("Tempo de execução: %s | TraceID: %s\n", timeElapsed, traceID)
	})
}
