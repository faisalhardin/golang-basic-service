package main

import (
	"log"

	src "task1/src"
	"task1/src/repo"
)

func main() {

	redisRepo := repo.NewRedisRepo(&repo.RedisOptions{
		Address: "127.0.0.1:6379",
	})

	// _, err := src.ListFiles("./subsetdata/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// stop := make(chan bool, 1)
	// transactionLine := make(chan []byte)
	trxLogs, err := src.ReadFiles("")
	if err != nil {
		log.Default().Print(err)
	}

	ohlc := src.NewOHLCRecords(&src.OHLC{
		Store: redisRepo, 
	})

	for _, trxLog := range trxLogs {
		ohlc.
	}

	log.Default().Print(trxLogs)

}
