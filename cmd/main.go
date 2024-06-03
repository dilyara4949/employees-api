package main

import (
	conf "github.com/dilyara4949/employees-api/internal/config"
	"log"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/repository/storage"
	"github.com/dilyara4949/employees-api/internal/route"
)

func main() {
	positionsStorage := storage.NewPositionsStorage()
	employeesStorage := storage.NewEmployeesStorage()

	positionRepo := position.NewPositionsRepository(positionsStorage)
	positionController := controller.NewPositionsController(positionRepo)

	employeeRepo := employee.NewEmployeesRepository(employeesStorage, positionRepo)
	employeeController := controller.NewEmployeesController(employeeRepo)

	mux := http.NewServeMux()

	config := conf.NewConfig()

	route.SetUpRouter(employeeController, positionController, config, mux)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
