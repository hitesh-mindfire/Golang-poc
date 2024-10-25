package main

import (
	"context"
	"log"
	"net"

	pb "grpc-microservices/generated"

	"google.golang.org/grpc"
)

// Server represents the gRPC server
type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// Mock user data
	user := &pb.UserResponse{
		UserId: req.UserId,
		Name:   "John Doe",
	}
	return user, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &Server{})

	log.Println("User Service running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
