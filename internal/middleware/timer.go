package middleware

import (
	"log"
	"net/http"
	"time"

	employees_api "github.com/dilyara4949/employees-api"
)

func Timer() Adapter {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(employees_api.CorrelationID)
			start := time.Now()

			h.ServeHTTP(w, r)

			duration := time.Since(start)
			log.Printf("Duration of request %v %v %v, correlationID=%v", duration, r.Method, r.URL, id)
		}
	}
}
