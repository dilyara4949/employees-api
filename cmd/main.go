package main

import (
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/route"
)

func main() {
	storageP := &domain.PositionStorage{
		Storage: make(map[string]domain.Position),
	}
	storageE := &domain.EmployeeStorage{
		Storage: make(map[string]domain.Employee),
	}

	positionRepo := position.NewPositionRepository(storageP)
	positionController := controller.NewPositionController(positionRepo)

	employeeRepo := employee.NewEmployeeRepository(storageE, positionRepo)
	employeeController := controller.NewEmployeeController(employeeRepo)

	route.SetUpRouter(employeeController, positionController)
}
