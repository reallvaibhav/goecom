FROM golang:1.22-alpine

WORKDIR /app/ecommerce

# Copy the entire ecommerce directory
COPY . .

RUN go mod download
RUN cd api-gateway && go build -o api-gateway

EXPOSE 8080

WORKDIR /app/ecommerce/api-gateway
CMD ["./api-gateway"] 