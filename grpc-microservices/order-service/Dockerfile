# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the shared generated folder
COPY generated/ ./generated/

# Copy the user-service source files
COPY order-service/ .

# Build the Go application
RUN go build -o order-service .

# Stage 2: Create a minimal container
FROM alpine:3.14

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/order-service .

# Command to run the executable
CMD ["./order-service"]
