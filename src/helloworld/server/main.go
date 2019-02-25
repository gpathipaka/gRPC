package main

import (
	"context"
	pb "gRPC/src/helloworld/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

//SayHello is
func (s *server) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "Hello " + req.Name}, nil
}

//SayHelloStream is streaming service.
func (s *server) SayHelloStream(req *pb.Request, stream pb.Greeter_SayHelloStreamServer) error {
	for {
		err := stream.Send(&pb.Response{Message: "hello " + req.Name})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

type server struct {
}

func main() {
	log.Println("Server started...")
	listerner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(listerner)
	log.Println("Server going down...")
}
