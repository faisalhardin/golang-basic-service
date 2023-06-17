package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	numOfConsumer := 2

	generateConsumer(fmt.Sprintf("groupID%v", numOfConsumer))
}

func generateConsumer(groupID string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "same",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("/tmp/log-kafka.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	c.SubscribeTopics([]string{"stock-transaction"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		var out string
		msg, err := c.ReadMessage(time.Second)
		time.Sleep(time.Second * 1)
		if err == nil {
			out = fmt.Sprintf("Message on groupID %s %s: %s #%s\n", groupID, msg.TopicPartition, string(msg.Value), string(msg.Key))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			out = fmt.Sprintf("Consumer groupID %s error: %v (%v)\n", groupID, err, msg)
		}

		if _, err = f.WriteString(out); err != nil {
			panic(err)
		}

	}

	c.Close()
}
