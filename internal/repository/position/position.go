package position

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type positionRepository struct {
	db *domain.PositionStorage
}

func NewPositionRepository(db *domain.PositionStorage) domain.PositionRepository {
	return &positionRepository{db: db}
}

func (p *positionRepository) Create(position *domain.Position) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	position.ID = uuid.New().String()
	p.db.Storage[position.ID] = *position

	return nil
}

func (p *positionRepository) Get(id string) (*domain.Position, error) {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return nil, fmt.Errorf("position does not exist")
	}

	position := p.db.Storage[id]
	return &position, nil
}

func (p *positionRepository) Update(position domain.Position) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[position.ID]; !ok {
		return fmt.Errorf("position does not exist")
	}

	p.db.Storage[position.ID] = position
	return nil
}

func (p *positionRepository) Delete(id string) error {
	p.db.Mu.Lock()
	defer p.db.Mu.Unlock()

	if _, ok := p.db.Storage[id]; !ok {
		return fmt.Errorf("position does not exist")
	}

	delete(p.db.Storage, id)
	return nil
}

func (p *positionRepository) GetAll() ([]domain.Position, error) {
	return nil, nil
}
