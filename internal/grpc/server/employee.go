package server

import (
	"context"
	"log"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/dilyara4949/employees-api/internal/repository/redis"
	pb "github.com/dilyara4949/employees-api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmployeeServer struct {
	Repo  domain.EmployeesRepository
	cache redis.EmployeeCache
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeServer(repo domain.EmployeesRepository, cache redis.EmployeeCache) *EmployeeServer {
	return &EmployeeServer{
		Repo:  repo,
		cache: cache,
	}
}

const (
	pageDefault     = 1
	pageSizeDefault = 50
)

func (s *EmployeeServer) GetAll(ctx context.Context, req *pb.GetAllEmployeesRequest) (*pb.EmployeesList, error) {
	page := req.GetPage()
	pageSize := req.GetPageSize()

	if page <= 0 || pageSize <= 0 {
		page = pageDefault
		pageSize = pageSizeDefault
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

	employee, err := s.cache.Get(ctx, id.GetValue())
	if err != nil || employee == nil {
		employee, err = s.Repo.Get(ctx, id.GetValue())
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		err = s.cache.Set(ctx, id.GetValue(), employee)
		if err != nil {
			log.Printf("error at caching employee: %v", err)
		}
	}
	return employeeToProto(employee), nil
}

func (s *EmployeeServer) Create(ctx context.Context, emp *pb.Employee) (*pb.Employee, error) {
	if emp == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil employee in create employee")
	}

	employee := protoToEmployee(emp)

	employee, err := s.Repo.Create(ctx, *employee)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.cache.Set(ctx, employee.ID, employee)
	if err != nil {
		log.Printf("error at caching employee: %v", err)
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

	err = s.cache.Set(ctx, employee.ID, employee)
	if err != nil {
		log.Printf("error at updating employee cache: %v", err)
	}
	return emp, nil
}

func (s *EmployeeServer) Delete(ctx context.Context, id *pb.Id) (*pb.Status, error) {
	if id == nil {
		return nil, status.Errorf(codes.InvalidArgument, "got nil id in delete employees")
	}

	err := s.Repo.Delete(ctx, id.GetValue())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = s.cache.Delete(ctx, id.GetValue())
	if err != nil {
		log.Printf("error at deleting employee from cache: %v", err)
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
