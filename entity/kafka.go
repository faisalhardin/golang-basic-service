package entity

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaHandler interface {
	HandleMessage(message *kafka.Message)
}

type HandlerFunc func(message *kafka.Message) error

func (h HandlerFunc) HandleMessage(m *kafka.Message) error {
	return h(m)
}

type KafkaInterface interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Flush(timeoutMs int) int
	IsClosed() bool
	String() string
	Events() chan kafka.Event
	Logs() chan kafka.LogEvent
	ProduceChannel() chan *kafka.Message
	Len() int
	Close()
	Purge(flags int) error
	GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error)
	TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode
	SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error
	SetOAuthBearerTokenFailure(errstr string) error
	InitTransactions(ctx context.Context) error
	BeginTransaction() error
	SendOffsetsToTransaction(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error
	CommitTransaction(ctx context.Context) error
	AbortTransaction(ctx context.Context) error
	GetFatalError() error
	OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	SetSaslCredentials(username string, password string) error
}
