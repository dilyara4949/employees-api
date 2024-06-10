package employee

import (
	"context"
	"errors"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type PositionsRepository interface {
	Get(ctx context.Context, id string) (*domain.Position, error)
}

type employeeRepository struct {
	employeesStorage *EmployeesStorage
	positionsRepo    PositionsRepository
}

func NewEmployeesRepository(employeesStorage *EmployeesStorage, positionsRepo PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{employeesStorage: employeesStorage, positionsRepo: positionsRepo}
}

type EmployeesStorage struct {
	mu      sync.Mutex
	storage map[string]domain.Employee
}

func NewEmployeesStorage() *EmployeesStorage {
	return &EmployeesStorage{
		storage: make(map[string]domain.Employee),
	}
}

func (e *employeeRepository) Create(ctx context.Context, employee *domain.Employee) error {
	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return fmt.Errorf("error to create employee: %w", err)
	}

	employee.ID = uuid.New().String()
	e.employeesStorage.mu.Lock()
	defer e.employeesStorage.mu.Unlock()
	e.employeesStorage.storage[employee.ID] = *employee

	return nil
}

func (e *employeeRepository) Get(ctx context.Context, id string) (*domain.Employee, error) {
	e.employeesStorage.mu.Lock()
	defer e.employeesStorage.mu.Unlock()

	if employee, ok := e.employeesStorage.storage[id]; ok {
		return &employee, nil
	}
	return nil, errors.New("employee not found")
}

func (e *employeeRepository) Update(ctx context.Context, employee domain.Employee) error {

	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return fmt.Errorf("error to update employee: %w", err)
	}
	e.employeesStorage.mu.Lock()
	defer e.employeesStorage.mu.Unlock()

	if _, ok := e.employeesStorage.storage[employee.ID]; !ok {
		return errors.New("employee not found")
	}

	e.employeesStorage.storage[employee.ID] = employee
	return nil
}

func (e *employeeRepository) Delete(ctx context.Context, id string) error {
	e.employeesStorage.mu.Lock()
	defer e.employeesStorage.mu.Unlock()

	if _, ok := e.employeesStorage.storage[id]; !ok {
		return errors.New("employee not found")
	}

	delete(e.employeesStorage.storage, id)
	return nil
}

func (e *employeeRepository) GetAll(ctx context.Context) []domain.Employee {
	e.employeesStorage.mu.Lock()
	defer e.employeesStorage.mu.Unlock()
	employees := make([]domain.Employee, 0)
	for _, employee := range e.employeesStorage.storage {
		employees = append(employees, employee)
	}
	return employees
}
