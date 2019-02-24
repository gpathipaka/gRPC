package main

import (
	"context"
	"errors"
	pb "gRPC/src/employee/emp"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type server struct {
	employees map[int32]*pb.EmployeeRequest
}

//Create a new Employee Record.
func (s *server) CreateEmployee(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	if s.employees == nil {
		s.employees = make(map[int32]*pb.EmployeeRequest)
	}
	s.employees[req.EmpID] = req
	log.Println("New Employee record has been added to DB and EMP ID: ", req.EmpID)
	res := &pb.EmployeeResponse{EmpId: req.EmpID, Success: true}
	return res, nil
}

// Get the emplyee using uniq employee Id.
func (s *server) GetEmployee(ctx context.Context, req *pb.SingleEmployeeRequest) (*pb.EmployeeRequest, error) {
	if s.employees != nil {
		emp, ok := s.employees[req.EmpId]
		if ok {
			return emp, nil
		}
	}
	return nil, errors.New("GetEmployee failed...Employee is not found")
}

func (s *server) UpdateEmployee(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	if s.employees != nil {
		_, ok := s.employees[req.EmpID]
		if ok {
			s.employees[req.EmpID] = req
			return &pb.EmployeeResponse{EmpId: req.EmpID, Success: true}, nil
		}
	}
	return nil, errors.New("Udpated fialed...Employee does not exist")
}

// Delete Employee Record
func (s *server) DeleteEmployee(ctx context.Context, req *pb.SingleEmployeeRequest) (*pb.EmployeeResponse, error) {
	if s.employees != nil {
		_, ok := s.employees[req.EmpId]
		if ok {
			delete(s.employees, req.EmpId)
			return &pb.EmployeeResponse{
				EmpId:   req.EmpId,
				Success: true}, nil
		}
	}
	return nil, errors.New("Delete fialed...Employee does not exist")
}

func (s *server) GetAllEmployees(filter *pb.EmployeeFilter, stream pb.Employee_GetAllEmployeesServer) error {
	for _, emp := range s.employees {
		if filter.Keyword != "" {
			if !strings.Contains(emp.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(emp); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Println("Server started....")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := &server{}
	pb.RegisterEmployeeServer(s, srv)
	s.Serve(listener)
	log.Println("Server about to go down......")
}
