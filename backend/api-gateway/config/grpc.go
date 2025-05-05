package config

import (
	"ecommerce/proto"
	"google.golang.org/grpc"
	"log"
)

type Clients struct {
	Inventory     proto.InventoryServiceClient
	Order         proto.OrderServiceClient
	inventoryConn *grpc.ClientConn
	orderConn     *grpc.ClientConn
}

func NewClients() (*Clients, error) {
	inventoryConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to inventory service: %v", err)
		return nil, err
	}

	orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to order service: %v", err)
		inventoryConn.Close()
		return nil, err
	}

	return &Clients{
		Inventory:     proto.NewInventoryServiceClient(inventoryConn),
		Order:         proto.NewOrderServiceClient(orderConn),
		inventoryConn: inventoryConn,
		orderConn:     orderConn,
	}, nil
}

func (c *Clients) Close() {
	if c.inventoryConn != nil {
		c.inventoryConn.Close()
	}
	if c.orderConn != nil {
		c.orderConn.Close()
	}
}
