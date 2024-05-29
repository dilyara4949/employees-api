package domain

import "sync"

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionStorage struct {
	Mu      sync.Mutex
	Storage map[string]Position
}

type PositionRepository interface {
	Create(*Position) error
	Get(id string) (*Position, error)
	Update(Position) error
	Delete(id string) error
	GetAll() ([]Position, error)
}
