package position

import (
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/storage"

	"github.com/google/uuid"
)

type positionsRepository struct {
	positionsStorage *storage.PositionsStorage
}

func NewPositionsRepository(positionsStorage *storage.PositionsStorage) domain.PositionsRepository {
	return &positionsRepository{positionsStorage: positionsStorage}
}

func (p *positionsRepository) Create(position *domain.Position) error {
	position.ID = uuid.New().String()
	p.positionsStorage.Add(*position)
	return nil
}

func (p *positionsRepository) Get(id string) (*domain.Position, error) {
	return p.positionsStorage.Get(id)
}

func (p *positionsRepository) Update(position domain.Position) error {
	return p.positionsStorage.Update(position)
}

func (p *positionsRepository) Delete(id string) error {
	return p.positionsStorage.Delete(id)
}

func (p *positionsRepository) GetAll() ([]domain.Position, error) {
	return p.positionsStorage.All()
}
