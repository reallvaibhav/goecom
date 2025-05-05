FROM golang:1.22-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

WORKDIR /app/ecommerce

# Copy the entire ecommerce directory
COPY . .

# Enable CGO and build
ENV CGO_ENABLED=1
ENV GOOS=linux
RUN cd order-service && go build -o order-service

EXPOSE 50051

WORKDIR /app/ecommerce/order-service
CMD ["./order-service"] 