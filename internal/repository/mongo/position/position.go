package position

import (
	"context"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type positionsRepository struct {
	collection *mongo.Collection
}

func NewPositionsRepository(c *mongo.Collection) domain.PositionsRepository {
	return &positionsRepository{collection: c}
}

var (
	ErrPositionNotFound = errors.New("position not found")
	ErrNothingChanged   = errors.New("nothing changed")
)

func (p *positionsRepository) Create(ctx context.Context, position domain.Position) (*domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.collection.InsertOne(ctx, bson.M{
		"id":         position.ID,
		"name":       position.Name,
		"salary":     position.Salary,
		"created_at": time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (p *positionsRepository) Get(ctx context.Context, id string) (*domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	position := domain.Position{}

	err := p.collection.FindOne(ctx, bson.M{"id": id}).Decode(&position)
	if err != nil {
		return nil, ErrPositionNotFound
	}

	return &position, nil
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil, nil
}
