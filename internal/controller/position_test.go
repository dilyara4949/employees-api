package controller

import (
	"errors"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type posRepoMock struct {
}

func (p posRepoMock) Create(position *domain.Position) error {
	if position.ID == "err" {
		return errors.New("Nope")
	}

	return nil
}

func (p posRepoMock) Get(id string) (*domain.Position, error) {
	if id == "err" {
		return nil, errors.New("Nope")
	}

	return &domain.Position{
		ID:     "id",
		Name:   "name",
		Salary: 100,
	}, nil
}

func (p posRepoMock) Update(position domain.Position) error {
	//TODO implement me
	panic("implement me")
}

func (p posRepoMock) Delete(id string) error {
	if id == "err" {
		return errors.New("test error")
	}
	return nil
}

func (p posRepoMock) GetAll() []domain.Position {
	//TODO implement me
	panic("implement me")
}

func TestPositionsController_GetPosition(t *testing.T) {
	repo := posRepoMock{}
	h := NewPositionsController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", h.GetPosition)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		id       string
		expected string
	}{
		"OK": {
			id:       "1",
			expected: "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
		},
		"err": {
			id:       "err",
			expected: "internal server error: Correlation id set incorrect\nerror getting position\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", svr.URL, tt.id), http.NoBody)
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
			if strResponse := string(response); strResponse != tt.expected {
				t.Fatalf(`expected "%s", got "%s"`, tt.expected, strResponse)
			}
		})
	}
}

func TestPositionsController_CreatePosition(t *testing.T) {
	repo := posRepoMock{}
	h := NewPositionsController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", h.CreatePosition)

	svr := httptest.NewServer(mux)
	defer svr.Close()

	tests := map[string]struct {
		body     string
		expected string
	}{
		"OK": {
			body:     "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
			expected: "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
		},
		"Empty body": {
			body:     "",
			expected: "invalid request body\n",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("POST", svr.URL, strings.NewReader(tt.body))
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

func TestPositionsController_DeletePosition(t *testing.T) {
	repo := posRepoMock{}
	h := NewPositionsController(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /{id}", h.DeletePosition)

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
			expected:     "error deleting position\n",
			expectedCode: 500,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", svr.URL, tt.id), http.NoBody)
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
