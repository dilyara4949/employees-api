//go:build integration
// +build integration

package integration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	conf "github.com/dilyara4949/employees-api/internal/config"
	mongoDB "github.com/dilyara4949/employees-api/internal/database/mongo"
	"github.com/dilyara4949/employees-api/internal/database/postgres"
	"github.com/dilyara4949/employees-api/internal/domain"
	mongoposition "github.com/dilyara4949/employees-api/internal/repository/mongo/position"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/position"
)

func InitDataPositions(posRepo domain.PositionsRepository) {
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

	for _, p := range positions {
		_, _ = posRepo.Create(context.Background(), p)
	}
}

func DeleteDataPositions(posRepo domain.PositionsRepository) {
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
	for _, p := range positions {
		_ = posRepo.Delete(context.Background(), p.ID)
	}
}

func InitPosRepo() (domain.PositionsRepository, error) {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	var positionRepo domain.PositionsRepository

	switch config.DatabaseType {
	case "testpostgres":
		db, err := postgres.ConnectPostgres(config.PostgresConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}

		positionRepo = position.NewPositionsRepository(db)

	case "testmongo":
		db, err := mongoDB.ConnectMongo(config.MongoConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}

		positionRepo = mongoposition.NewPositionsRepository(db, config.MongoConfig.Collections.Positions, config.MongoConfig.Collections.Employees)
	default:
		return nil, errors.New("Incorrect database given for tests")
	}
	return positionRepo, nil
}

func TestPositionRepository_Create(t *testing.T) {

	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataPositions(positionRepo)
	defer DeleteDataPositions(positionRepo)

	tests := []struct {
		name        string
		position    domain.Position
		expectedErr bool
	}{
		{
			name: "OK",
			position: domain.Position{
				ID:     "4",
				Name:   "name4",
				Salary: 4,
			},
		},
		{
			name: "position already exists",
			position: domain.Position{
				ID:     "1",
				Name:   "name1",
				Salary: 1,
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, err := positionRepo.Create(context.Background(), tt.position)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}

			if pos != nil && !reflect.DeepEqual(*pos, tt.position) {
				t.Errorf("expected: %v, got: %v", tt.position, &pos)
			}

			_ = positionRepo.Delete(context.Background(), tt.position.ID)
		})
	}
}

func TestPositionRepository_Get(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataPositions(positionRepo)
	defer DeleteDataPositions(positionRepo)

	tests := []struct {
		name        string
		position    domain.Position
		expectedErr bool
	}{
		{
			name: "OK",
			position: domain.Position{
				ID:     "3",
				Name:   "name3",
				Salary: 3,
			},
		},
		{
			name: "position not found",
			position: domain.Position{
				ID:     "4",
				Name:   "name3",
				Salary: 3,
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pos, err := positionRepo.Get(context.Background(), tt.position.ID)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}

			fmt.Println("-----------------------", tt.expectedErr, err, pos, tt.position)

			if !tt.expectedErr && !reflect.DeepEqual(*pos, tt.position) {
				t.Errorf("expected: %v, got: %v", tt.position, *pos)
			}
		})
	}
}

func TestPositionRepository_Update(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataPositions(positionRepo)
	defer DeleteDataPositions(positionRepo)

	tests := []struct {
		name        string
		position    domain.Position
		expectedErr bool
	}{
		{
			name: "OK",
			position: domain.Position{
				ID:     "3",
				Name:   "name3qw",
				Salary: 33,
			},
		},
		{
			name: "position not found",
			position: domain.Position{
				ID:     "4",
				Name:   "name3",
				Salary: 3,
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := positionRepo.Update(context.Background(), tt.position)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestPositionRepository_Delete(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataPositions(positionRepo)
	defer DeleteDataPositions(positionRepo)

	tests := []struct {
		name        string
		position    domain.Position
		expectedErr bool
	}{
		{
			name: "OK",
			position: domain.Position{
				ID: "1",
			},
		},
		{
			name: "position not found",
			position: domain.Position{
				ID: "5",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := positionRepo.Delete(context.Background(), tt.position.ID)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
		})
	}
}

func TestPositionRepository_GetAll(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	InitDataPositions(positionRepo)
	defer DeleteDataPositions(positionRepo)

	tests := []struct {
		name     string
		page     int64
		pageSize int64
		expected []domain.Position
	}{
		{
			name:     "First page, two records",
			page:     1,
			pageSize: 2,
			expected: []domain.Position{
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
			},
		},
		{
			name:     "Second page, one record",
			page:     2,
			pageSize: 2,
			expected: []domain.Position{
				{
					ID:     "3",
					Name:   "name3",
					Salary: 3,
				},
			},
		},
		{
			name:     "Page out of range",
			page:     3,
			pageSize: 2,
			expected: []domain.Position{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			positions, err := positionRepo.GetAll(context.Background(), tt.page, tt.pageSize)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(positions, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, positions)
			}
		})
	}
}
