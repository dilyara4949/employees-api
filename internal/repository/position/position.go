package position

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
)

type positionsRepository struct {
	db *sql.DB
}

func NewPositionsRepository(db *sql.DB) domain.PositionsRepository {
	return &positionsRepository{db: db}
}

func (p *positionsRepository) Create(ctx context.Context, position *domain.Position) error {
	position.ID = uuid.New().String()

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
		return nil, errors.New("position does not found")
	case nil:
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
		return errors.New("nothing updated")
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
		return errors.New("nothing deleted")
	}
	return nil
}

func (p *positionsRepository) GetAll(ctx context.Context) ([]domain.Position, error) {
	stmt := "select id, name, salary from positions;"
	rows, err := p.db.Query(stmt)
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
