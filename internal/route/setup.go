package route

import (
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/tokenutil"
	"net/http"
)

func SetUpRouter(employeeController *controller.EmployeeController, positionController *controller.PositionController, mux *http.ServeMux) {
	//mux.Handle("/positions/{id}", tokenutil.JwtMiddleware(positionController.GetPosition))
	mux.HandleFunc("GET /positions/{id}", tokenutil.JwtMiddleware(positionController.GetPosition))
	mux.HandleFunc("POST /positions", positionController.CreatePosition)
	mux.HandleFunc("DELETE /positions/{id}", positionController.DeletePosition)
	mux.HandleFunc("PUT /positions/{id}", positionController.UpdatePosition)

	mux.HandleFunc("GET /employees/{id}", employeeController.GetEmployee)
	mux.HandleFunc("POST /employees", employeeController.CreateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", employeeController.DeleteEmployee)
	mux.HandleFunc("PUT /employees/{id}", employeeController.UpdateEmployee)
}
