package middleware

import (
	"log"
	"net/http"
	"time"
)

func Timer() Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(CorrelationID)
			if id == nil {
				log.Println("Correlation id set incorrect")
				http.Error(w, "internal server error: Correlation id set incorrect", http.StatusInternalServerError)
			}
			start := time.Now()

			h.ServeHTTP(w, r)

			duration := time.Since(start)
			log.Printf("Duration of request %v %v %v, correlationID=%v", duration, r.Method, r.URL, id)
		}
	}
}
