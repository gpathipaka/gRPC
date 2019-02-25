package main

import (
	"context"
	pb "gRPC/src/helloworld/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func sayHello(client pb.GreeterClient, req *pb.Request) {
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Got the message from Server ", res.Message)
}

func streamServer(client pb.GreeterClient, req *pb.Request) {
	stream, err := client.SayHelloStream(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.Message)
	}
}

func main() {
	log.Println("Client Started....")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Not albe to get connection %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	req := &pb.Request{Name: "Gangadhar"}
	sayHello(client, req)

	streamServer(client, req)
}
