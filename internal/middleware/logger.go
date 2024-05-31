package middleware

import (
	"log"
	"net/http"

	employees_api "github.com/dilyara4949/employees-api"
)

func Logger() Adapter {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(employees_api.CorrelationID)
			log.Printf("%s %s, correlationID=%v", r.Method, r.URL, id)
			h.ServeHTTP(w, r)
		}
	}
}
