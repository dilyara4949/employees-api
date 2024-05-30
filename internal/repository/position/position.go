package position

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type positionsRepository struct {
	db *domain.PositionsStorage
}

func NewPositionsRepository(db *domain.PositionsStorage) domain.PositionsRepository {
	return &positionsRepository{db: db}
}

func (p *positionsRepository) Create(position *domain.Position) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	position.ID = uuid.New().String()
	p.db.Storage[position.ID] = *position

	return nil
}

func (p *positionsRepository) Get(id string) (*domain.Position, error) {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return nil, fmt.Errorf("position does not exist")
	}

	position := p.db.Storage[id]
	return &position, nil
}

func (p *positionsRepository) Update(position domain.Position) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[position.ID]; !ok {
		return fmt.Errorf("position does not exist")
	}

	p.db.Storage[position.ID] = position
	return nil
}

func (p *positionsRepository) Delete(id string) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return fmt.Errorf("position does not exist")
	}

	delete(p.db.Storage, id)
	return nil
}

func (p *positionsRepository) GetAll() ([]domain.Position, error) {
	positions := make([]domain.Position, 0)

	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	for _, v := range p.db.Storage {
		positions = append(positions, v)
	}
	return positions, nil
}
