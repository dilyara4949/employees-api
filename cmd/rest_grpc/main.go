package main

import (
	"fmt"
	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/grpc/server"
	"github.com/dilyara4949/employees-api/internal/repository/employee"
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/route"
	pb "github.com/dilyara4949/employees-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	positionsStorage := position.NewPositionsStorage()
	employeeStorage := employee.NewEmployeesStorage()

	positionRepo := position.NewPositionsRepository(positionsStorage)
	employeeRepo := employee.NewEmployeesRepository(employeeStorage, positionRepo)

	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	go func() {
		positionServer := server.NewPositionServer(positionRepo)
		employeeServer := server.NewEmployeeServer(employeeRepo)

		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Address, config.GrpcPort))
		if err != nil {
			log.Fatalf("Could not listen on port: %v", err)
		}

		s := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				server.CorrelationIDInterceptor(),
				server.LoggingInterceptor,
			),
			grpc.StreamInterceptor(server.StreamLoggingInterceptor),
		)
		pb.RegisterPositionServiceServer(s, positionServer)
		pb.RegisterEmployeeServiceServer(s, employeeServer)

		reflection.Register(s)

		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

		log.Printf("Hosting server on: %s", listen.Addr().String())
	}()

	positionController := controller.NewPositionsController(positionRepo)
	employeeController := controller.NewEmployeesController(employeeRepo)

	mux := http.NewServeMux()

	route.SetUpRouter(employeeController, positionController, config, mux)

	log.Printf("Starting server on :%s", config.RestPort)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", config.Address, config.RestPort), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
