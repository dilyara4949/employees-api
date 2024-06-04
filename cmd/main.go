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

	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	route.SetUpRouter(employeeController, positionController, config, mux)

	log.Printf("Starting server on :%s", config.Port)
	err = http.ListenAndServe(":"+config.Port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
