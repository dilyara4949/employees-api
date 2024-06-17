package route

import (
	conf "github.com/dilyara4949/employees-api/internal/config"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, config conf.Config, mux *http.ServeMux) {

	mux.HandleFunc("GET /positions/{id}", logCorrelationIDTimer(positionsController.GetPosition, config.JWTTokenSecret))
	mux.HandleFunc("POST /positions", logCorrelationIDTimer(positionsController.CreatePosition, config.JWTTokenSecret))
	mux.HandleFunc("DELETE /positions/{id}", logCorrelationIDTimer(positionsController.DeletePosition, config.JWTTokenSecret))
	mux.HandleFunc("PUT /positions/{id}", logCorrelationIDTimer(positionsController.UpdatePosition, config.JWTTokenSecret))
	mux.HandleFunc("GET /positions", logCorrelationIDTimer(positionsController.GetAllPositions, config.JWTTokenSecret))

	mux.HandleFunc("GET /employees/{id}", logCorrelationIDTimer(employeesController.GetEmployee, config.JWTTokenSecret))
	mux.HandleFunc("POST /employees", logCorrelationIDTimer(employeesController.CreateEmployee, config.JWTTokenSecret))
	mux.HandleFunc("DELETE /employees/{id}", logCorrelationIDTimer(employeesController.DeleteEmployee, config.JWTTokenSecret))
	mux.HandleFunc("PUT /employees/{id}", logCorrelationIDTimer(employeesController.UpdateEmployee, config.JWTTokenSecret))
	mux.HandleFunc("GET /employees", logCorrelationIDTimer(employeesController.GetAllEmployees, config.JWTTokenSecret))
}

func logCorrelationIDTimer(endpoint http.HandlerFunc, JWTTokenSecret string) http.HandlerFunc {
	auth := middleware.NewJWTAuth(JWTTokenSecret)
	middlewares := []middleware.Middleware{
		auth.Auth(),
		middleware.Timer(),
		middleware.Logger(),
		middleware.CorrelationIDMiddleware(),
	}

	return middleware.Chain(endpoint, middlewares...)
}
