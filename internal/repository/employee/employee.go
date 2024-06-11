package employee

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type PositionsRepository interface {
	Get(ctx context.Context, id string) (*domain.Position, error)
}

type employeeRepository struct {
	mu            sync.Mutex
	storage       map[string]domain.Employee
	positionsRepo PositionsRepository
}

func NewEmployeesRepository(positionsRepo PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{
		storage:       make(map[string]domain.Employee),
		positionsRepo: positionsRepo,
	}
}

func (e *employeeRepository) Create(ctx context.Context, employee *domain.Employee) error {
	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return fmt.Errorf("error to create employee: %w", err)
	}

	employee.ID = uuid.New().String()

	e.mu.Lock()
	defer e.mu.Unlock()
	e.storage[employee.ID] = *employee

	return nil
}

func (e *employeeRepository) Get(_ context.Context, id string) (*domain.Employee, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if employee, ok := e.storage[id]; ok {
		return &employee, nil
	}
	return nil, errors.New("employee not found")
}

func (e *employeeRepository) Update(ctx context.Context, employee domain.Employee) error {
	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return fmt.Errorf("error to update employee: %w", err)
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.storage[employee.ID]; !ok {
		return errors.New("employee not found")
	}

	e.storage[employee.ID] = employee
	return nil
}

func (e *employeeRepository) Delete(_ context.Context, id string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.storage[id]; !ok {
		return errors.New("employee not found")
	}

	delete(e.storage, id)
	return nil
}

func (e *employeeRepository) GetAll(_ context.Context) []domain.Employee {
	e.mu.Lock()
	defer e.mu.Unlock()

	employees := make([]domain.Employee, 0)

	for _, employee := range e.storage {
		employees = append(employees, employee)
	}
	return employees
}
