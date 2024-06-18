//go:build integration
// +build integration

package position

import (
	"context"
	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/database/postgres"
	"github.com/dilyara4949/employees-api/internal/domain"
	"log"
	"reflect"
	"testing"
)

func InitData(posRepo domain.PositionsRepository) {
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

func DeleteData(posRepo domain.PositionsRepository) {
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

func TestPositionRepository_Create(t *testing.T) {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := postgres.ConnectPostgres(config.DB)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := NewPositionsRepository(db)

	InitData(positionRepo)
	defer DeleteData(positionRepo)

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
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := postgres.ConnectPostgres(config.DB)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := NewPositionsRepository(db)

	InitData(positionRepo)
	defer DeleteData(positionRepo)

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

			if !tt.expectedErr && !reflect.DeepEqual(*pos, tt.position) {
				t.Errorf("expected: %v, got: %v", tt.position, *pos)
			}
		})
	}
}

func TestPositionRepository_Update(t *testing.T) {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := postgres.ConnectPostgres(config.DB)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := NewPositionsRepository(db)

	InitData(positionRepo)
	defer DeleteData(positionRepo)

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
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := postgres.ConnectPostgres(config.DB)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := NewPositionsRepository(db)

	InitData(positionRepo)
	defer DeleteData(positionRepo)

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
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	db, err := postgres.ConnectPostgres(config.DB)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}
	defer db.Close()

	positionRepo := NewPositionsRepository(db)

	InitData(positionRepo)
	defer DeleteData(positionRepo)

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
			name:     "Second page, two records",
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
			employees, err := positionRepo.GetAll(context.Background(), tt.page, tt.pageSize)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if !reflect.DeepEqual(employees, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, employees)
			}
		})
	}
}
