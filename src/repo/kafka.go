package repo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"task1/entity"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pkg/errors"
)

type KafkaOption struct {
	ServerHost      string
	Topic           string
	ConsumerGroupID string

	kafkaProducer  *kafka.Producer
	topicPartition kafka.TopicPartition
}

type KafkaConsumer struct {
	consumer        *kafka.Consumer
	opt             KafkaOption
	runningHandlers int32
	exitHandler     sync.Once
}

func NewKafkaProducer(opt *KafkaOption) (*KafkaOption, error) {
	if len(opt.ServerHost) == 0 {
		return opt, fmt.Errorf("initialization failed for NewKafkaProducer: no host provided")
	}

	opt.topicPartition = kafka.TopicPartition{
		Topic:     &opt.Topic,
		Partition: kafka.PartitionAny,
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": opt.ServerHost})
	if err != nil {
		return opt, err
	}

	opt.kafkaProducer = producer

	return opt, nil
}

func (k *KafkaOption) SendMessage(message *kafka.Message) (err error) {
	message.TopicPartition = k.topicPartition

	go func() {
		for e := range k.kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	err = k.kafkaProducer.Produce(message, nil)
	if err != nil {
		return errors.Wrap(err, "kafka.SendMessage")
	}

	k.kafkaProducer.Flush(15 * 1000)

	return nil
}

func (k *KafkaOption) CloseProducer() {
	k.kafkaProducer.Close()
}

func NewKafkaConsumer(opt KafkaOption) (kc KafkaConsumer, err error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  opt.ServerHost,
		"group.id":           opt.ConsumerGroupID,
		"auto.offset.reset":  "latest",
		"enable.auto.commit": true,
	})
	if err != nil {
		return kc, err
	}

	kc.consumer = c
	kc.opt = opt

	return
}

func (kc *KafkaConsumer) RegisterHandler(topic string, handler entity.HandlerFunc, concurrency int) {
	kc.consumer.SubscribeTopics([]string{topic}, nil)

	atomic.AddInt32(&kc.runningHandlers, int32(concurrency))
	for i := 0; i < concurrency; i++ {
		go kc.handlerLoop(handler)
	}

}

func (kc *KafkaConsumer) IsClosed() bool {
	return kc.consumer.IsClosed()
}

func (kc *KafkaConsumer) exit() {
	kc.exitHandler.Do(func() {
		kc.consumer.Close()
	})
}

func (kc *KafkaConsumer) handlerLoop(handler entity.HandlerFunc) {
	for {
		msg, err := kc.consumer.ReadMessage(time.Second * 2)
		if err == nil {
			handler.HandleMessage(msg)
		} else if !err.(kafka.Error).IsTimeout() {
			if atomic.AddInt32(&kc.runningHandlers, -1) == 0 {
				kc.exit()
			}

			break

		}

	}

}
