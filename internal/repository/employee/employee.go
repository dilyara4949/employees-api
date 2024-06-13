package employee

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"

	"github.com/google/uuid"
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

func (e *employeeRepository) Create(ctx context.Context, employee *domain.Employee) error {
	if _, err := e.positionsRepo.Get(ctx, employee.PositionID); err != nil {
		return err
	}

	employee.ID = uuid.New().String()

	stmt := "insert into positions (id, first_name, last_name, position_id, created_at) values ($1, $2, $3, $4, CURRENT_TIMESTAMP);"

	if _, err := e.db.Exec(stmt, employee.ID, employee.FirstName, employee.LastName, employee.PositionID); err != nil {
		return err
	}
	return nil
}

func (e *employeeRepository) Get(_ context.Context, id string) (*domain.Employee, error) {
	stmt := "select first_name, last_mame, position_id from employees where id = $1;"
	row := e.db.QueryRow(stmt, id)

	employee := domain.Employee{}
	switch err := row.Scan(&employee.FirstName, &employee.LastName, &employee.PositionID); err {
	case sql.ErrNoRows:
		return nil, errors.New("employee does not found")
	case nil:
		return &employee, nil
	default:
		return nil, err
	}
}

// update positions
// set name = $2, salary = $3, updated_at = CURRENT_TIMESTAMP
// where id = $1;

func (e *employeeRepository) Update(ctx context.Context, employee domain.Employee) error {
	stmt := "update employees set first_name = $2, last_name = $3, position_id = $4, updated_at = CURRENT_TIMESTAMP where id = $1;"

	res, err := e.db.Exec(stmt, employee.ID, employee.FirstName, employee.LastName, employee.PositionID)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return errors.New("nothing updated")
	}
	return nil
}

//  delete from employees where id = $1;

func (e *employeeRepository) Delete(_ context.Context, id string) error {
	stmt := "delete from employees where id = $1"

	res, err := e.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt != 1 {
		return errors.New("nothing deleted")
	}
	return nil
}

// select id, first_name, last_name, employee_id, created_at, updated_at from positions;

func (e *employeeRepository) GetAll(_ context.Context) ([]domain.Employee, error) {
	stmt := "select id, name, salary from employees;"
	rows, err := e.db.Query(stmt)
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
