package domain

import "sync"

type Positions struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionsStorage struct {
	Mu      sync.Mutex
	Storage map[string]Positions
}

type PositionsRepository interface {
	Create(*Positions) error
	Get(id string) (*Positions, error)
	Update(Positions) error
	Delete(id string) error
	GetAll() ([]Positions, error)
}
