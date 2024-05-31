package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal"
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, env *internal.Env, mux *http.ServeMux) {
	mux.HandleFunc("GET /positions/{id}", middleware.JwtMiddleware(positionsController.GetPosition, env.JWTTokenSecret))
	mux.HandleFunc("POST /positions", middleware.JwtMiddleware(positionsController.CreatePosition, env.JWTTokenSecret))
	mux.HandleFunc("DELETE /positions/{id}", middleware.JwtMiddleware(positionsController.DeletePosition, env.JWTTokenSecret))
	mux.HandleFunc("PUT /positions/{id}", middleware.JwtMiddleware(positionsController.UpdatePosition, env.JWTTokenSecret))
	mux.HandleFunc("GET /positions", middleware.JwtMiddleware(positionsController.GetAllPositions, env.JWTTokenSecret))

	mux.HandleFunc("GET /employees/{id}", middleware.JwtMiddleware(employeesController.GetEmployee, env.JWTTokenSecret))
	mux.HandleFunc("POST /employees", middleware.JwtMiddleware(employeesController.CreateEmployee, env.JWTTokenSecret))
	mux.HandleFunc("DELETE /employees/{id}", middleware.JwtMiddleware(employeesController.DeleteEmployee, env.JWTTokenSecret))
	mux.HandleFunc("PUT /employees/{id}", middleware.JwtMiddleware(employeesController.UpdateEmployee, env.JWTTokenSecret))
	mux.HandleFunc("GET /employees", middleware.JwtMiddleware(employeesController.GetAllEmployees, env.JWTTokenSecret))
}
