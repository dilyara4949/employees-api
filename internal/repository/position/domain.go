package position

import "sync"

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type Storage struct {
	mu      sync.Mutex
	storage map[string]Position
}

type Repository interface {
	Create(Position) (*Position, error)
	Get(id string) (*Position, error)
	Update(Position) error
	Delete(id string) error
	GetAll() ([]*Position, error)
}
