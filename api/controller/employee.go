package controller

import (
	"encoding/json"
	pos "github.com/dilyara4949/employees-api/internal/repository/employee"
	"io"
	"net/http"

	emp "github.com/dilyara4949/employees-api/internal/repository/employee"
)

type EmployeeController struct {
	Repo emp.Repository
}

func NewEmployeeController(repo emp.Repository) *EmployeeController {
	return &EmployeeController{repo}
}

func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return &HTTPError{Detail: "invalid method at get employee", Status: http.StatusMethodNotAllowed}
	}

	employeeID := r.PathValue("id")
	employee, err := c.Repo.Get(employeeID)

	if err != nil {
		return &HTTPError{Detail: "error getting employee", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(employee)
	if err != nil {
		return &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return &HTTPError{Detail: "invalid method at create employee", Status: http.StatusMethodNotAllowed}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err}
	}

	var employee pos.Employee
	if err := json.Unmarshal(body, &employee); err != nil {
		return &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err}
	}

	if err = c.Repo.Create(&employee); err != nil {
		return &HTTPError{Detail: "error creating employee", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(employee)
	if err != nil {
		return &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return nil
}

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		return &HTTPError{Detail: "invalid method at delete employee", Status: http.StatusMethodNotAllowed}
	}

	employeeID := r.PathValue("id")
	err := c.Repo.Delete(employeeID)

	if err != nil {
		return &HTTPError{Detail: "error deleting employee", Status: http.StatusInternalServerError, Cause: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return &HTTPError{Detail: "invalid method at update employee", Status: http.StatusMethodNotAllowed}
	}

	employeeID := r.PathValue("id")
	if employeeID == "" {
		return &HTTPError{Detail: "missing employee ID", Status: http.StatusBadRequest}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return &HTTPError{Detail: "error reading request body", Status: http.StatusBadRequest, Cause: err}
	}

	var employee emp.Employee
	if err := json.Unmarshal(body, &employee); err != nil {
		return &HTTPError{Detail: "invalid request body", Status: http.StatusBadRequest, Cause: err}
	}

	employee.ID = employeeID
	if err := c.Repo.Update(employee); err != nil {
		return &HTTPError{Detail: "error updating employee", Status: http.StatusInternalServerError, Cause: err}
	}

	response, err := json.Marshal(employee)
	if err != nil {
		return &HTTPError{Detail: "error at marshal employee", Status: http.StatusInternalServerError, Cause: err}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}
