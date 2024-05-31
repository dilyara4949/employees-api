package middleware

import (
	"net/http"
)

type Adapter func(http.Handler) http.HandlerFunc

func Adapt(endpoint http.HandlerFunc, middlewares ...Adapter) http.HandlerFunc {
	for _, adapter := range middlewares {
		endpoint = adapter(endpoint)
	}
	return endpoint
}
