//go:build all
// +build all

package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dilyara4949/employees-api/internal/domain"
)

type empRepoMock struct {
}

func (e empRepoMock) Create(employee *domain.Employee) error {
	if employee.ID == "err" {
		return errors.New("Nope")
	}

	return nil
}

func (e empRepoMock) Get(id string) (*domain.Employee, error) {
	if id == "err" {
		return nil, errors.New("Nope")
	}

	return &domain.Employee{
		ID:         "id",
		FirstName:  "first name",
		LastName:   "last name",
		PositionID: "position id",
	}, nil
}

func (e empRepoMock) Update(employee domain.Employee) error {
	if employee.ID == "err" {
		return errors.New("Nope")
	}

	return nil
}

func (e empRepoMock) Delete(id string) error {
	if id == "err" {
		return errors.New("Nope")
	}

	return nil
}

func (e empRepoMock) GetAll() []domain.Employee {
	return []domain.Employee{
		{
			ID:         "id",
			FirstName:  "first name",
			LastName:   "last name",
			PositionID: "position id",
		},
	}
}

func TestEmployeesController_GetEmployee(t *testing.T) {
	repo := empRepoMock{}
	h := NewEmployeesController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees/{id}", h.GetEmployee)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		id       string
		expected string
	}{
		"OK": {
			id:       "id",
			expected: "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
		},
		"err": {
			id:       "err",
			expected: "error getting employee\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), http.NoBody)
			if err != nil {
				t.Fatal(err)
			}
			hcl := http.Client{}
			resp, err := hcl.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			response, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if strResponse := string(response); strResponse != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, strResponse)
			}
		})
	}
}

func TestEmployeesController_CreateEmployee(t *testing.T) {
	repo := empRepoMock{}
	h := NewEmployeesController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees", h.CreateEmployee)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		body     string
		expected string
	}{
		"OK": {
			body:     "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
			expected: "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
		},
		"Empty body": {
			body:     "",
			expected: "invalid request body\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/employees", svr.URL), strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Error while making request: %s", err)
			}

			cl := http.Client{}
			resp, err := cl.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			response, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if res := string(response); res != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, res)
			}
		})
	}
}

func TestEmployeesController_DeleteEmployee(t *testing.T) {
	repo := empRepoMock{}
	h := NewEmployeesController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees/{id}", h.DeleteEmployee)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		id           string
		expected     string
		expectedCode int
	}{
		"OK": {
			id:           "10",
			expected:     "",
			expectedCode: 204,
		},
		"err": {
			id:           "err",
			expected:     "error deleting employee\n",
			expectedCode: 500,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			cl := http.Client{}
			resp, err := cl.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedCode {
				t.Fatalf(`expected "%d", got "%d"`, tt.expectedCode, resp.StatusCode)
			}

			response, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if res := string(response); res != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, res)
			}
		})
	}
}

func TestEmployeesController_UpdateEmployee(t *testing.T) {
	repo := empRepoMock{}
	h := NewEmployeesController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees/{id}", h.UpdateEmployee)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		id       string
		body     string
		expected string
	}{
		"OK": {
			id:       "id",
			body:     "{\"id\":\"id\",\"firstname\":\"updated first name\",\"lastname\":\"updated last name\",\"position_id\":\"position id\"}",
			expected: "{\"id\":\"id\",\"firstname\":\"updated first name\",\"lastname\":\"updated last name\",\"position_id\":\"position id\"}",
		},
		"Empty body": {
			id:       "id",
			body:     "",
			expected: "invalid request body\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("Error while making request: %s", err)
			}

			cl := http.Client{}
			resp, err := cl.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			response, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if res := string(response); res != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, res)
			}
		})
	}
}

func TestEmployeesController_GetAllEmployees(t *testing.T) {
	repo := empRepoMock{}
	h := NewEmployeesController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees", h.GetAllEmployees)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		expected string
	}{
		"OK": {
			expected: "[{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}]",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/employees", svr.URL), http.NoBody)
			if err != nil {
				t.Fatal(err)
			}

			cl := http.Client{}
			resp, err := cl.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			response, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if res := string(response); res != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, res)
			}
		})
	}
}
