package employee

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/storage"

	"github.com/google/uuid"
)

type PositionsRepository interface {
	Get(id string) (*domain.Position, error)
}

type employeeRepository struct {
	employeesStorage *storage.EmployeesStorage
	positionsRepo    PositionsRepository
}

func NewEmployeesRepository(employeesStorage *storage.EmployeesStorage, positionsRepo PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{employeesStorage: employeesStorage, positionsRepo: positionsRepo}
}

func (e *employeeRepository) Create(employee *domain.Employee) error {
	if _, err := e.positionsRepo.Get(employee.PositionID); err != nil {
		return fmt.Errorf("error to create employee: %w", err)
	}

	employee.ID = uuid.New().String()
	e.employeesStorage.Add(*employee)

	return nil
}

func (e *employeeRepository) Get(id string) (*domain.Employee, error) {
	return e.employeesStorage.Get(id)
}

func (e *employeeRepository) Update(employee domain.Employee) error {

	if _, err := e.positionsRepo.Get(employee.PositionID); err != nil {
		return fmt.Errorf("error to update employee: %w", err)
	}
	return e.employeesStorage.Update(employee)
}

func (e *employeeRepository) Delete(id string) error {
	return e.employeesStorage.Delete(id)
}

func (e *employeeRepository) GetAll() ([]domain.Employee, error) {
	return e.employeesStorage.All()
}
