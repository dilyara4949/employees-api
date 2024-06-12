package domain

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionsRepository interface {
	Create(*Position) error
	Get(id string) (*Position, error)
	Update(Position) error
	Delete(id string) error
	GetAll() ([]Position, error)
}
