package route

import (
	"github.com/dilyara4949/employees-api/internal/middleware"
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/tokenutil"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, mux *http.ServeMux) {

	mux.HandleFunc("GET /positions/{id}", middleware.Adapt(tokenutil.JwtMiddleware(positionsController.GetPosition), middleware.Logger()))
	mux.HandleFunc("POST /positions", middleware.Adapt(tokenutil.JwtMiddleware(positionsController.CreatePosition), middleware.Logger()))
	mux.HandleFunc("DELETE /positions/{id}", middleware.Adapt(tokenutil.JwtMiddleware(positionsController.DeletePosition), middleware.Logger()))
	mux.HandleFunc("PUT /positions/{id}", middleware.Adapt(tokenutil.JwtMiddleware(positionsController.UpdatePosition), middleware.Logger()))
	mux.HandleFunc("GET /positions", middleware.Adapt(tokenutil.JwtMiddleware(positionsController.GetAllPositions), middleware.Logger()))

	mux.HandleFunc("GET /employees/{id}", middleware.Adapt(tokenutil.JwtMiddleware(employeesController.GetEmployee), middleware.Logger()))
	mux.HandleFunc("POST /employees", middleware.Adapt(tokenutil.JwtMiddleware(employeesController.CreateEmployee), middleware.Logger()))
	mux.HandleFunc("DELETE /employees/{id}", middleware.Adapt(tokenutil.JwtMiddleware(employeesController.DeleteEmployee), middleware.Logger()))
	mux.HandleFunc("PUT /employees/{id}", middleware.Adapt(tokenutil.JwtMiddleware(employeesController.UpdateEmployee), middleware.Logger()))
	mux.HandleFunc("GET /employees", middleware.Adapt(tokenutil.JwtMiddleware(employeesController.GetAllEmployees), middleware.Logger()))
}
