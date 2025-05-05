FROM golang:1.23-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

# Create directory structure
RUN mkdir -p proto

# Copy proto files first
COPY proto/ proto/

# Copy the service files
COPY statistics-service/ .

# Generate proto files
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/statistics.proto

# Move generated files to the right place
RUN mkdir -p ecommerce/proto && mv proto/*.pb.go ecommerce/proto/

# Download dependencies and verify
RUN go mod download
RUN go mod tidy

# Build the application
RUN go build -o statistics-service

EXPOSE 50053
CMD ["./statistics-service"] 