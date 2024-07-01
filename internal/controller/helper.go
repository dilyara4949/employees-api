package controller

import (
	"github.com/dilyara4949/employees-api/internal/middleware"
	"log"
	"net/http"
)

type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		correlationId := r.Context().Value(middleware.CorrelationID)
		if correlationId == nil {
			log.Println("Correlation id set incorrect")
			http.Error(w, "internal server error: Correlation id set incorrect", http.StatusInternalServerError)
		}

		log.Printf("HTTP error at %v: %v, correlationID=%v", r.URL, err, correlationId)

		if httpErr, ok := err.(*HTTPError); ok {
			http.Error(w, httpErr.Detail, httpErr.Status)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
