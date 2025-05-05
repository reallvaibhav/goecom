package main

import (
	"log"
	"net"

	"statistics-service/internal/delivery"
	"statistics-service/internal/nats"
	"statistics-service/internal/repository"
	"statistics-service/internal/usecase"
	"statistics-service/proto"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting Statistics Service...")

	// Step 1: Initialize the database connection
	db, err := repository.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database connection initialized successfully.")

	// Step 2: Initialize the NATS connection
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer natsConn.Close()
	log.Println("Connected to NATS server successfully.")

	// Step 3: Initialize the use case layer
	statsUsecase := usecase.NewStatisticsUsecase(db)
	log.Println("Use case layer initialized successfully.")

	// Step 4: Start NATS message listeners
	go func() {
		log.Println("Starting listener for order events...")
		if err := nats.ListenToOrderEvents(natsConn, statsUsecase); err != nil {
			log.Fatalf("Error listening to order events: %v", err)
		}
	}()
	go func() {
		log.Println("Starting listener for inventory events...")
		if err := nats.ListenToInventoryEvents(natsConn, statsUsecase); err != nil {
			log.Fatalf("Error listening to inventory events: %v", err)
		}
	}()

	// Step 5: Start the gRPC server
	startGRPCServer(statsUsecase)
}

func startGRPCServer(usecase usecase.StatisticsUsecase) {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("Failed to listen on port 8084: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterStatisticsServiceServer(grpcServer, &delivery.Server{Usecase: usecase})

	log.Println("gRPC server running on port 8084")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
