package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal"
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/middleware"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, env *internal.Env, mux *http.ServeMux) {

	auth := middleware.NewJWTAuth(env.JWTTokenSecret)

	mux.HandleFunc("GET /positions/{id}", middleware.Adapt(auth.Auth(positionsController.GetPosition), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("POST /positions", middleware.Adapt(auth.Auth(positionsController.CreatePosition), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("DELETE /positions/{id}", middleware.Adapt(auth.Auth(positionsController.DeletePosition), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("PUT /positions/{id}", middleware.Adapt(auth.Auth(positionsController.UpdatePosition), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("GET /positions", middleware.Adapt(auth.Auth(positionsController.GetAllPositions), middleware.Logger(), middleware.СorrelationIDMiddleware()))

	mux.HandleFunc("GET /employees/{id}", middleware.Adapt(auth.Auth(employeesController.GetEmployee), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("POST /employees", middleware.Adapt(auth.Auth(employeesController.CreateEmployee), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("DELETE /employees/{id}", middleware.Adapt(auth.Auth(employeesController.DeleteEmployee), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("PUT /employees/{id}", middleware.Adapt(auth.Auth(employeesController.UpdateEmployee), middleware.Logger(), middleware.СorrelationIDMiddleware()))
	mux.HandleFunc("GET /employees", middleware.Adapt(auth.Auth(employeesController.GetAllEmployees), middleware.Logger(), middleware.СorrelationIDMiddleware()))
}
