package employee

import "sync"

type Storage struct {
	mu      sync.Mutex
	storage map[string]Employee
}

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	PositionID string `json:"position_id"`
}

type Repository interface {
	Create(Employee) (*Employee, error)
	Get(id string) (*Employee, error)
	Update(Employee) error
	Delete(id string) error
	GetAll() ([]Employee, error)
}
