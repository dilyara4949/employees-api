package employee

import (
	"context"
	"fmt"
	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/database"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"log"
	"reflect"
	"strconv"
	"testing"
)

func SetEnv(t *testing.T) {
	cfg := map[string]string{
		"JWT_TOKEN_SECRET": "my_secret_key",
		"REST_PORT":        "8080",
		"GRPC_PORT":        "50052",
		"ADDRESS":          "0.0.0.0",
		"DB_HOST":          "localhost",
		"DB_PORT":          "5432",
		"DB_USER":          "postgres",
		"DB_PASSWORD":      "12345",
		"DB_NAME":          "testpostgres",
	}
	for key, value := range cfg {
		t.Setenv(key, value)
	}
}

func InitData(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository) {
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
		_ = posRepo.Create(context.Background(), &p)
	}
	for _, e := range employees {
		_ = empRepo.Create(context.Background(), &e)
	}
}

func TestData(t *testing.T) {
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	for i := 0; i < 10000; i++ {
		employee := domain.Employee{ID: strconv.Itoa(i), FirstName: strconv.Itoa(i), LastName: strconv.Itoa(i), PositionID: "1"}
		err := employeeRepo.Create(context.Background(), &employee)
		if err != nil {
			log.Println(err)
		}
	}
}

func DeleteData(posRepo domain.PositionsRepository, empRepo domain.EmployeesRepository) {
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

func TestEmployeeRepository_Create(t *testing.T) {
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	InitData(positionRepo, employeeRepo)
	defer DeleteData(positionRepo, employeeRepo)

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
			err := employeeRepo.Create(context.Background(), &tt.employee)
			fmt.Println(err)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}

			_ = employeeRepo.Delete(context.Background(), tt.employee.ID)
		})
	}
}

func TestEmployeeRepository_Get(t *testing.T) {
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	InitData(positionRepo, employeeRepo)
	defer DeleteData(positionRepo, employeeRepo)

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
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	InitData(positionRepo, employeeRepo)
	defer DeleteData(positionRepo, employeeRepo)

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
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	InitData(positionRepo, employeeRepo)
	defer DeleteData(positionRepo, employeeRepo)

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
	SetEnv(t)
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := database.ConnectPostgres(config)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := position.NewPositionsRepository(db)
	employeeRepo := NewEmployeesRepository(db, positionRepo)

	InitData(positionRepo, employeeRepo)
	defer DeleteData(positionRepo, employeeRepo)

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
