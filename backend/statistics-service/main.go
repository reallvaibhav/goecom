package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "statistics-service/ecommerce/proto"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type StatisticsServer struct {
	pb.UnimplementedStatisticsServiceServer
	collection *mongo.Collection
	nc         *nats.Conn
}

func (s *StatisticsServer) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {
	var result struct {
		TotalOrders int32    `bson:"total_orders"`
		OrderIDs    []string `bson:"order_ids"`
	}
	err := s.collection.FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.UserOrderStatisticsResponse{
				TotalOrders: 0,
				OrderIds:    []string{},
			}, nil
		}
		return nil, err
	}
	return &pb.UserOrderStatisticsResponse{
		TotalOrders: result.TotalOrders,
		OrderIds:    result.OrderIDs,
	}, nil
}

func main() {
	fmt.Println("Statistics Service starting...")
	
	// Connect to MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:example@localhost:27017/statistics?authSource=admin"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Ping MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to MongoDB at %s", mongoURI)

	collection := client.Database("statistics").Collection("user_stats")

	// Connect to NATS
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	// Connect to NATS server with retries
	var nc *nats.Conn
	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(natsURL)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to NATS, retrying in 2 seconds... (attempt %d/5)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	log.Printf("Successfully connected to NATS at %s", natsURL)

	// Create statistics server
	statsServer := &StatisticsServer{
		collection: collection,
		nc:         nc,
	}

	// Subscribe to NATS subjects
	nc.Subscribe("order.created", func(msg *nats.Msg) {
		log.Printf("Received order creation event: %s", string(msg.Data))
		var orderEvent struct {
			UserID  string `json:"user_id"`
			OrderID string `json:"order_id"`
		}
		if err := json.Unmarshal(msg.Data, &orderEvent); err != nil {
			log.Printf("Error parsing order event: %v", err)
			return
		}

		// Update MongoDB with order statistics
		update := bson.M{
			"$inc": bson.M{"total_orders": 1},
			"$push": bson.M{"order_ids": orderEvent.OrderID},
		}
		opts := options.Update().SetUpsert(true)
		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"user_id": orderEvent.UserID},
			update,
			opts,
		)
		if err != nil {
			log.Printf("Error updating statistics: %v", err)
		}
	})

	nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		log.Printf("Received inventory update event: %s", string(msg.Data))
		// Process inventory event if needed
	})

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(grpcServer, statsServer)

	log.Printf("gRPC server listening on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
} 