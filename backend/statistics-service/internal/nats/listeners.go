package nats

import (
	"log"

	"statistics-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

func ListenToOrderEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("order.created", func(msg *nats.Msg) {
		log.Println("Received order.created event")
		usecase.ProcessOrderCreated(msg.Data)
	})
	return err
}

func ListenToInventoryEvents(nc *nats.Conn, usecase usecase.StatisticsUsecase) error {
	_, err := nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		log.Println("Received inventory.updated event")
		usecase.ProcessInventoryUpdated(msg.Data)
	})
	return err
}
