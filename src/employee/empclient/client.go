package main

import (
	"context"
	pb "gRPC/src/employee/emp"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func createEmployee(client pb.EmployeeClient, emp *pb.EmployeeRequest) {
	//CreateEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	res, err := client.CreateEmployee(context.Background(), emp)
	if err != nil {
		log.Fatalf("Could not create the employee %v ", err)
	}
	if res.Success {
		log.Println("A new Employee has been created...", res.EmpId)
	}
}

func getEmployee(client pb.EmployeeClient, singleEmpReq *pb.SingleEmployeeRequest) {
	res, err := client.GetEmployee(context.Background(), singleEmpReq)
	if err != nil {
		log.Println(err)
	}
	if res != nil {
		log.Println("Employee Object has been fetched...", res)
	}
}

func deleteEmployee(client pb.EmployeeClient, deleteEmp *pb.SingleEmployeeRequest) {
	log.Println("Delete the employee : ", deleteEmp.EmpId)
	res, err := client.DeleteEmployee(context.Background(), deleteEmp)
	if err != nil {
		log.Println(err)
	}
	if res != nil && res.Success {
		log.Println("A new Employee has been delete......", res.EmpId)
	}
}

func getAllEmployees(client pb.EmployeeClient, filter *pb.EmployeeFilter) {
	stream, err := client.GetAllEmployees(context.Background(), filter)
	if err != nil {
		log.Fatalf("Could not get customers %v", err)
	}
	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", emp)

	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not albe to get connection %v", err)
	}
	defer conn.Close()
	client := pb.NewEmployeeClient(conn)
	emp := &pb.EmployeeRequest{
		EmpID: 1001,
		Name:  "John D",
		Email: "johnd@gmail.com",
		Addresses: []*pb.EmployeeRequest_Address{
			&pb.EmployeeRequest_Address{
				Street: "1010 Main Street",
				City:   "Reston",
				State:  "VA",
				Zip:    "20151",
			},
			&pb.EmployeeRequest_Address{
				Street: "Elm Tree Drive",
				City:   "Aldie",
				State:  "VA",
				Zip:    "20181",
			},
		},
	}
	createEmployee(client, emp)

	emp = &pb.EmployeeRequest{
		EmpID: 1002,
		Name:  "John E",
		Email: "johnd@gmail.com",
		Addresses: []*pb.EmployeeRequest_Address{
			&pb.EmployeeRequest_Address{
				Street: "1010 Main Street",
				City:   "Reston",
				State:  "VA",
				Zip:    "20151",
			},
			&pb.EmployeeRequest_Address{
				Street: "Elm Tree Drive",
				City:   "Aldie",
				State:  "VA",
				Zip:    "20181",
			},
		},
	}
	createEmployee(client, emp)

	emp = &pb.EmployeeRequest{
		EmpID: 1003,
		Name:  "John F",
		Email: "johnd@gmail.com",
		Addresses: []*pb.EmployeeRequest_Address{
			&pb.EmployeeRequest_Address{
				Street: "1010 Main Street",
				City:   "Reston",
				State:  "VA",
				Zip:    "20151",
			},
			&pb.EmployeeRequest_Address{
				Street: "Elm Tree Drive",
				City:   "Aldie",
				State:  "VA",
				Zip:    "20181",
			},
		},
	}
	createEmployee(client, emp)

	//singleEmpReq := &pb.SingleEmployeeRequest{EmpId: 1001}
	//getEmployee(client, singleEmpReq)

	//singleEmpReq = &pb.SingleEmployeeRequest{EmpId: 1001}
	//deleteEmployee(client, singleEmpReq)
	//deleteEmployee(client, singleEmpReq)

	log.Println(strings.Repeat("*", 25))
	filter := &pb.EmployeeFilter{Keyword: ""}
	getAllEmployees(client, filter)
	log.Println(strings.Repeat("*", 25))
}
