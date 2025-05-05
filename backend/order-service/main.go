package main

import (
	"ecommerce/order-service/handlers"
	"ecommerce/order-service/repository"
	"ecommerce/order-service/services"
	"ecommerce/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	repo, err := repository.NewRepository("./order.db")
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	service := services.NewOrderService(repo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, handlers.NewOrderServer(service))

	log.Println("Order Service running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
