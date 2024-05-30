package storage

import (
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"sync"
)

type PositionsStorage struct {
	mu      sync.Mutex
	Storage map[string]domain.Position
}

func (storage *PositionsStorage) Add(position domain.Position) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	storage.Storage[position.ID] = position
}

func (storage *PositionsStorage) Get(id string) (*domain.Position, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if position, ok := storage.Storage[id]; ok {
		return &position, nil
	}
	return nil, errors.New("position not found")
}

func (storage *PositionsStorage) Update(position domain.Position) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.Storage[position.ID]; !ok {
		return errors.New("position not found")
	}

	storage.Storage[position.ID] = position
	return nil
}

func (storage *PositionsStorage) Delete(id string) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.Storage[id]; !ok {
		return errors.New("position not found")
	}

	delete(storage.Storage, id)
	return nil
}

func (storage *PositionsStorage) All() []domain.Position {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	var positions []domain.Position

	for _, position := range storage.Storage {
		positions = append(positions, position)
	}
	return positions
}
