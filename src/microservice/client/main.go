package main

import (
	"context"
	"encoding/json"
	pb "gRPC/src/microservice/proto/consignment"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	addres          = "localhost:8080"
	defaultFilename = "consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	var cons *pb.Consignment
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &cons)
	return cons, err

}

func createConsignment(client pb.ShippingServiceClient, cons *pb.Consignment) {
	res, err := client.CreateConsignment(context.Background(), cons)
	if err != nil {
		log.Printf("Could not Gree.. %v", err)
	}
	log.Println("Consignment has been created....", res.Created)
}

func getAllConsignments(client pb.ShippingServiceClient) {
	res, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Printf("Could not get the consignments.. %v", err)
	}
	log.Println(res.Consignments)
}
func main() {
	log.Println("Client Started...")
	conn, err := grpc.Dial(addres, grpc.WithInsecure())
	if err != nil {
		log.Printf("Did not connect to server...%v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	cons, err := parseFile(file)
	if err != nil {
		log.Printf("Could not read the file %v", err)
	}
	createConsignment(client, cons)

	getAllConsignments(client)
	log.Println("Client about to go down...")
}
