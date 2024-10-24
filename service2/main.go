
package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "grpc-nats/generated"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

const natsURL = "nats://localhost:4222"

type server struct {
	pb.UnimplementedServiceTwoServer
	lastMessage string
	mu          sync.Mutex
}

func (s *server) GetResponse(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &pb.Response{Reply: s.lastMessage}, nil
}

func main() {
	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	srv := &server{}

	nc.Subscribe("updates", func(m *nats.Msg) {
		log.Printf("Received message from NATS: %s", string(m.Data))
		srv.mu.Lock()
		srv.lastMessage = string(m.Data)
		srv.mu.Unlock()
	})

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceTwoServer(grpcServer, srv)

	log.Println("Service Two running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
