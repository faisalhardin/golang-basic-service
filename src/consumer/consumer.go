package main

import (
	"os"
	"sync"
	"task1/src/calculation"
	"task1/src/repo"
)

var (
	numOfConsumer = 3
)

func main() {
	i := StartConsumer()
	os.Exit(i)
}

func StartConsumer() int {
	redisRepo := repo.NewRedisRepo(&repo.RedisOptions{
		Address: "127.0.0.1:6379",
	})

	ohlc := calculation.NewOHLCRecords(&calculation.OHLC{
		Store: redisRepo,
	})

	consumer, err := repo.NewKafkaConsumer(repo.KafkaOption{
		ServerHost:      "localhost:29092",
		ConsumerGroupID: "stock-summary",
	})
	if err != nil {
		panic(err)
	}

	consumer.RegisterHandler("stock-transaction", ohlc.InsertNewRecordFromKafka, numOfConsumer)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(this *repo.KafkaConsumer) {

		for {
			if consumer.IsClosed() {
				wg.Done()
			}
		}

	}(&consumer)

	wg.Wait()

	return 0

}
