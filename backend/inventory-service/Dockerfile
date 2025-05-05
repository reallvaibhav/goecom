FROM golang:1.22-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

WORKDIR /app/ecommerce

# Copy the entire ecommerce directory
COPY . .

# Enable CGO and build
ENV CGO_ENABLED=1
ENV GOOS=linux
RUN cd inventory-service && go build -o inventory-service

EXPOSE 50052

WORKDIR /app/ecommerce/inventory-service
CMD ["./inventory-service"] 