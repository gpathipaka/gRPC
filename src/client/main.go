package main

import (
	"context"
	pb "gRPC/src/customer"
	"io"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "localhost:40044"
)

func createCustomer(client pb.CustomerClient, customer *pb.CustomerReqeust) {
	res, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("Could not create customer ", err)
	}
	if res.Success {
		log.Printf("A new customer has been creted with ID : %d", res.Id)
	}
}

func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Could not get customers %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}
func main() {
	log.Println("Client Main() Start...")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerReqeust{
		Id:    5001,
		Name:  "Reyan Pathipaka",
		Email: "Reyan@abc.com",
		Phone: "444-555-8524",
		Addresses: []*pb.CustomerReqeust_Address{
			&pb.CustomerReqeust_Address{
				Street:            "101 Main Stree",
				City:              "Washington DC",
				State:             "DC",
				Zip:               "10024",
				IsShippingAddress: true,
			},
			&pb.CustomerReqeust_Address{
				Street:            "112 Laurel Tree",
				City:              "Reston",
				State:             "VA",
				Zip:               "20170",
				IsShippingAddress: false,
			},
		},
	}

	createCustomer(client, customer)

	customer = &pb.CustomerReqeust{
		Id:    5002,
		Name:  "Ishu Pathipaka",
		Email: "Pathipaka@abc.com",
		Phone: "444-555-8524",
		Addresses: []*pb.CustomerReqeust_Address{
			&pb.CustomerReqeust_Address{
				Street:            "1280 West Stree SW",
				City:              "Washington DC",
				State:             "DC",
				Zip:               "10028",
				IsShippingAddress: true,
			},
		},
	}

	createCustomer(client, customer)

	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)

	log.Println("Client Main() Start...")

}
