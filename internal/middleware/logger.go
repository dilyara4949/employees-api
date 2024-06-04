package middleware

import (
	"log"
	"net/http"
)

func Logger() Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(CorrelationID)
			if id == nil {
				log.Println("Correlation id set incorrect")
				http.Error(w, "internal server error: Correlation id set incorrect", http.StatusInternalServerError)
			}

			log.Printf("%s %s, correlationID=%v", r.Method, r.URL, id)
			h.ServeHTTP(w, r)
		}
	}
}
