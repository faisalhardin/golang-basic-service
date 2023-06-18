package main

import (
	"log"
	src "task1/src"
	"task1/src/repo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	messagingProducers, err := repo.NewKafkaProducer(&repo.KafkaOption{
		ServerHost: "localhost:29092",
		Topic:      "stock-transaction",
	})
	if err != nil {
		panic(err)
	}

	listOfFiles, err := src.ListFiles("./subsetdata/")
	if err != nil {
		log.Fatal(err)
	}

	trxLogs := src.ReadFilesWithChannel("./subsetdata", listOfFiles)
	for record := range trxLogs {
		messagingProducers.SendMessage(&kafka.Message{
			Key:   []byte(record.StockCode),
			Value: []byte(record.TransactionLog),
		})
	}

}
