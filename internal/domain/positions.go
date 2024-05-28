package domain

import "sync"

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionStorage struct {
	mu      sync.Mutex
	storage map[string]Position
}

type PositionRepository interface {
	Create(Position) (Position, error)
	Get(id string) (Position, error)
	Update(Position) (Position, error)
	Delete(id string) error
	GetAll() ([]Position, error)
}
