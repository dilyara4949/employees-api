package route

import (
	"net/http"

	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/tokenutil"
)

func SetUpRouter(employeeController *controller.EmployeeController, positionController *controller.PositionController, mux *http.ServeMux) {
	mux.HandleFunc("GET /positions/{id}", tokenutil.JwtMiddleware(positionController.GetPosition))
	mux.HandleFunc("POST /positions", tokenutil.JwtMiddleware(positionController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", tokenutil.JwtMiddleware(positionController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", tokenutil.JwtMiddleware(positionController.UpdatePosition))

	mux.HandleFunc("GET /employees/{id}", tokenutil.JwtMiddleware(employeeController.GetEmployee))
	mux.HandleFunc("POST /employees", tokenutil.JwtMiddleware(employeeController.CreateEmployee))
	mux.HandleFunc("DELETE /employees/{id}", tokenutil.JwtMiddleware(employeeController.DeleteEmployee))
	mux.HandleFunc("PUT /employees/{id}", tokenutil.JwtMiddleware(employeeController.UpdateEmployee))
}
