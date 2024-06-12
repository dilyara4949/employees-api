package domain

import "context"

type Employee struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	PositionID string `json:"position_id"`
}

type EmployeesRepository interface {
	Create(ctx context.Context, emp *Employee) error
	Get(ctx context.Context, id string) (*Employee, error)
	Update(ctx context.Context, emp Employee) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Employee, error)
}
