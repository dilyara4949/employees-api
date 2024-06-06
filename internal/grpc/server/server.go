package main

import (
	"github.com/dilyara4949/employees-api/internal/repository/position"
	"github.com/dilyara4949/employees-api/internal/repository/storage"
	pb "github.com/dilyara4949/employees-api/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	positionsStorage := storage.NewPositionsStorage()

	positionRepo := position.NewPositionsRepository(positionsStorage)

	svr := NewPositionServer(positionRepo)

	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPositionServiceServer(s, svr)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}
