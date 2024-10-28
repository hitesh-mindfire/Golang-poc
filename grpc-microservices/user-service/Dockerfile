FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY generated/ ./generated/

COPY user-service/ .

RUN go build -o user-service .

FROM alpine:3.14

WORKDIR /root/

COPY --from=builder /app/user-service .

CMD ["./user-service"]
