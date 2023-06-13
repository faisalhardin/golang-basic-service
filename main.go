package main

import (
	"fmt"
	"log"

	src "task1/src"
	"task1/src/repo"
)

func main() {

	redisRepo := repo.NewRedisRepo(&repo.RedisOptions{
		Address: "127.0.0.1:6379",
	})

	ohlc := src.NewOHLCRecords(&src.OHLC{
		Store: redisRepo,
	})

	listOfFiles, err := src.ListFiles("./subsetdata/")
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range listOfFiles {
		trxLogs, err := src.ReadFiles(fmt.Sprintf("./subsetdata/%s", fileName))
		if err != nil {
			log.Default().Print(err)
		}

		for _, trx := range trxLogs {
			log.Default().Print(trx)
			err = ohlc.InsertNewRecord(trx)
			if err != nil {
				log.Fatal(err)
			}
		}

		summ, err := ohlc.GetRedisSummaryLog("BBCA")
		if err != nil {
			log.Fatal(err)
		}

		log.Default().Print(summ)
	}

}
