package controller

import (
	"encoding/json"
	"net/http"

	pos "github.com/dilyara4949/employees-api/internal/repository/position"
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

type PositionController struct {
	Repo pos.Repository
}

func NewPositionController(repo pos.Repository) *PositionController {
	return &PositionController{repo}
}

func (c *PositionController) GetPosition(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return &HTTPError{Detail: "invalid method at get position", Status: http.StatusMethodNotAllowed}
	}

	positionID := r.PathValue("id")
	position, err := c.Repo.Get(positionID)

	if err != nil {
		return &HTTPError{Detail: "error getting position", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(position)
	if err != nil {
		return &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Write(response)

	w.WriteHeader(http.StatusOK)
	return nil
}
