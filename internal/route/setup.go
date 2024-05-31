package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, mux *http.ServeMux) {
	mux.HandleFunc("GET /positions/{id}", middleware.JwtMiddleware(positionsController.GetPosition))
	mux.HandleFunc("POST /positions", middleware.JwtMiddleware(positionsController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", middleware.JwtMiddleware(positionsController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", middleware.JwtMiddleware(positionsController.UpdatePosition))
	mux.HandleFunc("GET /positions", middleware.JwtMiddleware(positionsController.GetAllPositions))

	mux.HandleFunc("GET /employees/{id}", middleware.JwtMiddleware(employeesController.GetEmployee))
	mux.HandleFunc("POST /employees", middleware.JwtMiddleware(employeesController.CreateEmployee))
	mux.HandleFunc("DELETE /employees/{id}", middleware.JwtMiddleware(employeesController.DeleteEmployee))
	mux.HandleFunc("PUT /employees/{id}", middleware.JwtMiddleware(employeesController.UpdateEmployee))
	mux.HandleFunc("GET /employees", middleware.JwtMiddleware(employeesController.GetAllEmployees))
}
