package route

import (
	"log"
	"net/http"

	"github.com/dilyara4949/employees-api/api/controller"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
)

func NewRouter(employeeStorage *employee.Storage, positionStorage *position.Storage) {

	positionRepo := position.NewPositionRepository(positionStorage)
	positionController := controller.NewPositionController(positionRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /positions/{id}", controllerErrorHandler(positionController.GetPosition))
	mux.HandleFunc("POST /positions", controllerErrorHandler(positionController.CreatePosition))
	mux.HandleFunc("DELETE /positions/{id}", controllerErrorHandler(positionController.DeletePosition))
	mux.HandleFunc("PUT /positions/{id}", controllerErrorHandler(positionController.UpdatePosition))

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
			log.Printf("HTTP error: %v", err)
			if httpErr, ok := err.(*controller.HTTPError); ok {
				http.Error(w, httpErr.Detail, httpErr.Status)
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
	}
}
