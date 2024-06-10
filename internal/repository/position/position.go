package position

import (
	"context"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type positionsRepository struct {
	positionsStorage *PositionsStorage
}

func NewPositionsRepository(positionsStorage *PositionsStorage) domain.PositionsRepository {
	return &positionsRepository{positionsStorage: positionsStorage}
}

type PositionsStorage struct {
	mu      sync.Mutex
	storage map[string]domain.Position
}

func NewPositionsStorage() *PositionsStorage {
	return &PositionsStorage{
		storage: make(map[string]domain.Position),
	}
}

func (p *positionsRepository) Create(ctx context.Context, position *domain.Position) error {
	position.ID = uuid.New().String()
	p.positionsStorage.mu.Lock()
	defer p.positionsStorage.mu.Unlock()
	p.positionsStorage.storage[position.ID] = *position
	return nil
}

func (p *positionsRepository) Get(ctx context.Context, id string) (*domain.Position, error) {
	p.positionsStorage.mu.Lock()
	defer p.positionsStorage.mu.Unlock()

	if position, ok := p.positionsStorage.storage[id]; ok {
		return &position, nil
	}
	return nil, errors.New("position not found")
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	p.positionsStorage.mu.Lock()
	defer p.positionsStorage.mu.Unlock()

	if _, ok := p.positionsStorage.storage[position.ID]; !ok {
		return errors.New("position not found")
	}

	p.positionsStorage.storage[position.ID] = position
	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	p.positionsStorage.mu.Lock()
	defer p.positionsStorage.mu.Unlock()

	if _, ok := p.positionsStorage.storage[id]; !ok {
		return errors.New("position not found")
	}

	delete(p.positionsStorage.storage, id)
	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context) []domain.Position {
	p.positionsStorage.mu.Lock()
	defer p.positionsStorage.mu.Unlock()

	positions := make([]domain.Position, 0)

	for _, position := range p.positionsStorage.storage {
		positions = append(positions, position)
	}
	return positions
}
