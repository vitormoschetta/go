package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	CorrelationIDHeader = "X-Correlation-ID"
)

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(CorrelationIDHeader)
		if traceID == "" {
			traceID = generateTraceID()
			log.Println("New trace ID generated:", traceID)
		}
		ctx := context.WithValue(r.Context(), CorrelationIDHeader, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateTraceID() string {
	pid := os.Getpid()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d-%d", pid, timestamp)
}
