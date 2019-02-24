package main

import (
	"context"
	pb "gRPC/src/customer/customerProto"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":40044"
)

type server struct {
	savedCustomers []*pb.CustomerReqeust
}

func (s *server) CreateCustomer(ctx context.Context, input *pb.CustomerReqeust) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, input)
	log.Println("New Customer Created ... ID ", input.Id)
	return &pb.CustomerResponse{Id: input.Id, Success: true}, nil
}

// GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Println("Server Started....")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &server{})
	s.Serve(listener)
	log.Println("Server is about to go down.....")
}
