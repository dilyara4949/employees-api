package route

import (
	"github.com/dilyara4949/employees-api/api/controller"
	"github.com/dilyara4949/employees-api/internal/repository/position"
)

func NewRouter(db) {
	positionRepository := position.NewPositionRepository()
	positionController := controller.NewPositionController()
}
