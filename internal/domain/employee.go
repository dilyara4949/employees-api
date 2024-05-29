package domain

import "sync"

type EmployeeStorage struct {
	Mu      sync.Mutex
	Storage map[string]Employee
}

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	PositionID string `json:"position_id"`
}

type EmployeeRepository interface {
	Create(*Employee) error
	Get(id string) (*Employee, error)
	Update(Employee) error
	Delete(id string) error
	GetAll() ([]Employee, error)
}