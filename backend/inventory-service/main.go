package main

import (
	"ecommerce/inventory-service/handlers"
	"ecommerce/inventory-service/repository"
	"ecommerce/inventory-service/services"
	"ecommerce/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	repo, err := repository.NewRepository("./inventory.db")
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	service := services.NewProductService(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterInventoryServiceServer(s, handlers.NewInventoryServer(service))

	log.Println("Inventory Service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
