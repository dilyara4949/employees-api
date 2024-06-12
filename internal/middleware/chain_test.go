package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChain(t *testing.T) {
	type args struct {
		endpoint    http.HandlerFunc
		middlewares []Middleware
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "1.2.3",
			args: args{
				endpoint: func(writer http.ResponseWriter, request *http.Request) {
					writer.Write([]byte("3"))
				},
				middlewares: []Middleware{
					func(h http.Handler) http.HandlerFunc {
						return func(w http.ResponseWriter, r *http.Request) {
							w.Write([]byte("2."))
							h.ServeHTTP(w, r)
						}
					},
					func(h http.Handler) http.HandlerFunc {
						return func(w http.ResponseWriter, r *http.Request) {
							_, err := w.Write([]byte("1."))
							if err != nil {
								println(err)
							}
							h.ServeHTTP(w, r)
						}
					},
				},
			},
			expected: "1.2.3",
		},
		{
			name: "no middleware",
			args: args{
				endpoint: func(writer http.ResponseWriter, request *http.Request) {
					writer.Write([]byte("hi"))
				},
			},
			expected: "hi",
		},
		{
			name: "dont serve endpoint",
			args: args{
				endpoint: func(writer http.ResponseWriter, request *http.Request) {
					writer.Write([]byte("hi"))
				},
				middlewares: []Middleware{
					func(h http.Handler) http.HandlerFunc {
						return func(w http.ResponseWriter, r *http.Request) {
							w.Write([]byte("world"))
						}
					},
				},
			},
			expected: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()

			chain := Chain(tt.args.endpoint, tt.args.middlewares...)
			req, err := http.NewRequest("GET", "/", http.NoBody)
			if err != nil {
				t.Fatal(err)
			}
			chain.ServeHTTP(responseRecorder, req)

			defer responseRecorder.Result().Body.Close()

			resp, err := io.ReadAll(responseRecorder.Result().Body)
			if err != nil {
				t.Fatal(err)
			}

			if strResponse := string(resp); strResponse != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, strResponse)
			}
		})
	}
}
