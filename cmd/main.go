package main

import (
	"fmt"
	mongoDB "github.com/dilyara4949/employees-api/internal/database/mongo"
	"github.com/dilyara4949/employees-api/internal/database/postgres"
	"github.com/dilyara4949/employees-api/internal/database/redis"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/grpc/server"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/employee"
	redis2 "github.com/dilyara4949/employees-api/internal/repository/redis"

	mongoemployee "github.com/dilyara4949/employees-api/internal/repository/mongo/employee"
	mongoposition "github.com/dilyara4949/employees-api/internal/repository/mongo/position"
	"github.com/dilyara4949/employees-api/internal/repository/postgres/position"
	pb "github.com/dilyara4949/employees-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"

	conf "github.com/dilyara4949/employees-api/internal/config"
	"github.com/dilyara4949/employees-api/internal/controller"
	"github.com/dilyara4949/employees-api/internal/route"
)

func main() {

	config, err := conf.NewConfig()
	if err != nil {
		log.Fatalf("Error while getting config: %s", err)
	}

	var positionRepo domain.PositionsRepository
	var employeeRepo domain.EmployeesRepository

	switch config.DatabaseType {
	case "postgres":
		db, err := postgres.ConnectPostgres(config.PostgresConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}
		defer db.Close()

		positionRepo = position.NewPositionsRepository(db)
		employeeRepo = employee.NewEmployeesRepository(db, positionRepo)
	case "mongo":
		db, err := mongoDB.ConnectMongo(config.MongoConfig)
		if err != nil {
			log.Fatalf("Connection to database failed: %s", err)
		}

		positionRepo = mongoposition.NewPositionsRepository(db, config.MongoConfig.Collections.Positions, config.MongoConfig.Collections.Employees)
		employeeRepo = mongoemployee.NewEmployeesRepository(db, config.MongoConfig.Collections.Employees, config.MongoConfig.Collections.Positions)

	default:
		log.Fatalf("%s is unknown database", config.DatabaseType)
	}

	log.Println("Successfully connected to database")

	go func() {
		positionServer := server.NewPositionServer(positionRepo)
		employeeServer := server.NewEmployeeServer(employeeRepo)

		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Address, config.GrpcPort))
		if err != nil {
			log.Fatalf("Could not listen on port: %v", err)
		}

		svr := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				server.CorrelationIDInterceptor(),
				server.LoggingInterceptor,
			),
		)
		pb.RegisterPositionServiceServer(svr, positionServer)
		pb.RegisterEmployeeServiceServer(svr, employeeServer)

		reflection.Register(svr)

		if err := svr.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

		log.Printf("Hosting server on: %s", listen.Addr().String())
	}()

	rdb := redis.ConnectRedis(config.RedisConfig)

	positionCache := redis2.NewPositionCache(rdb, config.RedisConfig.Timeout)

	positionController := controller.NewPositionsController(positionRepo, positionCache)
	employeeController := controller.NewEmployeesController(employeeRepo)

	mux := http.NewServeMux()

	route.SetUpRouter(employeeController, positionController, config, mux)

	log.Printf("Starting server on :%s", config.RestPort)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", config.Address, config.RestPort), mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
