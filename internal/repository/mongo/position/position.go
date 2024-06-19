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
	collection   *mongo.Collection
	employeeRepo EmployeesRepository
}

type EmployeesRepository interface {
	GetByPosition(ctx context.Context, id string) (*domain.Employee, error)
}

func NewPositionsRepository(c *mongo.Collection) domain.PositionsRepository {
	return &positionsRepository{collection: c}
}

func AddEmpRepo(c *mongo.Collection, empRepo EmployeesRepository) domain.PositionsRepository {
	return &positionsRepository{collection: c, employeeRepo: empRepo}
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
		return nil, errors.Join(err)
	}

	return &position, nil
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":       position.Name,
			"salary":     position.Salary,
			"updated_at": time.Now(),
		},
	}

	res, err := p.collection.UpdateOne(ctx, bson.M{"id": position.ID}, update)
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

	_, err := p.employeeRepo.GetByPosition(ctx, id)
	if err == nil {
		return errors.New("position in use by employee")
	}

	res, err := p.collection.DeleteOne(ctx, bson.M{"id": id})
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

	findOptions := options.Find()
	positions := make([]domain.Position, 0)

	cur, err := p.collection.Find(context.TODO(), bson.D{{}}, findOptions)
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
