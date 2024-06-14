package server

import (
	"context"

	"github.com/dilyara4949/employees-api/internal/domain"
	pb "github.com/dilyara4949/employees-api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmployeeServer struct {
	Repo domain.EmployeesRepository
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeServer(repo domain.EmployeesRepository) *EmployeeServer {
	return &EmployeeServer{
		Repo: repo,
	}
}

func (s *EmployeeServer) GetAll(ctx context.Context, req *pb.GetAllEmployeesRequest) (*pb.EmployeesList, error) {
	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page <= 0 || pageSize <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "page and page size cannot be less than 1")
	}
	employees, err := s.Repo.GetAll(ctx, page, pageSize)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	employeeProtos := make([]*pb.Employee, len(employees))
	for i, emp := range employees {
		employeeProtos[i] = employeeToProto(&emp)
	}
	return &pb.EmployeesList{Employee: employeeProtos, Page: page, PageSize: pageSize}, nil
}

func (s *EmployeeServer) Get(ctx context.Context, id *pb.Id) (*pb.Employee, error) {
	if id == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil id in get employee")
	}

	employee, err := s.Repo.Get(ctx, id.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return employeeToProto(employee), nil
}

func (s *EmployeeServer) Create(ctx context.Context, emp *pb.Employee) (*pb.Employee, error) {
	if emp == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil employee in create employee")
	}

	employee := protoToEmployee(emp)

	err := s.Repo.Create(ctx, employee)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return employeeToProto(employee), nil
}

func (s *EmployeeServer) Update(ctx context.Context, emp *pb.Employee) (*pb.Employee, error) {
	if emp == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil employee in update employee")
	}

	employee := protoToEmployee(emp)

	err := s.Repo.Update(ctx, *employee)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return emp, nil
}

func (s *EmployeeServer) Delete(ctx context.Context, id *pb.Id) (*pb.Status, error) {
	if id == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil id in delete employees")
	}

	err := s.Repo.Delete(ctx, id.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.Status{Status: 0}, nil
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
