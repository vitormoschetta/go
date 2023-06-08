package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = generateTraceID()
			r.Header.Set("X-Trace-ID", traceID)
		}
		next.ServeHTTP(w, r)
	})
}

func generateTraceID() string {
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d-%d", pid, timestamp)
}
