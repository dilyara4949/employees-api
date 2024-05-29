package position

import (
	"fmt"

	"github.com/google/uuid"
)

type positionRepository struct {
	db *Storage
}

func NewPositionRepository(db *Storage) Repository {
	return &positionRepository{db: db}
}

func (p *positionRepository) Create(position *Position) error {
	p.db.mu.Lock()
	defer p.db.mu.Unlock()

	position.ID = uuid.New().String()
	p.db.Storage[position.ID] = *position

	return nil
}

func (p *positionRepository) Get(id string) (*Position, error) {
	p.db.mu.Lock()
	defer p.db.mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return nil, fmt.Errorf("position does not exist")
	}

	position := p.db.Storage[id]
	return &position, nil
}

func (p *positionRepository) Update(position Position) error {
	p.db.mu.Lock()
	defer p.db.mu.Unlock()

	if _, ok := p.db.Storage[position.ID]; !ok {
		return fmt.Errorf("position does not exist")
	}

	p.db.Storage[position.ID] = position
	return nil
}

func (p *positionRepository) Delete(id string) error {
	p.db.mu.Lock()
	defer p.db.mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return fmt.Errorf("position does not exist")
	}

	delete(p.db.Storage, id)
	return nil
}

func (p *positionRepository) GetAll() ([]Position, error) {
	return nil, nil
}
