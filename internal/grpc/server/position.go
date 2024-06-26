package server

import (
	"context"
	"log"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/redis"
	pb "github.com/dilyara4949/employees-api/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PositionServer struct {
	Repo  domain.PositionsRepository
	cache redis.PositionCache
	pb.UnimplementedPositionServiceServer
}

func (s *PositionServer) GetAll(ctx context.Context, req *pb.GetAllPositionsRequest) (*pb.PositionsList, error) {
	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page <= 0 || pageSize <= 0 {
		page = pageDefault
		pageSize = pageSizeDefault
	}

	positions, err := s.Repo.GetAll(ctx, page, pageSize)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	positionProtos := make([]*pb.Position, len(positions))
	for i, pos := range positions {
		positionProtos[i] = positionToProto(&pos)
	}
	return &pb.PositionsList{Position: positionProtos, Page: page, PageSize: pageSize}, nil
}

func NewPositionServer(repo domain.PositionsRepository, cache redis.PositionCache) *PositionServer {
	return &PositionServer{
		Repo:  repo,
		cache: cache,
	}
}

func (s *PositionServer) Get(ctx context.Context, id *pb.Id) (*pb.Position, error) {
	if id == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil id in get position")
	}

	position, err := s.cache.Get(ctx, id.GetValue())
	if err != nil || position == nil {
		position, err = s.Repo.Get(ctx, id.Value)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error getting position: %v", err)
		}

		err = s.cache.Set(ctx, id.GetValue(), position)
		if err != nil {
			log.Printf("error at caching position: %v", err)
		}
	}

	return positionToProto(position), nil
}

func (s *PositionServer) Create(ctx context.Context, pos *pb.Position) (*pb.Position, error) {
	if pos == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil position in create position")
	}

	position := protoToPosition(pos)
	position.ID = uuid.New().String()

	position, err := s.Repo.Create(ctx, *position)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.cache.Set(ctx, position.ID, position)
	if err != nil {
		log.Printf("error at caching position: %v", err)
	}
	return positionToProto(position), nil
}

func (s *PositionServer) Update(ctx context.Context, pos *pb.Position) (*pb.Position, error) {
	if pos == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil position in update position")
	}

	position := protoToPosition(pos)

	err := s.Repo.Update(ctx, *position)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.cache.Set(ctx, position.ID, position)
	if err != nil {
		log.Printf("error at updating position cache: %v", err)
	}
	return pos, nil
}

func (s *PositionServer) Delete(ctx context.Context, id *pb.Id) (*pb.Status, error) {
	if id == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil id in delete positions")
	}

	err := s.Repo.Delete(ctx, id.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.cache.Delete(ctx, id.Value)
	if err != nil {
		log.Printf("error at deleting position from cache: %v", err)
	}
	return &pb.Status{Status: 0}, nil
}

func positionToProto(p *domain.Position) *pb.Position {
	if p == nil {
		return nil
	}
	return &pb.Position{
		Id: p.ID, Name: p.Name, Salary: int32(p.Salary),
	}
}

func protoToPosition(p *pb.Position) *domain.Position {
	if p == nil {
		return nil
	}
	return &domain.Position{
		ID: p.Id, Name: p.Name, Salary: int(p.Salary),
	}
}
