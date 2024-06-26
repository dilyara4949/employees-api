//go:build !integration
// +build !integration

package controller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dilyara4949/employees-api/internal/domain"
)

type posRepoMock struct {
	err error
}

type posCacheMock struct {
}

func (c posCacheMock) Set(_ context.Context, _ string, _ *domain.Position) error {
	return nil
}

func (c posCacheMock) Get(_ context.Context, _ string) (*domain.Position, error) {
	return nil, nil
}

func (c posCacheMock) Delete(_ context.Context, _ string) error {
	return nil
}

func (p posRepoMock) Create(_ context.Context, position domain.Position) (*domain.Position, error) {
	if p.err != nil {
		return nil, p.err
	}
	return &domain.Position{
		ID:     "id",
		Name:   "name",
		Salary: 100,
	}, nil
}

func (p posRepoMock) Get(_ context.Context, id string) (*domain.Position, error) {
	if p.err != nil {
		return nil, p.err
	}
	return &domain.Position{
		ID:     "id",
		Name:   "name",
		Salary: 100,
	}, nil
}

func (p posRepoMock) Update(_ context.Context, position domain.Position) error {
	if p.err != nil {
		return p.err
	}
	return nil
}

func (p posRepoMock) Delete(_ context.Context, id string) error {
	if p.err != nil {
		return p.err
	}
	return nil
}

func (p posRepoMock) GetAll(_ context.Context, page, pageSize int64) ([]domain.Position, error) {
	if p.err != nil {

		log.Println(p.err)
		return nil, p.err
	}

	return []domain.Position{
		{
			ID:     "id",
			Name:   "name",
			Salary: 100,
		},
	}, nil
}

func TestPositionsController_GetPosition(t *testing.T) {
	tests := map[string]struct {
		id       string
		expected string
		repo     posRepoMock
		cache    posCacheMock
	}{
		"OK": {
			id:       "1",
			expected: "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"err": {
			id:       "err",
			expected: "error getting position\n",
			repo:     posRepoMock{err: errors.New("error")},
			cache:    posCacheMock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewPositionsController(tt.repo, tt.cache)

			mux := http.NewServeMux()
			mux.HandleFunc("GET /{id}", h.GetPosition)

			svr := httptest.NewServer(mux)
			defer svr.Close()

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
	tests := map[string]struct {
		body     string
		expected string
		repo     posRepoMock
		cache    posCacheMock
	}{
		"OK": {
			body:     "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
			expected: "{\"id\":\"id\",\"name\":\"name\",\"salary\":100}",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"Empty body": {
			body:     "",
			expected: "invalid request body\n",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"err": {
			body:     "{\"id\":\"err\",\"name\":\"name\",\"salary\":100}",
			expected: "error creating position\n",
			repo:     posRepoMock{err: errors.New("error")},
			cache:    posCacheMock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewPositionsController(tt.repo, tt.cache)

			mux := http.NewServeMux()
			mux.HandleFunc("POST /", h.CreatePosition)

			svr := httptest.NewServer(mux)
			defer svr.Close()

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
	tests := map[string]struct {
		id           string
		expected     string
		expectedCode int
		repo         posRepoMock
		cache        posCacheMock
	}{
		"OK": {
			id:           "10",
			expected:     "",
			expectedCode: 204,
			repo:         posRepoMock{},
			cache:        posCacheMock{},
		},
		"err": {
			id:           "err",
			expected:     "error deleting position\n",
			expectedCode: 500,
			repo:         posRepoMock{err: errors.New("error")},
			cache:        posCacheMock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewPositionsController(tt.repo, tt.cache)

			mux := http.NewServeMux()
			mux.HandleFunc("DELETE /{id}", h.DeletePosition)

			svr := httptest.NewServer(mux)
			defer svr.Close()

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

func TestPositionsController_UpdatePosition(t *testing.T) {
	tests := map[string]struct {
		id       string
		body     string
		expected string
		repo     posRepoMock
		cache    posCacheMock
	}{
		"OK": {
			id:       "1",
			body:     "{\"id\":\"1\",\"name\":\"updated name\",\"salary\":200}",
			expected: "{\"id\":\"1\",\"name\":\"updated name\",\"salary\":200}",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"Empty body": {
			id:       "1",
			body:     "",
			expected: "invalid request body\n",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"err": {
			id:       "err",
			body:     "{\"err\":\"1\",\"name\":\"updated name\",\"salary\":200}",
			expected: "error updating position\n",
			repo:     posRepoMock{err: errors.New("error")},
			cache:    posCacheMock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewPositionsController(tt.repo, tt.cache)

			mux := http.NewServeMux()
			mux.HandleFunc("PUT /{id}", h.UpdatePosition)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", svr.URL, tt.id), strings.NewReader(tt.body))
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

func TestPositionsController_GetAllPositions(t *testing.T) {
	tests := map[string]struct {
		expected string
		repo     posRepoMock
		cache    posCacheMock
	}{
		"OK": {
			expected: "[{\"id\":\"id\",\"name\":\"name\",\"salary\":100}]",
			repo:     posRepoMock{},
			cache:    posCacheMock{},
		},
		"error": {
			expected: "error at getting all positions\nnull",
			repo:     posRepoMock{err: errors.New("error")},
			cache:    posCacheMock{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewPositionsController(tt.repo, tt.cache)

			mux := http.NewServeMux()
			mux.HandleFunc("GET /", h.GetAllPositions)

			svr := httptest.NewServer(mux)
			defer svr.Close()

			req, err := http.NewRequest("GET", svr.URL, http.NoBody)
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
