package position

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/google/uuid"
)

type positionsRepository struct {
	db *sql.DB
}

func NewPositionsRepository(db *sql.DB) domain.PositionsRepository {
	return &positionsRepository{db: db}
}

var (
	ErrPositionNotFound = errors.New("position not found")
	ErrNothingChanged   = errors.New("nothing changed")
)

const (
	createPosition  = "insert into positions (id, name, salary, created_at) values ($1, $2, $3, CURRENT_TIMESTAMP);"
	getPosition     = "select name, salary from positions where id = $1;"
	updatePosition  = "update positions set name = $2, salary = $3, updated_at = CURRENT_TIMESTAMP where id = $1;"
	deletePosition  = "delete from positions where id = $1"
	getAllPositions = "select id, name, salary from positions limit $1 offset $2;"
)

func (p *positionsRepository) Create(ctx context.Context, position domain.Position) (*domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	position.ID = uuid.New().String()

	if _, err := p.db.ExecContext(ctx, createPosition, position.ID, position.Name, position.Salary); err != nil {
		return nil, err
	}
	return &position, nil
}

func (p *positionsRepository) Get(ctx context.Context, id string) (*domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := p.db.QueryRowContext(ctx, getPosition, id)
	position := domain.Position{}

	err := row.Scan(&position.Name, &position.Salary)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPositionNotFound
		}
		return nil, err
	}

	position.ID = id
	return &position, nil
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := p.db.ExecContext(ctx, updatePosition, position.ID, position.Name, position.Salary)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := p.db.ExecContext(ctx, deletePosition, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	offset := (page - 1) * pageSize

	rows, err := p.db.QueryContext(ctx, getAllPositions, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	positions := make([]domain.Position, 0)
	for rows.Next() {
		position := domain.Position{}

		err = rows.Scan(&position.ID, &position.Name, &position.Salary)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}
