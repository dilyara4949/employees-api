package main

import (
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/repository/storage"
	"github.com/dilyara4949/employees-api/internal/route"
	"log"
	"net/http"
)

func main() {
	storageP := &storage.PositionsStorage{
		Storage: make(map[string]domain.Position),
	}
	storageE := &storage.EmployeesStorage{
		Storage: make(map[string]domain.Employee),
	}

	positionRepo := position.NewPositionsRepository(storageP)
	positionController := controller.NewPositionsController(positionRepo)

	employeeRepo := employee.NewEmployeesRepository(storageE, positionRepo)
	employeeController := controller.NewEmployeesController(employeeRepo)

	mux := http.NewServeMux()

	route.SetUpRouter(employeeController, positionController, mux)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
