package employee

import (
	"fmt"

	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/google/uuid"
)

type PositionRepository interface {
	Get(id string) (position.Position, error)
}

type employeeRepository struct {
	db *Storage
	p  PositionRepository
}

func NewEmployeeRepository(db *Storage, p PositionRepository) Repository {
	return &employeeRepository{db: db, p: p}
}

func (e *employeeRepository) Create(employee Employee) (*Employee, error) {
	e.db.mu.Lock()
	defer e.db.mu.Unlock()

	if _, err := e.p.Get(employee.PositionID); err != nil {
		return nil, fmt.Errorf("error to create employee: %w", err)
	}

	employee.ID = uuid.New().String()
	e.db.storage[employee.ID] = employee

	return &employee, nil
}

func (e *employeeRepository) Get(id string) (*Employee, error) {
	e.db.mu.Lock()
	defer e.db.mu.Unlock()

	if _, ok := e.db.storage[id]; !ok {
		return nil, fmt.Errorf("employee with id %s does not exists", id)
	}
	employee := e.db.storage[id]
	return &employee, nil
}

func (e *employeeRepository) Update(employee Employee) error {
	e.db.mu.Lock()
	defer e.db.mu.Unlock()

	if _, ok := e.db.storage[employee.ID]; !ok {
		return fmt.Errorf("employee does not exist")
	}

	e.db.storage[employee.ID] = employee
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	e.db.mu.Lock()
	defer e.db.mu.Unlock()

	if _, ok := e.db.storage[id]; !ok {
		return fmt.Errorf("employee does not exist")
	}

	delete(e.db.storage, id)
	return nil
}

func (e *employeeRepository) GetAll() ([]Employee, error) {
	return nil, nil
}
