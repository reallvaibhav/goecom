FROM golang:1.20

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o statistics-service .

# Expose the port
EXPOSE 50051

# Run the application
CMD ["./statistics-service"]