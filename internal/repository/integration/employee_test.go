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

func initData(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository) ([]*domain.Position, []*domain.Employee, error) {
	positions := []*domain.Position{
		{
			Name:   "name1",
			Salary: 1,
		},
		{
			Name:   "name2",
			Salary: 2,
		},
		{
			Name:   "name3",
			Salary: 3,
		},
	}

	employees := []*domain.Employee{
		{
			FirstName: "firstname1",
			LastName:  "lastname1",
		},
		{
			FirstName: "firstname2",
			LastName:  "lastname2",
		},
		{
			FirstName: "firstname3",
			LastName:  "lastname3",
		},
	}

	var errs, err error

	for i := 0; i < len(positions); i++ {
		positions[i], err = posRepo.Create(context.Background(), *positions[i])
		if err != nil {
			errs = errors.Join(errs, err)
		}

		employees[i].PositionID = positions[i].ID
		employees[i], err = empRepo.Create(context.Background(), *employees[i])
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return positions, employees, errs
}

func DeleteData(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository, employees []*domain.Employee, positions []*domain.Position) error {
	var errs error
	for _, e := range employees {
		err := empRepo.Delete(context.Background(), e.ID)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	for _, p := range positions {
		err := posRepo.Delete(context.Background(), p.ID)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return errs
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

	poss, emps, err := initData(positionRepo, employeeRepo)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := DeleteData(positionRepo, employeeRepo, emps, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

	tests := []struct {
		name        string
		employee    domain.Employee
		expectedErr bool
	}{
		{
			name: "OK",
			employee: domain.Employee{
				FirstName:  "firstname4",
				LastName:   "lastname4",
				PositionID: poss[0].ID,
			},
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
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
			if emp != nil {
				tt.employee.ID = emp.ID
			}
			if emp != nil && !reflect.DeepEqual(*emp, tt.employee) {
				t.Errorf("expected: %v, got: %v", tt.employee, *emp)
			}
			err = employeeRepo.Delete(context.Background(), tt.employee.ID)
			if err != nil {
				fmt.Println("error at deleting created employee: v", err, tt.employee, emp)
			}
		})
	}
}

func TestEmployeeRepository_Get(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, emps, err := initData(positionRepo, employeeRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := DeleteData(positionRepo, employeeRepo, emps, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

	tests := []struct {
		name        string
		employees   []*domain.Employee
		expectedErr bool
	}{
		{
			name:      "OK",
			employees: emps,
		},
		{
			name: "employee not found",
			employees: []*domain.Employee{
				{
					ID:         "5",
					FirstName:  "firstname1",
					LastName:   "lastname1",
					PositionID: "1",
				},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, employee := range tt.employees {
				emp, err := employeeRepo.Get(context.Background(), employee.ID)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}

				if !tt.expectedErr && !reflect.DeepEqual(emp, employee) {
					t.Errorf("expected: %v, got: %v", employee, *emp)
				}
			}
		})
	}
}

func TestEmployeeRepository_Update(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, emps, err := initData(positionRepo, employeeRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := DeleteData(positionRepo, employeeRepo, emps, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

	tests := []struct {
		name        string
		employees   []*domain.Employee
		expectedErr bool
	}{
		{
			name:      "OK",
			employees: emps,
		},
		{
			name: "employee not found",
			employees: []*domain.Employee{
				{
					ID:         "5",
					FirstName:  "firstname1",
					LastName:   "lastname1",
					PositionID: "1",
				},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, emp := range tt.employees {
				err := employeeRepo.Update(context.Background(), *emp)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}
			}
		})
	}
}

func TestEmployeeRepository_Delete(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}
	_, emps, err := initData(positionRepo, employeeRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}

	tests := []struct {
		name        string
		employees   []*domain.Employee
		expectedErr bool
	}{
		{
			name:      "OK",
			employees: emps,
		},
		{
			name: "employee not found",
			employees: []*domain.Employee{
				{
					ID: "5",
				},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, emp := range tt.employees {
				err := employeeRepo.Delete(context.Background(), emp.ID)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}
			}
		})
	}
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	employeeRepo, positionRepo, err := initEmpRepo()
	if err != nil {
		t.Fatal(err)
	}
	poss, emps, err := initData(positionRepo, employeeRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := DeleteData(positionRepo, employeeRepo, emps, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

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
				*emps[0],
				*emps[1],
			},
		},
		{
			name:     "Second page, two records",
			page:     2,
			pageSize: 2,
			expected: []domain.Employee{
				*emps[2],
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
