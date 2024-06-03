package route

import (
	"github.com/dilyara4949/employees-api/internal/config"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, config config.Config, mux *http.ServeMux) {

	auth := middleware.NewJWTAuth(config.JWTTokenSecret)

	mux.HandleFunc("GET /positions/{id}", auth.Auth(positionsController.GetPosition))
	mux.HandleFunc("POST /positions", auth.Auth(positionsController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", auth.Auth(positionsController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", auth.Auth(positionsController.UpdatePosition))
	mux.HandleFunc("GET /positions", auth.Auth(positionsController.GetAllPositions))

	mux.HandleFunc("GET /employees/{id}", auth.Auth(employeesController.GetEmployee))
	mux.HandleFunc("POST /employees", auth.Auth(employeesController.CreateEmployee))
	mux.HandleFunc("DELETE /employees/{id}", auth.Auth(employeesController.DeleteEmployee))
	mux.HandleFunc("PUT /employees/{id}", auth.Auth(employeesController.UpdateEmployee))
	mux.HandleFunc("GET /employees", auth.Auth(employeesController.GetAllEmployees))
}
