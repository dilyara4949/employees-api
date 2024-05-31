package middleware

import (
	"context"
	"log"
	"net/http"

	employees_api "github.com/dilyara4949/employees-api"
	"github.com/google/uuid"
)

//const CorrelationIDHeader = "X-Correlation-ID"

func Logger() Adapter {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(employees_api.CorrelationID)
			log.Printf("%s %s, correlationID=%v", r.Method, r.URL, id)
			h.ServeHTTP(w, r)
		}
	}
}

func СorrelationIDMiddleware() Adapter {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			correlationID := r.Header.Get(employees_api.CorrelationID)
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			ctx := context.WithValue(r.Context(), employees_api.CorrelationID, correlationID)
			r = r.WithContext(ctx)

			w.Header().Set(employees_api.CorrelationID, correlationID)

			h.ServeHTTP(w, r)
		}
	}
}
