package entity

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaHandler interface {
	HandleMessage(message *kafka.Message)
}

type HandlerFunc func(message *kafka.Message) error

func (h HandlerFunc) HandleMessage(m *kafka.Message) error {
	return h(m)
}
