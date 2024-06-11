package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

func Chain(endpoint http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middle := range middlewares {
		endpoint = middle(endpoint)
	}

	return endpoint
}
