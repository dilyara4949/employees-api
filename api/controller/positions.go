package controller

import (
	"encoding/json"
	"io"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}

func (c *PositionController) CreatePosition(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return &HTTPError{Detail: "invalid method at create position", Status: http.StatusMethodNotAllowed}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err}
	}

	var position pos.Position
	if err := json.Unmarshal(body, &position); err != nil {
		return &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err}
	}

	if position, err = c.Repo.Create(position); err != nil {
		return &HTTPError{Detail: "error creating position", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(position)
	if err != nil {
		return &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}

func (c *PositionController) DeletePosition(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		return &HTTPError{Detail: "invalid method at delete position", Status: http.StatusMethodNotAllowed}
	}

	positionID := r.PathValue("id")
	err := c.Repo.Delete(positionID)

	if err != nil {
		return &HTTPError{Detail: "error deleting position", Status: http.StatusInternalServerError, Cause: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *PositionController) UpdatePosition(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return &HTTPError{Detail: "invalid method at update position", Status: http.StatusMethodNotAllowed}
	}

	positionID := r.PathValue("id")
	if positionID == "" {
		return &HTTPError{Detail: "missing position ID", Status: http.StatusBadRequest}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err}
	}

	var position pos.Position
	if err := json.Unmarshal(body, &position); err != nil {
		return &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err}
	}

	position.ID = positionID
	if err := c.Repo.Update(position); err != nil {
		return &HTTPError{Detail: "error updating position", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(position)
	if err != nil {
		return &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}
