# Microservices with gRPC, NATS, Docker, and Kubernetes

This project demonstrates a microservices architecture using **gRPC** for service-to-service communication and **NATS** as an event streaming platform. Each service runs in a Docker container and is managed on Kubernetes.

## Project Overview

The microservices architecture contains the following services:

- **User Service**: Exposes user information via gRPC.
- **Order Service**: Creates orders and publishes events via NATS.
- **Notification Service**: Listens for order events on NATS and handles notifications.

### Key Tools and Technologies

- **gRPC**: Remote procedure call (RPC) framework for efficient communication between services.
- **NATS**: Lightweight messaging system for event streaming between services.
- **Docker**: Containerization platform to package and distribute services.
- **Kubernetes**: Manages and scales the containerized services.

## Service Descriptions

### 1. User Service

- **Purpose**: Provides basic user information based on User ID.
- **Implementation**:
  - Written in Go with gRPC for service-to-service calls.
  - Defines the `GetUser` RPC method to return user details.
- **Tools**: `gRPC`, `Docker`, `Kubernetes`.

#### Example Request

Using gRPC, a client can request user details by sending a `UserRequest` with the user ID, and the server responds with a `UserResponse` containing the user’s name.

### 2. Order Service

- **Purpose**: Generates orders for users and publishes an event on NATS.
- **Implementation**:
  - Calls the User Service using gRPC to retrieve user information.
  - Publishes an `order.created` event to NATS with order details.
- **Tools**: `gRPC`, `NATS`, `Docker`, `Kubernetes`.

#### Example Flow

1. Calls `GetUser` from User Service.
2. Publishes a message to the NATS `order.created` subject with order details.

### 3. Notification Service

- **Purpose**: Subscribes to the `order.created` subject on NATS and prints notifications when an order is created.
- **Implementation**:
  - Written in Node.js, using the `nats` package to connect to the NATS server.
- **Tools**: `NATS`, `Node.js`, `Docker`, `Kubernetes`.

#### Example Flow

1. Subscribes to `order.created` on NATS.
2. Logs the notification message whenever a new order event is received.

---

# Docker Setup

Each service is containerized with Docker. To build the images, use:

```bash
docker build -t user-service -f user-service/Dockerfile .
docker build -t order-service -f order-service/Dockerfile .
docker build -t notification-service -f notification-service/Dockerfile .
```

# Kubernetes Setup
The Kubernetes configurations (YAML files) for each service should be deployed in your cluster. Apply the configurations:

```bash
kubectl apply -f .\kubernetes\user-service.yaml
kubectl apply -f .\kubernetes\order-service.yaml
kubectl apply -f .\kubernetes\notification-service.yaml
```

# Running the Project

## Prerequisites

Steps
1. Build Docker images:

2. Deploy to Kubernetes: Apply each service’s deployment and service YAML files to set up networking and  scaling.

3. Verify the Workflow:
   - Trigger an order through the Order Service.
   - Confirm that the Notification Service receives the event and logs the message.