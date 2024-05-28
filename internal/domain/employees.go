package domain

import "sync"

type EmployeeStorage struct {
	mu      sync.Mutex
	storage map[string]Employee
}

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	PositionID string `json:"position_id"`
}

type EmployeeRepository interface {
	Create(Employee) (Employee, error)
	Get(id string) (Employee, error)
	Update(Employee) (Employee, error)
	Delete(id string) error
	GetAll() ([]Employee, error)
}
