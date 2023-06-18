package repo

import (
	"context"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var GetKafkaProduceFunc func(msg *kafka.Message, deliveryChan chan kafka.Event) error
var GetKafkaFlushFunc func(timeoutMs int) int
var GetKafkaIsClosedFunc func() bool
var GetKafkaStringFunc func() string
var GetKafkaEventFunc func() chan kafka.Event
var GetKafkaLogsFunc func() chan kafka.LogEvent
var GetKafkaProduceChannelFunc func() chan *kafka.Message
var GetKafkaLenFunc func() int
var GetKafkaCloseFunc func()
var GetKafkaPurgeFunc func(flags int) error
var GetKafkaNewProducerFunc func(conf *kafka.ConfigMap) (*kafka.Producer, error)
var GetKafkaGetMetadataFunc func(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
var GetKafkaQueryWatermarkOffsets func(topic string, partition int32, timeoutMs int) (low, high int64, err error)
var GetKafkaTestFatalErrorFunc func(code kafka.ErrorCode, str string) kafka.ErrorCode
var GetKafkaSetOAuthBearerTokenFunc func(oauthBearerToken kafka.OAuthBearerToken) error
var GetKafkaSetOAuthBearerTokenFailureFunc func(errstr string) error
var GetKafkaInitTransactionsFunc func(ctx context.Context) error
var GetKafkaBeginTransactionFunc func() error
var GetKafkaSendOffsetsToTransactionFunc func(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error
var GetKafkaCommitTransactionFunc func(ctx context.Context) error
var GetKafkaAbortTransactionFunc func(ctx context.Context) error
var GetGetFatalErrorFunc func() error
var GetOffsetsForTimesFunc func(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
var GetKafkaSetSaslCredentialsFunc func(username string, password string) error

type MockKafkaProducer struct {
	GetProduceFunc                         func(msg *kafka.Message, deliveryChan chan kafka.Event) error
	GetKafkaFlushFunc                      func(timeoutMs int) int
	GetIsClosedFunc                        func() bool
	GetStringFunc                          func() string
	GetKafkaEventFunc                      func() chan kafka.Event
	GetKafkaLogsFunc                       func() chan kafka.LogEvent
	GetKafkaProduceChannelFunc             func() chan *kafka.Message
	GetKafkaLenFunc                        func() int
	GetKafkaCloseFunc                      func()
	GetKafkaPurge                          func(flags int) error
	GetKafkaNewProducerFunc                func(conf *kafka.ConfigMap) (*kafka.Producer, error)
	GetKafkaGetMetadataFunc                func(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	GetKafkaQueryWatermarkOffsets          func(topic string, partition int32, timeoutMs int) (low, high int64, err error)
	GetKafkaTestFatalErrorFunc             func(code kafka.ErrorCode, str string) kafka.ErrorCode
	GetKafkaSetOAuthBearerTokenFunc        func(oauthBearerToken kafka.OAuthBearerToken) error
	GetKafkaSetOAuthBearerTokenFailureFunc func(errstr string) error
	GetKafkaInitTransactionsFunc           func(ctx context.Context) error
	GetKafkaBeginTransactionFunc           func() error
	GetKafkaSendOffsetsToTransactionFunc   func(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error
	GetKafkaCommitTransactionFunc          func(ctx context.Context) error
	GetKafkaAbortTransactionFunc           func(ctx context.Context) error
	GetGetFatalErrorFunc                   func() error
	GetOffsetsForTimesFunc                 func(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	GetKafkaSetSaslCredentialsFunc         func(username string, password string) error
}

func (p *MockKafkaProducer) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	return GetKafkaProduceFunc(msg, deliveryChan)
}

func (p *MockKafkaProducer) Flush(timeoutMs int) int {
	return GetKafkaFlushFunc(timeoutMs)
}

func (p *MockKafkaProducer) IsClosed() bool {
	return GetKafkaIsClosedFunc()
}

func (p *MockKafkaProducer) String() string {
	return GetKafkaStringFunc()
}

func (p *MockKafkaProducer) Events() chan kafka.Event {
	return GetKafkaEventFunc()
}

func (p *MockKafkaProducer) Logs() chan kafka.LogEvent {
	return GetKafkaLogsFunc()
}

func (p *MockKafkaProducer) ProduceChannel() chan *kafka.Message {
	return GetKafkaProduceChannelFunc()
}

func (p *MockKafkaProducer) Len() int {
	return GetKafkaLenFunc()
}

func (p *MockKafkaProducer) Close() {
	GetKafkaCloseFunc()
}

func (p *MockKafkaProducer) Purge(flag int) error {
	return GetKafkaPurgeFunc(flag)
}

func (p *MockKafkaProducer) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	return GetKafkaGetMetadataFunc(topic, allTopics, timeoutMs)
}

func (p *MockKafkaProducer) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
	return GetKafkaQueryWatermarkOffsets(topic, partition, timeoutMs)
}

func (p *MockKafkaProducer) TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode {
	return GetKafkaTestFatalErrorFunc(code, str)
}

func (p *MockKafkaProducer) SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error {
	return GetKafkaSetOAuthBearerTokenFunc(oauthBearerToken)
}

func (p *MockKafkaProducer) SetOAuthBearerTokenFailure(errstr string) error {
	return GetKafkaSetOAuthBearerTokenFailureFunc(errstr)
}

func (p *MockKafkaProducer) InitTransactions(ctx context.Context) error {
	return GetKafkaInitTransactionsFunc(ctx)
}

func (p *MockKafkaProducer) BeginTransaction() error {
	return GetKafkaBeginTransactionFunc()
}

func (p *MockKafkaProducer) SendOffsetsToTransaction(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error {
	return GetKafkaSendOffsetsToTransactionFunc(ctx, offsets, consumerMetadata)
}

func (p *MockKafkaProducer) CommitTransaction(ctx context.Context) error {
	return GetKafkaCommitTransactionFunc(ctx)
}

func (p *MockKafkaProducer) AbortTransaction(ctx context.Context) error {
	return GetKafkaAbortTransactionFunc(ctx)
}

func (p *MockKafkaProducer) GetFatalError() error {
	return GetGetFatalErrorFunc()
}

func (p *MockKafkaProducer) OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error) {
	return GetOffsetsForTimesFunc(times, timeoutMs)
}

func (p *MockKafkaProducer) SetSaslCredentials(username string, password string) error {
	return GetKafkaSetSaslCredentialsFunc(username, password)
}

func Test_SendMessage(t *testing.T) {

	mockKafkaProducer := MockKafkaProducer{}

	type args struct {
		Address string
		Message *kafka.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name: "Successful set",
			args: args{
				Address: "",
				Message: &kafka.Message{
					Key:   []byte{},
					Value: []byte{},
				},
			},
			patch: func() {

				GetKafkaEventFunc = func() chan kafka.Event {
					e := make(chan kafka.Event, 1)
					return e
				}

				GetKafkaProduceFunc = func(msg *kafka.Message, deliveryChan chan kafka.Event) error {
					return nil
				}

				GetKafkaFlushFunc = func(timeoutMs int) int {
					return 1
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			kafkaOption := &KafkaOption{
				kafkaProducer:  &mockKafkaProducer,
				topicPartition: tt.args.Message.TopicPartition,
			}
			tt.patch()
			err := kafkaOption.SendMessage(tt.args.Message)
			if (err != nil) != tt.wantErr {
				t.Errorf("kaf.SendMessage() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
