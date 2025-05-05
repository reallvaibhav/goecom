package client

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NATSClient struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

func NewNATSClient(url string) (*NATSClient, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, err
	}

	return &NATSClient{
		conn: nc,
		js:   js,
	}, nil
}

func (c *NATSClient) Close() {
	c.conn.Close()
}

func (c *NATSClient) PublishOrderCreated(orderData []byte) error {
	_, err := c.js.Publish("order.created", orderData)
	return err
}

func (c *NATSClient) PublishInventoryUpdated(inventoryData []byte) error {
	_, err := c.js.Publish("inventory.updated", inventoryData)
	return err
}

func (c *NATSClient) SubscribeOrderCreated(handler func([]byte)) error {
	_, err := c.js.Subscribe("order.created", func(msg *nats.Msg) {
		handler(msg.Data)
		msg.Ack()
	})
	return err
}

func (c *NATSClient) SubscribeInventoryUpdated(handler func([]byte)) error {
	_, err := c.js.Subscribe("inventory.updated", func(msg *nats.Msg) {
		handler(msg.Data)
		msg.Ack()
	})
	return err
} 