package route

import (
	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/redis/go-redis/v9"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, config conf.Config, mux *http.ServeMux, cache *redis.Client) {

	mux.HandleFunc("GET /positions/{id}", logCorrelationIDTimer(positionsController.GetPosition, config, cache))
	mux.HandleFunc("POST /positions", logCorrelationIDTimer(positionsController.CreatePosition, config, cache))
	mux.HandleFunc("DELETE /positions/{id}", logCorrelationIDTimer(positionsController.DeletePosition, config, cache))
	mux.HandleFunc("PUT /positions/{id}", logCorrelationIDTimer(positionsController.UpdatePosition, config, cache))
	mux.HandleFunc("GET /positions", logCorrelationIDTimer(positionsController.GetAllPositions, config, cache))

	mux.HandleFunc("GET /employees/{id}", logCorrelationIDTimer(employeesController.GetEmployee, config, cache))
	mux.HandleFunc("POST /employees", logCorrelationIDTimer(employeesController.CreateEmployee, config, cache))
	mux.HandleFunc("DELETE /employees/{id}", logCorrelationIDTimer(employeesController.DeleteEmployee, config, cache))
	mux.HandleFunc("PUT /employees/{id}", logCorrelationIDTimer(employeesController.UpdateEmployee, config, cache))
	mux.HandleFunc("GET /employees", logCorrelationIDTimer(employeesController.GetAllEmployees, config, cache))
}

func logCorrelationIDTimer(endpoint http.HandlerFunc, config conf.Config, cache *redis.Client) http.HandlerFunc {
	auth := middleware.NewJWTAuth(config.JWTTokenSecret)
	middlewares := []middleware.Middleware{
		middleware.Cache(cache, config.RedisConfig.Ttl),
		auth.Auth(),
		middleware.Logger(),
		middleware.Timer(),
		middleware.CorrelationIDMiddleware(),
	}

	return middleware.Chain(endpoint, middlewares...)
}
