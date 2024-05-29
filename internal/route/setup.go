package route

import (
	"github.com/dilyara4949/employees-api/internal/controller"
	"log"
	"net/http"
)

func SetUpRouter(employeeController *controller.EmployeeController, positionController *controller.PositionController) {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /positions/{id}", controllerErrorHandler(positionController.GetPosition))
	mux.HandleFunc("POST /positions", controllerErrorHandler(positionController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", controllerErrorHandler(positionController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", controllerErrorHandler(positionController.UpdatePosition))

	mux.HandleFunc("GET /employees/{id}", controllerErrorHandler(employeeController.GetEmployee))
	mux.HandleFunc("POST /employees", controllerErrorHandler(employeeController.CreateEmployee))
	mux.HandleFunc("DELETE /employees/{id}", controllerErrorHandler(employeeController.DeleteEmployee))
	mux.HandleFunc("PUT /employees/{id}", controllerErrorHandler(employeeController.UpdateEmployee))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}

func controllerErrorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			log.Printf("HTTP error at %v: %v", r.URL, err)
			if httpErr, ok := err.(*controller.HTTPError); ok {
				http.Error(w, httpErr.Detail, httpErr.Status)
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
	}
}
