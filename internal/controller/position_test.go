package controller

import (
	"errors"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type posRepoMock struct {
}

func (p posRepoMock) Create(position *domain.Position) error {
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (p posRepoMock) GetAll() []domain.Position {
	//TODO implement me
	panic("implement me")
}

func TestPositionsController_GetPosition(t *testing.T) {

	//    expected := "dummy data"
	//    svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//        fmt.Fprintf(w, expected)
	//    }))
	//    defer svr.Close()
	//    c := NewClient(svr.URL)
	//    res, err := c.UpperCase("anything")
	//    if err != nil {
	//        t.Errorf("expected err to be nil got %v", err)
	//    }
	//    // res: expected\r\n
	//    // due to the http protocol cleanup response
	//    res = strings.TrimSpace(res)
	//    if res != expected {
	//        t.Errorf("expected res to be %s got %s", expected, res)
	//    }

	repo := posRepoMock{}

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
			expected: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			h := NewPositionsController(repo)
			svr := httptest.NewServer(http.HandlerFunc(h.GetPosition))

			defer svr.Close()

			req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", svr.URL, tt.id), http.NoBody)
			//req.SetPathValue("id", "err")
			fmt.Println(req.PathValue("id"))
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
