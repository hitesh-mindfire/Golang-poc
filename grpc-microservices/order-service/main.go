package main

import (
	"context"
	"log"
	"time"

	pb "grpc-microservices/generated"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to gRPC User Service
	conn, err := grpc.NewClient("user-service:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to User Service: %v", err)
	}
	defer conn.Close()
	userServiceClient := pb.NewUserServiceClient(conn)

	// Connect to NATS
	nc, err := nats.Connect("nats-server:4222")

	if err != nil {
		log.Fatalf("Could not connect to NATS: %v", err)
	}
	defer nc.Close()
	createOrder(userServiceClient, nc, "123")
	select {}
}

func createOrder(userServiceClient pb.UserServiceClient, nc *nats.Conn, userID string) {
	// Fetch user details using gRPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	userResponse, err := userServiceClient.GetUser(ctx, &pb.UserRequest{UserId: userID})
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}
	log.Printf("Creating order for user: %s\n", userResponse.Name)

	// Publish event to NATS
	orderID := "order-1234"
	message := "Order created for " + userResponse.Name + " (Order ID: " + orderID + ")"
	err = nc.Publish("order.created", []byte(message))
	if err != nil {
		log.Fatalf("Error publishing to NATS: %v", err)
	}
	log.Println("Order event published to NATS.")
}
