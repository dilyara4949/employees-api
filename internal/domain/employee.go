package domain

import "context"

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname" bson:"first_name"`
	LastName   string `json:"lastname" bson:"last_name"`
	PositionID string `json:"position_id" bson:"position_id"`
}

type EmployeesRepository interface {
	Create(ctx context.Context, emp Employee) (*Employee, error)
	Get(ctx context.Context, id string) (*Employee, error)
	Update(ctx context.Context, emp Employee) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, page, pageSize int64) ([]Employee, error)
	GetByPosition(ctx context.Context, positionId string) (*Employee, error)
}
