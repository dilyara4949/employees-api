package middleware

import (
	"log"
	"net/http"
)

const CorrelationIDHeader = "X-Correlation-ID"

func Logger() Adapter {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL)
			h.ServeHTTP(w, r)
		}
	}
}

//func СorrelationIDMiddleware(next http.Handler) Adapter {
//	return func(w httpttp.ResponseWriter, r *http.Request) http.HandlerFunc {
//		correlationID := r.Header.Get(CorrelationIDHeader)
//		if correlationID == "" {
//			correlationID = uuid.New().String()
//		}
//
//		ctx := context.WithValue(r.Context(), CorrelationIDHeader, correlationID)
//		r = r.WithContext(ctx)
//
//		w.Header().Set(CorrelationIDHeader, correlationID)
//
//		next.ServeHTTP(w, r)
//	}
//}
