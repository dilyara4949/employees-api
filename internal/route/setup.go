package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/tokenutil"
)

func SetUpRouter(employeesController *controller.EmployeesController, positionsController *controller.PositionsController, mux *http.ServeMux) {
	mux.HandleFunc("GET /positions/{id}", tokenutil.JwtMiddleware(positionsController.GetPosition))
	mux.HandleFunc("POST /positions", tokenutil.JwtMiddleware(positionsController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", tokenutil.JwtMiddleware(positionsController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", tokenutil.JwtMiddleware(positionsController.UpdatePosition))
	mux.HandleFunc("GET /positions", tokenutil.JwtMiddleware(positionsController.GetAllPositions))

	mux.HandleFunc("GET /employees/{id}", tokenutil.JwtMiddleware(employeesController.GetEmployee))
	mux.HandleFunc("POST /employees", tokenutil.JwtMiddleware(employeesController.CreateEmployee))
	mux.HandleFunc("DELETE /employees/{id}", tokenutil.JwtMiddleware(employeesController.DeleteEmployee))
	mux.HandleFunc("PUT /employees/{id}", tokenutil.JwtMiddleware(employeesController.UpdateEmployee))
	mux.HandleFunc("GET /employees", tokenutil.JwtMiddleware(employeesController.GetAllEmployees))
}
