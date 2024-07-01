package middleware

import (
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strings"
	"time"
)

func Cache(cache *redis.Client, ttl time.Duration) Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				h.ServeHTTP(w, r)
				return
			}

			id := r.PathValue("id")
			if id == "" {
				h.ServeHTTP(w, r)
				return
			}

			res, err := cache.Get(r.Context(), id).Result()
			if err == nil {
				log.Println("Cache hit for key:", id)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(res))
				return
			}

			log.Println("Cache miss for key:", id)

			rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
			h.ServeHTTP(rec, r)

			if rec.statusCode == http.StatusOK {
				cache.Set(r.Context(), id, rec.body.String(), ttl)
			}
		}
	}
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       strings.Builder
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b)
	return rec.ResponseWriter.Write(b)
}
