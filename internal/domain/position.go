package domain

import "context"

type Position struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type PositionsRepository interface {
	Create(ctx context.Context, pos *Position) error
	Get(ctx context.Context, id string) (*Position, error)
	Update(ctx context.Context, pos Position) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Position, error)
}
