package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

type ErrorResponse struct {
	Errors        []string `json:"errors"`
	CorrelationID string   `json:"correlation_id"`
}

func ErrorHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID, ok := r.Context().Value(CorrelationKey).(string)
		if !ok {
			correlationID = "unknown"
		}
		defer func() {
			if r := recover(); r != nil {
				log.Print(correlationID, " ", r)
				debug.PrintStack()
				output := ErrorResponse{
					Errors:        []string{"Internal error"},
					CorrelationID: correlationID,
				}
				w.WriteHeader(http.StatusInternalServerError)
				outputJson, _ := json.Marshal(output)
				w.Write(outputJson)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
