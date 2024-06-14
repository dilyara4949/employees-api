package position

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
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

func (p *positionsRepository) Create(ctx context.Context, position *domain.Position) error {
	stmt := "insert into positions (id, name, salary, created_at) values ($1, $2, $3, CURRENT_TIMESTAMP);"
	if _, err := p.db.Exec(stmt, position.ID, position.Name, position.Salary); err != nil {
		return err
	}
	return nil
}

func (p *positionsRepository) Get(ctx context.Context, id string) (*domain.Position, error) {
	stmt := "select name, salary from positions where id = $1;"
	row := p.db.QueryRow(stmt, id)
	position := domain.Position{}

	switch err := row.Scan(&position.Name, &position.Salary); err {
	case sql.ErrNoRows:
		return nil, ErrPositionNotFound
	case nil:
		position.ID = id
		return &position, nil
	default:
		return nil, err
	}
}

func (p *positionsRepository) Update(ctx context.Context, position domain.Position) error {
	stmt := "update positions set name = $2, salary = $3, updated_at = CURRENT_TIMESTAMP where id = $1;"

	res, err := p.db.Exec(stmt, position.ID, position.Name, position.Salary)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id string) error {
	stmt := "delete from positions where id = $1"

	res, err := p.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Position, error) {
	offset := (page - 1) * pageSize

	stmt := "select id, name, salary from positions limit $1 offset $2;"
	rows, err := p.db.Query(stmt, pageSize, offset)
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
