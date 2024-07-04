package controller

import (
	"context"
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
	err error
}

func (e empRepoMock) Create(_ context.Context, employee domain.Employee) (*domain.Employee, error) {
	if e.err != nil {
		return nil, e.err
	}

	return &domain.Employee{
		ID:         "id",
		FirstName:  "first name",
		LastName:   "last name",
		PositionID: "position id",
	}, nil
}

func (e empRepoMock) Get(_ context.Context, id string) (*domain.Employee, error) {
	if e.err != nil {
		return nil, e.err
	}

	return &domain.Employee{
		ID:         "id",
		FirstName:  "first name",
		LastName:   "last name",
		PositionID: "position id",
	}, nil
}

func (e empRepoMock) Update(_ context.Context, employee domain.Employee) error {
	if e.err != nil {
		return e.err
	}

	return nil
}

func (e empRepoMock) Delete(_ context.Context, id string) error {
	if e.err != nil {
		return e.err
	}

	return nil
}

func (e empRepoMock) GetAll(_ context.Context, page, pageSize int64) ([]domain.Employee, error) {
	if e.err != nil {
		return nil, e.err
	}

	return []domain.Employee{
		{
			ID:         "id",
			FirstName:  "first name",
			LastName:   "last name",
			PositionID: "position id",
		},
	}, nil
}

func (e empRepoMock) GetByPosition(_ context.Context, id string) (*domain.Employee, error) {
	if e.err != nil {
		return nil, e.err
	}

	return &domain.Employee{
		ID:         "id",
		FirstName:  "first name",
		LastName:   "last name",
		PositionID: "position id",
	}, nil
}

func TestEmployeesController_GetEmployee(t *testing.T) {
	tests := map[string]struct {
		id       string
		expected string
		repo     empRepoMock
	}{
		"OK": {
			id:       "id",
			expected: "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
			repo:     empRepoMock{},
		},
		"err": {
			id:       "err",
			expected: "error getting employee\n",
			repo:     empRepoMock{err: errors.New("error")},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewEmployeesController(tt.repo)

			mux := http.NewServeMux()
			mux.HandleFunc("/employees/{id}", h.GetEmployee)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), http.NoBody)
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
	tests := map[string]struct {
		body     string
		expected string
		repo     empRepoMock
	}{
		"OK": {
			body:     "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
			expected: "{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
			repo:     empRepoMock{},
		},
		"Empty body": {
			body:     "",
			expected: "invalid request body\n",
			repo:     empRepoMock{},
		},
		"err": {
			body:     "{\"id\":\"err\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}",
			expected: "error creating employee\n",
			repo:     empRepoMock{err: errors.New("error")},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewEmployeesController(tt.repo)

			mux := http.NewServeMux()
			mux.HandleFunc("/employees", h.CreateEmployee)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/employees", svr.URL), strings.NewReader(tt.body))
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
	tests := map[string]struct {
		id           string
		expected     string
		expectedCode int
		repo         empRepoMock
	}{
		"OK": {
			id:           "10",
			expected:     "",
			expectedCode: 204,
			repo:         empRepoMock{},
		},
		"err": {
			id:           "err",
			expected:     "error deleting employee\n",
			expectedCode: 500,
			repo:         empRepoMock{err: errors.New("error")},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewEmployeesController(tt.repo)

			mux := http.NewServeMux()
			mux.HandleFunc("/employees/{id}", h.DeleteEmployee)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), http.NoBody)
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
	tests := map[string]struct {
		id       string
		body     string
		expected string
		repo     empRepoMock
	}{
		"OK": {
			id:       "id",
			body:     "{\"id\":\"id\",\"firstname\":\"updated first name\",\"lastname\":\"updated last name\",\"position_id\":\"position id\"}",
			expected: "{\"id\":\"id\",\"firstname\":\"updated first name\",\"lastname\":\"updated last name\",\"position_id\":\"position id\"}",
			repo:     empRepoMock{},
		},
		"Empty body": {
			id:       "id",
			body:     "",
			expected: "invalid request body\n",
			repo:     empRepoMock{},
		},
		"err": {
			id:       "err",
			body:     "{\"id\":\"err\",\"firstname\":\"updated first name\",\"lastname\":\"updated last name\",\"position_id\":\"position id\"}",
			expected: "error updating employee\n",
			repo:     empRepoMock{err: errors.New("error")},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewEmployeesController(tt.repo)

			mux := http.NewServeMux()
			mux.HandleFunc("/employees/{id}", h.UpdateEmployee)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/employees/%s", svr.URL, tt.id), strings.NewReader(tt.body))
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
	tests := map[string]struct {
		expected string
		repo     empRepoMock
	}{
		"OK": {
			repo:     empRepoMock{},
			expected: "[{\"id\":\"id\",\"firstname\":\"first name\",\"lastname\":\"last name\",\"position_id\":\"position id\"}]",
		},
		"err": {
			repo:     empRepoMock{err: errors.New("error")},
			expected: "error at getting all employees\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewEmployeesController(tt.repo)

			mux := http.NewServeMux()
			mux.HandleFunc("/employees", h.GetAllEmployees)

			svr := httptest.NewServer(mux)
			defer svr.Close()

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
