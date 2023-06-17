package repo

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pkg/errors"
)

type KafkaOption struct {
	ServerHost    string
	kafkaProducer *kafka.Producer
}

func NewKafkaProducer(opt *KafkaOption) (*KafkaOption, error) {
	if len(opt.ServerHost) == 0 {
		return opt, fmt.Errorf("initialization failed for NewKafkaProducer: no host provided")
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": opt.ServerHost})
	if err != nil {
		return opt, err
	}

	opt.kafkaProducer = producer

	return opt, nil
}

func (k *KafkaOption) SendMessage(message *kafka.Message) (err error) {
	err = k.kafkaProducer.Produce(message, nil)
	if err != nil {
		return errors.Wrap(err, "kafka.SendMessage")
	}

	k.kafkaProducer.Flush(15 * 1000)

	return nil
}
