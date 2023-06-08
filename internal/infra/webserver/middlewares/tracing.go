package middlewares

import (
	"context"
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
		}
		ctx := context.WithValue(r.Context(), "X-Trace-ID", traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateTraceID() string {
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d-%d", pid, timestamp)
}
