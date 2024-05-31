package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal"
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, env *internal.Env, mux *http.ServeMux) {

	mux.HandleFunc("GET /positions/{id}", middleware.Adapt(middleware.JwtMiddleware(positionsController.GetPosition, env.JWTTokenSecret),
		middleware.СorrelationIDMiddleware(),
		middleware.Logger()))
	mux.HandleFunc("POST /positions", middleware.Adapt(middleware.JwtMiddleware(positionsController.CreatePosition, env.JWTTokenSecret),
		middleware.СorrelationIDMiddleware(),
		middleware.Logger()))
	mux.HandleFunc("DELETE /positions/{id}", middleware.Adapt(middleware.JwtMiddleware(positionsController.DeletePosition, env.JWTTokenSecret),
		middleware.СorrelationIDMiddleware(),
		middleware.Logger()))
	mux.HandleFunc("PUT /positions/{id}", middleware.Adapt(middleware.JwtMiddleware(positionsController.UpdatePosition, env.JWTTokenSecret),
		middleware.СorrelationIDMiddleware(),
		middleware.Logger()))
	mux.HandleFunc("GET /positions", middleware.Adapt(middleware.JwtMiddleware(positionsController.GetAllPositions, env.JWTTokenSecret),
		middleware.СorrelationIDMiddleware(),
		middleware.Logger()))

	mux.HandleFunc("GET /employees/{id}", middleware.Adapt(middleware.JwtMiddleware(employeesController.GetEmployee, env.JWTTokenSecret), middleware.Logger()))
	mux.HandleFunc("POST /employees", middleware.Adapt(middleware.JwtMiddleware(employeesController.CreateEmployee, env.JWTTokenSecret), middleware.Logger()))
	mux.HandleFunc("DELETE /employees/{id}", middleware.Adapt(middleware.JwtMiddleware(employeesController.DeleteEmployee, env.JWTTokenSecret), middleware.Logger()))
	mux.HandleFunc("PUT /employees/{id}", middleware.Adapt(middleware.JwtMiddleware(employeesController.UpdateEmployee, env.JWTTokenSecret), middleware.Logger()))
	mux.HandleFunc("GET /employees", middleware.Adapt(middleware.JwtMiddleware(employeesController.GetAllEmployees, env.JWTTokenSecret), middleware.Logger()))

}
