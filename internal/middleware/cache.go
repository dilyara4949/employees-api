package middleware

import (
	"net/http"
)

func Cache() Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			h.ServeHTTP(w, r)
		}
	}
}
