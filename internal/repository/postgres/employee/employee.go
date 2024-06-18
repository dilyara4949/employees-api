package employee

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"time"
)

type PositionsRepository interface {
	Get(ctx context.Context, id string) (*domain.Position, error)
}

type employeeRepository struct {
	db            *sql.DB
	positionsRepo PositionsRepository
}

func NewEmployeesRepository(db *sql.DB, positionsRepo PositionsRepository) domain.EmployeesRepository {
	return &employeeRepository{
		db:            db,
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

	stmt := "insert into employees (id, first_name, last_name, position_id, created_at) values ($1, $2, $3, $4, CURRENT_TIMESTAMP);"

	if _, err := e.db.ExecContext(ctx, stmt, employee.ID, employee.FirstName, employee.LastName, employee.PositionID); err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *employeeRepository) Get(ctx context.Context, id string) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := "select first_name, last_name, position_id from employees where id = $1;"
	row := e.db.QueryRowContext(ctx, stmt, id)

	employee := domain.Employee{}

	err := row.Scan(&employee.FirstName, &employee.LastName, &employee.PositionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEmployeeNotFound
		}
		return nil, err
	}

	employee.ID = id
	return &employee, nil
}

func (e *employeeRepository) Update(ctx context.Context, employee domain.Employee) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := "update employees set first_name = $2, last_name = $3, position_id = $4, updated_at = CURRENT_TIMESTAMP where id = $1;"

	res, err := e.db.ExecContext(ctx, stmt, employee.ID, employee.FirstName, employee.LastName, employee.PositionID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (e *employeeRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := "delete from employees where id = $1"

	res, err := e.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return ErrNothingChanged
	}
	return nil
}

func (e *employeeRepository) GetAll(ctx context.Context, page, pageSize int64) ([]domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	offset := (page - 1) * pageSize

	stmt := "select id, first_name, last_name, position_id from employees limit $1 offset $2;"
	rows, err := e.db.QueryContext(ctx, stmt, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]domain.Employee, 0)
	for rows.Next() {
		employee := domain.Employee{}

		err = rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.PositionID)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
