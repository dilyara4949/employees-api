package position

import (
	"context"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type positionsRepository struct {
	positionCollection *mongo.Collection
	employeeCollection *mongo.Collection
}

type EmployeesRepository interface {
	GetByPosition(ctx context.Context, id string) (*domain.Employee, error)
}

type positionMongo struct {
	domain.Position
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func NewPositionsRepository(db *mongo.Database, pos, emp string) domain.PositionsRepository {
	return &positionsRepository{
		employeeCollection: db.Collection(emp),
		positionCollection: db.Collection(pos),
	}
}

var (
	ErrPositionNotFound = errors.New("position not found")
	ErrNothingChanged   = errors.New("nothing changed")
)

func (p *positionsRepository) Create(ctx context.Context, position domain.Position) (*domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := p.positionCollection.FindOne(ctx, bson.M{"id": position.ID}).Decode(position)
	if err == nil {
		return nil, errors.New("position already exists")
	}

	_, err = p.positionCollection.InsertOne(ctx, positionMongo{
		Position:  position,
		CreatedAt: time.Now(),
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

	err := p.positionCollection.FindOne(ctx, bson.M{"id": id}).Decode(&position)
	if err != nil {
		return nil, errors.Join(err)
	}

	return &position, nil
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": positionMongo{
			Position:  position,
			UpdatedAt: time.Now(),
		},
	}

	res, err := p.positionCollection.UpdateOne(ctx, bson.M{"id": position.ID}, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return ErrPositionNotFound
	}

	if res.ModifiedCount == 0 {
		return ErrNothingChanged
	}

	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	employee := domain.Employee{}
	err := p.employeeCollection.FindOne(ctx, bson.M{"position_id": id}).Decode(&employee)
	if err == nil {
		return errors.New("position in use by employee")
	}

	res, err := p.positionCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNothingChanged
	}

	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	skip := (page - 1) * pageSize
	findOptions := options.Find().SetSkip(skip).SetLimit(pageSize)

	positions := make([]domain.Position, 0)

	cur, err := p.positionCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		pos := domain.Position{}
		err := cur.Decode(&pos)
		if err != nil {
			log.Fatal(err)
		}

		positions = append(positions, pos)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
