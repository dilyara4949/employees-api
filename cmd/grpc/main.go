package main

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/grpc/server"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/repository/storage"
	pb "github.com/dilyara4949/employees-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	positionsStorage := storage.NewPositionsStorage()
	employeeStorage := storage.NewEmployeesStorage()

	employeeRepo := employee.NewEmployeesRepository(employeeStorage, positionsStorage)
	positionRepo := position.NewPositionsRepository(positionsStorage)

	positionServer := server.NewPositionServer(positionRepo)
	employeeServer := server.NewEmployeeServer(employeeRepo)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error while reading configs: %s", err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Address, cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			server.CorrelationIDInterceptor(),
			server.LoggingInterceptor,
		),
	)
	pb.RegisterPositionServiceServer(s, positionServer)
	pb.RegisterEmployeeServiceServer(s, employeeServer)

	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}
