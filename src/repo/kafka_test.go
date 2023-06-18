package repo

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Test_SendMessage(t *testing.T) {
	mockTopic := "stock-transaction"
	type args struct {
		Address string
		Message *kafka.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful set",
			args: args{
				Address: "localhost:29092",
				Message: &kafka.Message{
					Key:            []byte("key"),
					Value:          []byte("message4"),
					TopicPartition: kafka.TopicPartition{Topic: &mockTopic, Partition: kafka.PartitionAny},
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kaf, err := NewKafkaProducer(&KafkaOption{ServerHost: tt.args.Address})
			if (err != nil) != tt.wantErr {
				t.Errorf("kaf.NewKafkaProducer() err = %v, wantErr = %v", err, tt.wantErr)
			}

			defer func() {
				kaf.kafkaProducer.Close()
			}()

			err = kaf.SendMessage(tt.args.Message)
			if (err != nil) != tt.wantErr {
				t.Errorf("kaf.SendMessage() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
