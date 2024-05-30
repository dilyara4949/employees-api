package domain

import "sync"

type EmployeesStorage struct {
	Mu      sync.Mutex
	Storage map[string]Employees
}

type Employees struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	PositionID string `json:"position_id"`
}

type EmployeesRepository interface {
	Create(*Employees) error
	Get(id string) (*Employees, error)
	Update(Employees) error
	Delete(id string) error
	GetAll() ([]Employees, error)
}
