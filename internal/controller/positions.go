package controller

import (
	"encoding/json"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
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
	position, err := c.Repo.Get(r.Context(), positionID)

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

	var position domain.Position
	if err := json.Unmarshal(body, &position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	position.ID = uuid.New().String()

	if err = c.Repo.Create(r.Context(), &position); err != nil {
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
	err := c.Repo.Delete(r.Context(), positionID)

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

	var position domain.Position
	if err := json.Unmarshal(body, &position); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	position.ID = positionID
	if err := c.Repo.Update(r.Context(), position); err != nil {
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

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "page format is incorrect", Status: http.StatusBadGateway, Cause: err})
		return
	}

	pageSize, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 32)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "page size format is incorrect", Status: http.StatusBadGateway, Cause: err})
		return
	}

	positions, err := c.Repo.GetAll(r.Context(), page, pageSize)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at get all positions", Status: http.StatusInternalServerError, Cause: err})
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
