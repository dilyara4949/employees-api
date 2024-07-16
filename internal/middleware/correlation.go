package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

const CorrelationID = "X-Correlation-ID"

func CorrelationIDMiddleware() Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			correlationID := r.Header.Get(CorrelationID)
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			ctx := context.WithValue(r.Context(), CorrelationID, correlationID)
			r = r.WithContext(ctx)

			w.Header().Set(CorrelationID, correlationID)

			h.ServeHTTP(w, r)
		}
	}
}

func VerifyCorrelation(ctx context.Context) (any, error) {
	id := ctx.Value(CorrelationID)
	if id == nil {
		return "", errors.New("correlation id set incorrect")
	}
	return id, nil
}
