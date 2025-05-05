# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o nats-service .

# Stage 2: Create a minimal image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/nats-service .

# Expose the ports used by the NATS service (adjust as needed)
EXPOSE 4222
EXPOSE 8222

# Command to run the application
CMD ["./nats-service"]
