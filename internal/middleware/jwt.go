package middleware

import (
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTAuth struct {
	secret string
}

func NewJWTAuth(secret string) *JWTAuth {
	return &JWTAuth{secret}
}

func (j *JWTAuth) Auth() Middleware {
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)

				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Bearer token required", http.StatusUnauthorized)

				return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
				}

				return []byte(j.secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			h.ServeHTTP(w, r)
		}
	}
}
