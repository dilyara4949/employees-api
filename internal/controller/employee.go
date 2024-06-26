package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/redis"
	"github.com/google/uuid"
)

type EmployeesController struct {
	Repo  domain.EmployeesRepository
	cache redis.EmployeeCache
}

const (
	pageDefault     = 1
	pageSizeDefault = 50
)

func NewEmployeesController(repo domain.EmployeesRepository, cache redis.EmployeeCache) *EmployeesController {
	return &EmployeesController{repo, cache}
}

func (c *EmployeesController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at get employee", Status: http.StatusMethodNotAllowed})
		return
	}

	employeeID := r.PathValue("id")
	if employeeID == "" {
		errorHandler(w, r, &HTTPError{Detail: "error getting employee: id is incorrect", Status: http.StatusInternalServerError})
		return
	}

	employee, err := c.cache.Get(r.Context(), employeeID)
	if err != nil || employee == nil {
		employee, err = c.Repo.Get(r.Context(), employeeID)
		if err != nil {
			errorHandler(w, r, &HTTPError{Detail: "error getting employee", Status: http.StatusInternalServerError, Cause: err})
			return
		}

		err = c.cache.Set(r.Context(), employeeID, employee)
		if err != nil {
			log.Printf("error at caching employee: %v", err)
		}
	}

	response, err := json.Marshal(employee)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *EmployeesController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at create employee", Status: http.StatusMethodNotAllowed})
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	var employee *domain.Employee
	if err := json.Unmarshal(body, &employee); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	employee.ID = uuid.New().String()

	employee, err = c.Repo.Create(r.Context(), *employee)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error creating employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	err = c.cache.Set(r.Context(), employee.ID, employee)
	if err != nil {
		log.Printf("error at caching employee: %v", err)
	}

	response, err := json.Marshal(employee)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *EmployeesController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at delete employee", Status: http.StatusMethodNotAllowed})
		return
	}

	employeeID := r.PathValue("id")
	if employeeID == "" {
		errorHandler(w, r, &HTTPError{Detail: "error deleting employee: id is incorrect", Status: http.StatusInternalServerError})
		return
	}

	err := c.Repo.Delete(r.Context(), employeeID)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error deleting employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	err = c.cache.Delete(r.Context(), employeeID)
	if err != nil {
		log.Printf("error at deleting employee from cache: %v", err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *EmployeesController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at update employee", Status: http.StatusMethodNotAllowed})
		return
	}

	employeeID := r.PathValue("id")
	if employeeID == "" {
		errorHandler(w, r, &HTTPError{Detail: "error updating employee: id is incorrect", Status: http.StatusInternalServerError})
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	var employee domain.Employee
	if err := json.Unmarshal(body, &employee); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err})
		return
	}

	employee.ID = employeeID
	if err := c.Repo.Update(r.Context(), employee); err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error updating employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	err = c.cache.Set(r.Context(), employee.ID, &employee)
	if err != nil {
		log.Printf("error at updating employee cache: %v", err)
	}

	response, err := json.Marshal(employee)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *EmployeesController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, &HTTPError{Detail: "invalid method at get all employees", Status: http.StatusMethodNotAllowed})
		return
	}

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)

	if page <= 0 || pageSize <= 0 {
		page = pageDefault
		pageSize = pageSizeDefault
	}

	employees, err := c.Repo.GetAll(r.Context(), page, pageSize)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at getting all employees", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	response, err := json.Marshal(employees)
	if err != nil {
		errorHandler(w, r, &HTTPError{Detail: "error at marshal employees", Status: http.StatusInternalServerError, Cause: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
