package controller

import (
	"log"
	"net/http"

	employees_api "github.com/dilyara4949/employees-api"
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

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		id := r.Context().Value(employees_api.CorrelationID)
		log.Printf("HTTP error at %v: %v, correlationID=%v", r.URL, err, id)
		if httpErr, ok := err.(*HTTPError); ok {
			http.Error(w, httpErr.Detail, httpErr.Status)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
