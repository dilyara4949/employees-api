package middleware

import (
	"context"
	"net/http"

	employees_api "github.com/dilyara4949/employees-api"
	"github.com/google/uuid"
)

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
