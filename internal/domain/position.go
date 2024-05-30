package domain

import "sync"

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionsStorage struct {
	Mu      sync.Mutex
	Storage map[string]Position
}

type PositionsRepository interface {
	Create(*Position) error
	Get(id string) (*Position, error)
	Update(Position) error
	Delete(id string) error
	GetAll() ([]Position, error)
}
