package main

import (
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/repository/storage"
	pb "github.com/dilyara4949/employees-api/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	positionsStorage := storage.NewPositionsStorage()
	employeeStorage := storage.NewEmployeesStorage()

	employeeRepo := employee.NewEmployeesRepository(employeeStorage, positionsStorage)
	positionRepo := position.NewPositionsRepository(positionsStorage)

	positionServer := NewPositionServer(positionRepo)
	employeeServer := NewEmployeeServer(employeeRepo)

	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPositionServiceServer(s, positionServer)
	pb.RegisterEmployeeServiceServer(s, employeeServer)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}
