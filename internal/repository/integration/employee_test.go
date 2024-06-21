//go:build integration
// +build integration

package integration

import (
	"context"
	"errors"
	"fmt"
	conf "github.com/dilyara4949/employees-api/internal/config"
	mongoDB "github.com/dilyara4949/employees-api/internal/database/mongo"
	"github.com/dilyara4949/employees-api/internal/database/postgres"
	"github.com/dilyara4949/employees-api/internal/domain"
	mongoemployee "github.com/dilyara4949/employees-api/internal/repository/mongo/employee"
	mongoposition "github.com/dilyara4949/employees-api/internal/repository/mongo/position"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/employee"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/position"
	"log"
	"reflect"
	"testing"
)

func InitDataEmployees(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository) {
	positions := []domain.Position{
		{
			ID:     "1",
			Name:   "name1",
			Salary: 1,
		},
		{
			ID:     "2",
			Name:   "name2",
			Salary: 2,
		},
		{
			ID:     "3",
			Name:   "name3",
			Salary: 3,
		},
	}

	employees := []domain.Employee{
		{
			ID:         "1",
			FirstName:  "firstname1",
			LastName:   "lastname1",
			PositionID: "1",
		},
		{
			ID:         "2",
			FirstName:  "firstname2",
			LastName:   "lastname2",
			PositionID: "2",
		},
		{
			ID:         "3",
			FirstName:  "firstname3",
			LastName:   "lastname3",
			PositionID: "3",
		},
	}

	for _, p := range positions {
		_, _ = posRepo.Create(context.Background(), p)
	}
	for _, e := range employees {
		_, _ = empRepo.Create(context.Background(), e)
	}
}

func DeleteDataEmployees(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository) {
	positions := []domain.Position{
		{
			ID:     "1",
			Name:   "name1",
			Salary: 1,
		},
		{
			ID:     "2",
			Name:   "name2",
			Salary: 2,
		},
		{
			ID:     "3",
			Name:   "name3",
			Salary: 3,
		},
	}

	employees := []domain.Employee{
		{
			ID:         "1",
			FirstName:  "firstname1",
			LastName:   "lastname1",
			PositionID: "1",
		},
		{
			ID:         "2",
			FirstName:  "firstname2",
			LastName:   "lastname2",
			PositionID: "2",
		},
		{
			ID:         "3",
			FirstName:  "firstname3",
			LastName:   "lastname3",
			PositionID: "3",
		},
	}

	for _, e := range employees {
		_ = empRepo.Delete(context.Background(), e.ID)
	}

	for _, p := range positions {
		_ = posRepo.Delete(context.Background(), p.ID)
	}
}

func initEmpRepo() (domain.EmployeesRepository, domain.PositionsRepository, error) {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	var employeeRepo domain.EmployeesRepository
	var positionRepo domain.PositionsRepository

	switch config.DatabaseType {
	case "postgres":
		db, err := postgres.ConnectPostgres(config.PostgresConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}

		positionRepo = position.NewPositionsRepository(db)
		employeeRepo = employee.NewEmployeesRepository(db, positionRepo)

	case "mongo":
		db, err := mongoDB.ConnectMongo(config.MongoConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}

		positionRepo = mongoposition.NewPositionsRepository(db, config.MongoConfig.Collections.Positions, config.MongoConfig.Collections.Employees)
		employeeRepo = mongoemployee.NewEmployeesRepository(db, config.MongoConfig.Collections.Employees, config.MongoConfig.Collections.Positions)
	default:
		return nil, nil, errors.New("Incorrect database given for tests")
	}
	return employeeRepo, positionRepo, nil
}

func TestEmployeeRepository_Create(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataEmployees(positionRepo, employeeRepo)
	defer DeleteDataEmployees(positionRepo, employeeRepo)

	tests := []struct {
		name        string
		employee    domain.Employee
		expectedErr bool
	}{
		{
			name: "OK",
			employee: domain.Employee{

				ID:         "4",
				FirstName:  "firstname4",
				LastName:   "lastname4",
				PositionID: "1",
			},
		},
		{
			name: "employee already exists",
			employee: domain.Employee{

				ID:         "1",
				FirstName:  "firstname1",
				LastName:   "lastname1",
				PositionID: "1",
			},
			expectedErr: true,
		},
		{
			name: "position not found",
			employee: domain.Employee{
				ID:         "5",
				FirstName:  "firstname5",
				LastName:   "lastname5",
				PositionID: "5",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emp, err := employeeRepo.Create(context.Background(), tt.employee)
			fmt.Println(err)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}

			if emp != nil && !reflect.DeepEqual(*emp, tt.employee) {
				t.Errorf("expected: %v, got: %v", tt.employee, &emp)
			}

			_ = employeeRepo.Delete(context.Background(), tt.employee.ID)
		})
	}
}

func TestEmployeeRepository_Get(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataEmployees(positionRepo, employeeRepo)
	defer DeleteDataEmployees(positionRepo, employeeRepo)

	tests := []struct {
		name        string
		employee    domain.Employee
		expectedErr bool
	}{
		{
			name: "OK",
			employee: domain.Employee{
				ID:         "1",
				FirstName:  "firstname1",
				LastName:   "lastname1",
				PositionID: "1",
			},
		},
		{
			name: "employee not found",
			employee: domain.Employee{
				ID:         "5",
				FirstName:  "firstname1",
				LastName:   "lastname1",
				PositionID: "1",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			emp, err := employeeRepo.Get(context.Background(), tt.employee.ID)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}

			if !tt.expectedErr && !reflect.DeepEqual(*emp, tt.employee) {
				t.Errorf("expected: %v, got: %v", tt.employee, *emp)
			}
		})
	}
}

func TestEmployeeRepository_Update(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataEmployees(positionRepo, employeeRepo)
	defer DeleteDataEmployees(positionRepo, employeeRepo)

	tests := []struct {
		name        string
		employee    domain.Employee
		expectedErr bool
	}{
		{
			name: "OK",
			employee: domain.Employee{
				ID:         "1",
				FirstName:  "firstname1qw",
				LastName:   "lastname1qw",
				PositionID: "1",
			},
		},
		{
			name: "employee not found",
			employee: domain.Employee{
				ID:         "5",
				FirstName:  "firstname1",
				LastName:   "lastname1",
				PositionID: "1",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := employeeRepo.Update(context.Background(), tt.employee)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestEmployeeRepository_Delete(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataEmployees(positionRepo, employeeRepo)
	defer DeleteDataEmployees(positionRepo, employeeRepo)

	tests := []struct {
		name        string
		employee    domain.Employee
		expectedErr bool
	}{
		{
			name: "OK",
			employee: domain.Employee{
				ID: "1",
			},
		},
		{
			name: "employee not found",
			employee: domain.Employee{
				ID: "5",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := employeeRepo.Delete(context.Background(), tt.employee.ID)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataEmployees(positionRepo, employeeRepo)
	defer DeleteDataEmployees(positionRepo, employeeRepo)

	tests := []struct {
		name     string
		page     int64
		pageSize int64
		expected []domain.Employee
	}{
		{
			name:     "First page, two records",
			page:     1,
			pageSize: 2,
			expected: []domain.Employee{
				{
					ID:         "1",
					FirstName:  "firstname1",
					LastName:   "lastname1",
					PositionID: "1",
				},
				{
					ID:         "2",
					FirstName:  "firstname2",
					LastName:   "lastname2",
					PositionID: "2",
				},
			},
		},
		{
			name:     "Second page, two records",
			page:     2,
			pageSize: 2,
			expected: []domain.Employee{
				{
					ID:         "3",
					FirstName:  "firstname3",
					LastName:   "lastname3",
					PositionID: "3",
				},
			},
		},
		{
			name:     "Page out of range",
			page:     3,
			pageSize: 2,
			expected: []domain.Employee{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employees, err := employeeRepo.GetAll(context.Background(), tt.page, tt.pageSize)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(employees, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, employees)
			}
		})
	}
}
