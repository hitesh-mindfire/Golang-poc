package main

import (
	"context"
	"log"
	"net"

	pb "grpc-nats/generated"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

const natsURL = "nats://localhost:4222"

type server struct {
	pb.UnimplementedServiceOneServer
	nc *nats.Conn
}

func (s *server) PublishMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	msg := req.Message

	err := s.nc.Publish("updates", []byte(msg))
	if err != nil {
		return nil, err
	}
	log.Printf("Published to NATS: %s", msg)

	return &pb.Response{Reply: "Message published!"}, nil
}

func main() {
	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceOneServer(s, &server{nc: nc})

	log.Println("Service One running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
