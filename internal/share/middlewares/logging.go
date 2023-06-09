package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		traceID := r.Context().Value(CorrelationIDHeader).(string)
		log.Printf("%s Método: %s | Rota: %s\n", traceID, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		timeElapsed := time.Since(timeStart)
		log.Printf("%s Tempo de execução: %s\n", traceID, timeElapsed)
	})
}
