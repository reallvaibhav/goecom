package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("NATS Service starting...")
	// Get NATS URL from environment or use default
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	// Connect to NATS server with retries
	var nc *nats.Conn
	var err error
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

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Create stream if it doesn't exist
	streamName := "ECOMMERCE"
	streamSubjects := []string{
		"order.*",
		"inventory.*",
	}

	// Check if stream exists, if not create it
	_, err = js.StreamInfo(streamName)
	if err != nil {
		// Stream doesn't exist, create it
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: streamSubjects,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created stream: %s", streamName)
	} else {
		log.Printf("Stream already exists: %s", streamName)
	}

	// Subscribe to order events
	js.Subscribe("order.created", func(msg *nats.Msg) {
		log.Printf("Received order created event: %s", string(msg.Data))
		// Process the order event
		msg.Ack()
	})

	// Subscribe to inventory events
	js.Subscribe("inventory.updated", func(msg *nats.Msg) {
		log.Printf("Received inventory update event: %s", string(msg.Data))
		// Process the inventory event
		msg.Ack()
	})

	log.Printf("NATS service is running, connected to %s...", natsURL)

	// Keep the service running
	select {}
} 