package main

import (
	"context"
	"log"
	"time"

	pb "grpc-nats/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to Service 1
	conn1, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Service 1: %v", err)
	}
	defer conn1.Close()
	client1 := pb.NewServiceOneClient(conn1)

	_, err = client1.PublishMessage(context.Background(), &pb.Request{Message: "Hello from Service 1"})
	if err != nil {
		log.Fatalf("Error calling PublishMessage: %v", err)
	}
	log.Println("Message published to NATS via Service 1")

	time.Sleep(2 * time.Second)

	// Connect to Service 2
	conn2, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Service 2: %v", err)
	}
	defer conn2.Close()
	client2 := pb.NewServiceTwoClient(conn2)

	res, err := client2.GetResponse(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error calling GetResponse: %v", err)
	}
	log.Printf("Message from Service 2: %s", res.Reply)
}
