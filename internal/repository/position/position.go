package position

import (
	"context"
	"errors"
	"sync"

	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type positionsRepository struct {
	mu      sync.Mutex
	storage map[string]domain.Position
}

func NewPositionsRepository() domain.PositionsRepository {
	return &positionsRepository{storage: make(map[string]domain.Position)}
}

func (p *positionsRepository) Create(ctx context.Context, position *domain.Position) error {
	position.ID = uuid.New().String()

	p.mu.Lock()
	defer p.mu.Unlock()

	p.storage[position.ID] = *position
	return nil
}

func (p *positionsRepository) Get(ctx context.Context, id string) (*domain.Position, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if position, ok := p.storage[id]; ok {
		return &position, nil
	}
	return nil, errors.New("position not found")
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.storage[position.ID]; !ok {
		return errors.New("position not found")
	}

	p.storage[position.ID] = position
	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.storage[id]; !ok {
		return errors.New("position not found")
	}

	delete(p.storage, id)
	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context) []domain.Position {
	p.mu.Lock()
	defer p.mu.Unlock()

	positions := make([]domain.Position, 0)

	for _, position := range p.storage {
		positions = append(positions, position)
	}
	return positions
}
