package controller

import (
	"encoding/json"
	"github.com/dilyara4949/employees-api/internal/domain"
	"io"
	"net/http"
)

type PositionsController struct {
	Repo domain.PositionsRepository
}

func NewPositionsController(repo domain.PositionsRepository) *PositionsController {
	return &PositionsController{repo}
}

func (c *PositionsController) GetPosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at get position", Status: http.StatusMethodNotAllowed})
		return
	}

	positionID := r.PathValue("id")
	position, err := c.Repo.Get(positionID)

	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error getting position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	response, err := json.Marshal(position)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *PositionsController) CreatePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at create position", Status: http.StatusMethodNotAllowed})
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	var position domain.Positions
	if err := json.Unmarshal(body, &position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	if err = c.Repo.Create(&position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error creating position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	response, err := json.Marshal(position)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *PositionsController) DeletePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at delete position", Status: http.StatusMethodNotAllowed})
		return
	}

	positionID := r.PathValue("id")
	err := c.Repo.Delete(positionID)

	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error deleting position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *PositionsController) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at update position", Status: http.StatusMethodNotAllowed})
		return
	}

	positionID := r.PathValue("id")
	if positionID == "" {
		errorHandler(w, r, &HTTPError{Detail: "missing position ID", Status: http.StatusBadRequest})
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	var position domain.Positions
	if err := json.Unmarshal(body, &position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	position.ID = positionID
	if err := c.Repo.Update(position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error updating position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	response, err := json.Marshal(position)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal position", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *PositionsController) GetAllPositions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at get all positions", Status: http.StatusMethodNotAllowed})
		return
	}

	positions, err := c.Repo.GetAll()
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error getting positions", Status: http.StatusInternalServerError, Cause: err})
	}

	response, err := json.Marshal(positions)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal positions", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
