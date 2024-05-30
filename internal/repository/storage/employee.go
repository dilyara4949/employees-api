package storage

import (
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"sync"
)

type EmployeesStorage struct {
	mu      sync.Mutex
	Storage map[string]domain.Employee
}

func (storage *EmployeesStorage) Add(employee domain.Employee) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	storage.Storage[employee.ID] = employee
}

func (storage *EmployeesStorage) Get(id string) (*domain.Employee, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if employee, ok := storage.Storage[id]; ok {
		return &employee, nil
	}
	return nil, errors.New("employee not found")
}

func (storage *EmployeesStorage) Update(employee domain.Employee) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.Storage[employee.ID]; !ok {
		return errors.New("employee not found")
	}

	storage.Storage[employee.ID] = employee
	return nil
}

func (storage *EmployeesStorage) Delete(id string) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.Storage[id]; !ok {
		return errors.New("employee not found")
	}

	delete(storage.Storage, id)
	return nil
}

func (storage *EmployeesStorage) All() []domain.Employee {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	var employees []domain.Employee
	for _, employee := range storage.Storage {
		employees = append(employees, employee)
	}
	return employees
}
