package main

import (
	"github.com/dilyara4949/employees-api/api/route"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
)

func main() {

	storageP := &position.Storage{
		Storage: make(map[string]position.Position),
	}
	storageE := &employee.Storage{
		Storage: make(map[string]employee.Employee),
	}
	route.NewRouter(storageE, storageP)
}
