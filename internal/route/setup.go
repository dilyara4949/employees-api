package route

import (
	conf "github.com/dilyara4949/employees-api/internal/config"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, config conf.Config, mux *http.ServeMux) {

	//auth := middleware.NewJWTAuth(config.JWTTokenSecret)

	mux.HandleFunc("GET /positions/{id}", LogCorrelationIDTimer(positionsController.GetPosition, config.JWTTokenSecret))
	mux.HandleFunc("POST /positions", LogCorrelationIDTimer(positionsController.CreatePosition, config.JWTTokenSecret))
	mux.HandleFunc("DELETE /positions/{id}", LogCorrelationIDTimer(positionsController.DeletePosition, config.JWTTokenSecret))
	mux.HandleFunc("PUT /positions/{id}", LogCorrelationIDTimer(positionsController.UpdatePosition, config.JWTTokenSecret))
	mux.HandleFunc("GET /positions", LogCorrelationIDTimer(positionsController.GetAllPositions, config.JWTTokenSecret))

	mux.HandleFunc("GET /employees/{id}", LogCorrelationIDTimer(employeesController.GetEmployee, config.JWTTokenSecret))
	mux.HandleFunc("POST /employees", LogCorrelationIDTimer(employeesController.CreateEmployee, config.JWTTokenSecret))
	mux.HandleFunc("DELETE /employees/{id}", LogCorrelationIDTimer(employeesController.DeleteEmployee, config.JWTTokenSecret))
	mux.HandleFunc("PUT /employees/{id}", LogCorrelationIDTimer(employeesController.UpdateEmployee, config.JWTTokenSecret))
	mux.HandleFunc("GET /employees", LogCorrelationIDTimer(employeesController.GetAllEmployees, config.JWTTokenSecret))
}

func LogCorrelationIDTimer(endpoint http.HandlerFunc, JWTTokenSecret string) http.HandlerFunc {
	auth := middleware.NewJWTAuth(JWTTokenSecret)
	middlewares := []middleware.Middleware{
		auth.Auth(),
		middleware.Logger(),
		middleware.Timer(),
		middleware.CorrelationIDMiddleware(),
	}

	return middleware.Chain(endpoint, middlewares...)
}
