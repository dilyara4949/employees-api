package server

import (
	"context"
	"errors"
	"github.com/dilyara4949/employees-api/internal/domain"
	pb "github.com/dilyara4949/employees-api/proto"
)

type EmployeeServer struct {
	Repo domain.EmployeesRepository
	pb.UnimplementedEmployeeServiceServer
}

func (s *EmployeeServer) GetAll(empty *pb.Empty, stream pb.EmployeeService_GetAllServer) error {
	employees := s.Repo.GetAll()
	for _, emp := range employees {
		if err := stream.Send(employeeToProto(&emp)); err != nil {
			return err
		}
	}
	return nil
}

func NewEmployeeServer(repo domain.EmployeesRepository) *EmployeeServer {
	return &EmployeeServer{
		Repo: repo,
	}
}

func (s *EmployeeServer) Get(ctx context.Context, id *pb.Id) (*pb.Employee, error) {
	if id == nil {
		return nil, errors.New("got nil id in get employee")
	}

	employee, err := s.Repo.Get(id.Value)
	if err != nil {
		return nil, err
	}
	return employeeToProto(employee), nil
}

func (s *EmployeeServer) Create(_ context.Context, emp *pb.Employee) (*pb.Employee, error) {
	if emp == nil {
		return nil, errors.New("got nil employee in create employee")
	}

	employee := protoToEmployee(emp)
	err := s.Repo.Create(employee)
	if err != nil {
		return nil, err
	}
	return employeeToProto(employee), nil
}

func (s *EmployeeServer) Update(_ context.Context, emp *pb.Employee) (*pb.Employee, error) {
	if emp == nil {
		return nil, errors.New("got nil employee in update employee")
	}

	employee := protoToEmployee(emp)
	err := s.Repo.Update(*employee)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (s *EmployeeServer) Delete(_ context.Context, id *pb.Id) (*pb.Status, error) {
	if id == nil {
		return nil, errors.New("got nil id in delete employees")
	}

	err := s.Repo.Delete(id.Value)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: 204}, nil
}

func employeeToProto(e *domain.Employee) *pb.Employee {
	if e == nil {
		return nil
	}
	return &pb.Employee{
		Id: e.ID, Firstname: e.FirstName, Lastname: e.LastName, PositionId: e.PositionID,
	}
}

func protoToEmployee(e *pb.Employee) *domain.Employee {
	if e == nil {
		return nil
	}
	return &domain.Employee{
		ID: e.Id, FirstName: e.Firstname, LastName: e.Lastname, PositionID: e.PositionId,
	}
}
