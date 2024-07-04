package employee

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/google/uuid"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeesRepository(db *sql.DB) domain.EmployeesRepository {
	return &employeeRepository{
		db: db,
	}
}

var (
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrNothingChanged   = errors.New("nothing changed")
)

const (
	createPosition  = "insert into employees (id, first_name, last_name, position_id, created_at) values ($1, $2, $3, $4, CURRENT_TIMESTAMP);"
	getPosition     = "select first_name, last_name, position_id from employees where id = $1;"
	updatePosition  = "update employees set first_name = $2, last_name = $3, position_id = $4, updated_at = CURRENT_TIMESTAMP where id = $1;"
	deletePositions = "delete from employees where id = $1"
	getAllEmployees = "select id, first_name, last_name, position_id from employees limit $1 offset $2;"
)

func (e *employeeRepository) Create(ctx context.Context, employee domain.Employee) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	employee.ID = uuid.New().String()

	if _, err := e.db.ExecContext(ctx, createPosition, employee.ID, employee.FirstName, employee.LastName, employee.PositionID); err != nil {
		return nil, err
	}
	return &employee, nil
}

func (e *employeeRepository) Get(ctx context.Context, id string) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := e.db.QueryRowContext(ctx, getPosition, id)

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

	res, err := e.db.ExecContext(ctx, updatePosition, employee.ID, employee.FirstName, employee.LastName, employee.PositionID)
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

	res, err := e.db.ExecContext(ctx, deletePositions, id)
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

	rows, err := e.db.QueryContext(ctx, getAllEmployees, pageSize, offset)
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

func (e *employeeRepository) GetByPosition(ctx context.Context, position_id string) (*domain.Employee, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := "select id, first_name, last_name, position_id from employees where position_id = $1;"
	row := e.db.QueryRowContext(ctx, stmt, position_id)

	employee := domain.Employee{}

	err := row.Scan(&employee.FirstName, &employee.LastName, &employee.PositionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEmployeeNotFound
		}
		return nil, err
	}
	return &employee, nil
}
