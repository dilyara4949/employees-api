package employee

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

type PositionsRepository interface {
	Get(ctx context.Context, id string) (*domain.Position, error)
}

type employeeRepository struct {
	collection    *mongo.Collection
	positionsRepo PositionsRepository
}

func NewEmployeesRepository(c *mongo.Collection, positionsRepo PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{
		collection:    c,
		positionsRepo: positionsRepo,
	}
}

var (
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrNothingChanged   = errors.New("nothing changed")
)

func (e *employeeRepository) Create(ctx context.Context, employee domain.Employee) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return nil, err
	}

	_, err := e.collection.InsertOne(ctx, bson.M{
		"id":          employee.ID,
		"first_name":  employee.FirstName,
		"last_name":   employee.LastName,
		"position_id": employee.PositionID,
		"created_at":  time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *employeeRepository) Get(ctx context.Context, id string) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	employee := domain.Employee{}
	err := e.collection.FindOne(ctx, bson.M{"id": id}).Decode(&employee)
	if err != nil {
		return nil, errors.Join(err)
	}

	return &employee, nil
}

func (e *employeeRepository) Update(ctx context.Context, employee domain.Employee) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"first_name":  employee.FirstName,
			"last_name":   employee.LastName,
			"position_id": employee.PositionID,
			"updated_at":  time.Now(),
		},
	}

	res, err := e.collection.UpdateOne(ctx, bson.M{"id": employee.ID}, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return ErrEmployeeNotFound
	}

	if res.ModifiedCount == 0 {
		return ErrNothingChanged
	}
	return nil
}

func (e *employeeRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := e.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNothingChanged
	}
	return nil
}

func (e *employeeRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	skip := (page - 1) * pageSize
	findOptions := options.Find().SetSkip(skip).SetLimit(pageSize)

	employees := make([]domain.Employee, 0)

	cur, err := e.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		emp := domain.Employee{}
		err := cur.Decode(&emp)
		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, emp)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *employeeRepository) GetByPosition(ctx context.Context, positionId string) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	employee := domain.Employee{}
	err := e.collection.FindOne(ctx, bson.M{"position_id": positionId}).Decode(&employee)
	if err != nil {
		return nil, errors.Join(err)
	}

	return &employee, nil
}
