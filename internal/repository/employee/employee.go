package employee

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type PositionsRepository interface {
	Get(id string) (*domain.Positions, error)
}

type employeeRepository struct {
	db *domain.EmployeesStorage
	p  PositionsRepository
}

func NewEmployeesRepository(db *domain.EmployeesStorage, p PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{db: db, p: p}
}

func (e *employeeRepository) Create(employee *domain.Employees) error {
	e.db.Mu.Lock()
	defer e.db.Mu.Unlock()

	if _, err := e.p.Get(employee.PositionID); err != nil {
		return fmt.Errorf("error to create employee: %w", err)
	}

	employee.ID = uuid.New().String()
	e.db.Storage[employee.ID] = *employee

	return nil
}

func (e *employeeRepository) Get(id string) (*domain.Employees, error) {
	e.db.Mu.Lock()
	defer e.db.Mu.Unlock()

	if _, ok := e.db.Storage[id]; !ok {
		return nil, fmt.Errorf("employee with id %s does not exists", id)
	}
	employee := e.db.Storage[id]
	return &employee, nil
}

func (e *employeeRepository) Update(employee domain.Employees) error {
	e.db.Mu.Lock()
	defer e.db.Mu.Unlock()

	if _, ok := e.db.Storage[employee.ID]; !ok {
		return fmt.Errorf("employee does not exist")
	}

	if _, err := e.p.Get(employee.PositionID); err != nil {
		return fmt.Errorf("error to update employee: %w", err)
	}

	e.db.Storage[employee.ID] = employee
	return nil
}

func (e *employeeRepository) Delete(id string) error {
	e.db.Mu.Lock()
	defer e.db.Mu.Unlock()

	if _, ok := e.db.Storage[id]; !ok {
		return fmt.Errorf("employee does not exist")
	}

	delete(e.db.Storage, id)
	return nil
}

func (e *employeeRepository) GetAll() ([]domain.Employees, error) {
	employees := make([]domain.Employees, 0)

	e.db.Mu.Lock()
	defer e.db.Mu.Unlock()
	for _, employee := range e.db.Storage {
		employees = append(employees, employee)
	}
	return employees, nil
}