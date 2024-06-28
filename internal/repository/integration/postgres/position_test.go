//go:build integration
// +build integration

package postgres

import (
	"context"
	"errors"
	"fmt"
	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/database/postgres"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/position"
	"log"
	"reflect"
	"testing"
)

func initDataPos(posRepo domain.PositionsRepository) ([]domain.Position, error) {
	positions := []domain.Position{
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
	var errs error
	for i := 0; i < len(positions); i++ {
		pos, err := posRepo.Create(context.Background(), positions[i])
		if err != nil {
			errs = errors.Join(errs, err)
		}
		positions[i] = *pos
	}
	return positions, errs
}

func deleteDataPos(posRepo domain.PositionsRepository, positions []domain.Position) error {
	var errs error
	for i := 0; i < len(positions); i++ {
		err := posRepo.Delete(context.Background(), positions[i].ID)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return errs
}

func InitPosRepo() (domain.PositionsRepository, error) {
	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	var positionRepo domain.PositionsRepository

	db, err := postgres.ConnectPostgres(config.PostgresConfig)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}

	positionRepo = position.NewPositionsRepository(db)

	return positionRepo, nil
}

func TestPositionRepository_Create(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, err := initDataPos(positionRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := deleteDataPos(positionRepo, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

	tests := []struct {
		name        string
		position    domain.Position
		expectedErr bool
	}{
		{
			name: "OK",
			position: domain.Position{
				Name:   "name4",
				Salary: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, err := positionRepo.Create(context.Background(), tt.position)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
			}
			if pos != nil {
				tt.position.ID = pos.ID
			}
			if pos != nil && !reflect.DeepEqual(*pos, tt.position) {
				t.Errorf("expected: %v, got: %v", tt.position, *pos)
			}
			err = positionRepo.Delete(context.Background(), tt.position.ID)
			if err != nil {
				fmt.Println("error at deleting created employee: v", err, tt.position, pos)
			}
		})
	}
}

func TestPositionRepository_Get(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, err := initDataPos(positionRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := deleteDataPos(positionRepo, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

	tests := []struct {
		name        string
		positions   []domain.Position
		expectedErr bool
	}{
		{
			name:      "OK",
			positions: poss,
		},
		{
			name: "position not found",
			positions: []domain.Position{
				{
					ID:     "4",
					Name:   "name3",
					Salary: 3},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, position := range tt.positions {
				pos, err := positionRepo.Get(context.Background(), position.ID)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}

				if !tt.expectedErr && !reflect.DeepEqual(*pos, position) {
					t.Errorf("expected: %v, got: %v", position, *pos)
				}
			}
		})
	}
}

func TestPositionRepository_Update(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, err := initDataPos(positionRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := deleteDataPos(positionRepo, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()
	tests := []struct {
		name        string
		positions   []domain.Position
		expectedErr bool
	}{
		{
			name:      "OK",
			positions: poss,
		},
		{
			name: "position not found",
			positions: []domain.Position{
				{
					ID:     "4",
					Name:   "name3",
					Salary: 3,
				},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, position := range tt.positions {
				err := positionRepo.Update(context.Background(), position)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}
			}
		})
	}
}

func TestPositionRepository_Delete(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, err := initDataPos(positionRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}

	tests := []struct {
		name        string
		positions   []domain.Position
		expectedErr bool
	}{
		{
			name:      "OK",
			positions: poss,
		},
		{
			name: "position not found",
			positions: []domain.Position{
				{ID: "5"},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, position := range tt.positions {
				err := positionRepo.Delete(context.Background(), position.ID)
				if (err != nil) != tt.expectedErr {
					t.Errorf("expected error: %v, got: %s", tt.expectedErr, err)
				}
			}
		})
	}
}

func TestPositionRepository_GetAll(t *testing.T) {
	positionRepo, err := InitPosRepo()
	if err != nil {
		t.Fatal(err)
	}

	poss, err := initDataPos(positionRepo)
	if err != nil {
		t.Errorf("error to init data: %v", err)
	}
	defer func() {
		err := deleteDataPos(positionRepo, poss)
		if err != nil {
			t.Errorf("error at deleting helper data: %v", err)
		}
	}()

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
				poss[0],
				poss[1],
			},
		},
		{
			name:     "Second page, one record",
			page:     2,
			pageSize: 2,
			expected: []domain.Position{
				poss[2],
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
