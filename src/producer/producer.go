package main

import (
	"log"
	"os"
	"task1/src/filereader"
	"task1/src/repo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	i := StartProducer()
	os.Exit(i)
}

func StartProducer() int {

	messagingProducers, err := repo.NewKafkaProducer(&repo.KafkaOption{
		ServerHost: "localhost:29092",
		Topic:      "stock-transaction",
	})
	if err != nil {
		panic(err)
	}

	defer messagingProducers.CloseProducer()

	listOfFiles, err := filereader.ListFiles("./subsetdata/")
	if err != nil {
		log.Fatal(err)
	}

	trxLogs := filereader.ReadFilesWithChannel("./subsetdata", listOfFiles)
	for record := range trxLogs {
		messagingProducers.SendMessage(&kafka.Message{
			Key:   []byte(record.StockCode),
			Value: []byte(record.TransactionLog),
		})
	}

	return 0

}
