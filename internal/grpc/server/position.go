package main

import (
	"context"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/protobuf/proto"
)

type PositionServer struct {
	Repo domain.PositionsRepository
	proto.UnimplementedPositionServiceServer
}

func (s *PositionServer) GetAll(empty *proto.Empty, stream proto.PositionService_GetAllServer) error {
	positions := s.Repo.GetAll()
	for _, pos := range positions {
		if err := stream.Send(positionToProto(&pos)); err != nil {
			return err
		}
	}
	return nil
}

func NewPositionServer(repo domain.PositionsRepository) *PositionServer {
	return &PositionServer{
		Repo: repo,
	}
}

func (s *PositionServer) Get(_ context.Context, id *proto.Id) (*proto.Position, error) {
	if id == nil {
		return nil, errors.New("got nil id in get position")
	}

	position, err := s.Repo.Get(id.Value)
	if err != nil {
		return nil, err
	}
	return positionToProto(position), nil
}

func (s *PositionServer) Create(_ context.Context, pos *proto.Position) (*proto.Position, error) {
	if pos == nil {
		return nil, errors.New("got nil position in create position")
	}

	position := protoToPosition(pos)
	err := s.Repo.Create(position)
	if err != nil {
		return nil, err
	}
	return positionToProto(position), nil
}

func (s *PositionServer) Update(_ context.Context, pos *proto.Position) (*proto.Position, error) {
	if pos == nil {
		return nil, errors.New("got nil position in update position")
	}

	position := protoToPosition(pos)
	err := s.Repo.Update(*position)
	if err != nil {
		return nil, err
	}
	return pos, nil
}

func (s *PositionServer) Delete(_ context.Context, id *proto.Id) (*proto.Status, error) {
	if id == nil {
		return nil, errors.New("got nil id in delete positions")
	}

	err := s.Repo.Delete(id.Value)
	if err != nil {
		return nil, err
	}
	return &proto.Status{Status: 204}, nil
}

func positionToProto(p *domain.Position) *proto.Position {
	if p == nil {
		return nil
	}
	return &proto.Position{
		Id: p.ID, Name: p.Name, Salary: int32(p.Salary),
	}
}

func protoToPosition(p *proto.Position) *domain.Position {
	if p == nil {
		return nil
	}
	return &domain.Position{
		ID: p.Id, Name: p.Name, Salary: int(p.Salary),
	}
}
