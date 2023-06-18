package calculation

import (
	"fmt"
	"log"
	"strconv"
	"task1/entity"
	"task1/src/filereader"
	"task1/src/repo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var (
	StoreSummaryStockAsKeyFormat = "summarystock:%s"
)

type OHLC struct {
	transactionLog map[string][]entity.MstTransaction
	summaryLog     map[string]entity.Summary
	Store          *repo.Storage
}

func NewOHLCRecords(records *OHLC) *OHLC {
	newLogs := make(map[string][]entity.MstTransaction)
	newSummary := make(map[string]entity.Summary)

	return &OHLC{
		transactionLog: newLogs,
		summaryLog:     newSummary,
		Store:          records.Store,
	}
}

func (rec OHLC) GetTransactionLog(stockCode string) []entity.MstTransaction {
	return rec.transactionLog[stockCode]
}

func (rec OHLC) GetRedisSummaryLog(stockCode string) (summary entity.Summary, err error) {
	summary, err = rec.Store.HGetSummary(fmt.Sprintf(StoreSummaryStockAsKeyFormat, stockCode))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		err = errors.Wrap(err, "GetSummaryLog")
	}

	return summary, err
}

func (rec OHLC) SetRedisSummaryLog(stockCode string, summary entity.Summary) (err error) {
	_, err = rec.Store.Del(fmt.Sprintf(StoreSummaryStockAsKeyFormat, stockCode))
	if err != nil {
		err = errors.Wrap(err, "GetSummaryLog")
		return err
	}

	_, err = rec.Store.HSetSummary(fmt.Sprintf(StoreSummaryStockAsKeyFormat, stockCode), summary)
	if err != nil {
		err = errors.Wrap(err, "GetSummaryLog")
		return err
	}
	return nil
}

func (rec OHLC) GetSummaryLog(stockCode string) entity.Summary {
	return rec.summaryLog[stockCode]
}

func (rec OHLC) SetSummaryLog(stockCode string, summary entity.Summary) entity.Summary {
	rec.summaryLog[stockCode] = summary
	return summary
}

func (rec OHLC) InsertNewRecord(trx entity.Transaction) (err error) {

	var quantity int64 = 0
	if trx.Quantity != "" {
		quantity, err = strconv.ParseInt(trx.Quantity, 10, 64)
		if err != nil {
			err = errors.Wrap(err, "InsertNewRecord. Quantity"+trx.Quantity)
			return
		}
	}

	var executedQuantity int64 = 0
	if trx.ExecutedQuantity != "" {
		executedQuantity, err = strconv.ParseInt(trx.ExecutedQuantity, 10, 64)
		if err != nil {
			err = errors.Wrap(err, "InsertNewRecord. Executed Quantity"+trx.ExecutedQuantity)
			return
		}
	}

	var price int64
	if trx.Price != "" {
		price, err = strconv.ParseInt(trx.Price, 10, 64)
		if err != nil {
			err = errors.Wrap(err, "InsertNewRecord. Price ="+trx.Price)
			return
		}
	}

	var executedPrice int64 = 0
	if trx.ExecutedPrice != "" {
		executedPrice, err = strconv.ParseInt(trx.ExecutedPrice, 10, 64)
		if err != nil {
			err = errors.Wrap(err, "InsertNewRecord. Executed Price ="+trx.ExecutedPrice)
			return
		}
	}

	newEntry := entity.MstTransaction{
		Type:             trx.Type,
		Stock:            trx.Stock,
		Quantity:         quantity,
		ExecutedQuantity: executedQuantity,
		Price:            price,
		ExecutedPrice:    executedPrice,
	}

	err = rec.CalculateRecordsByStockCode(newEntry)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecord")
	}

	return
}

func (rec OHLC) CalculateRecordsByStockCode(trx entity.MstTransaction) (err error) {
	summary, err := rec.GetRedisSummaryLog(trx.Stock)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
		return err
	}
	if err != nil && errors.Is(err, redis.ErrNil) {
		summary.LowestPrice = 0
		err = nil
	}

	if trx.Type == "E" || trx.Type == "P" {
		summary.Volume += trx.ExecutedQuantity
		summary.Value += trx.ExecutedQuantity * trx.ExecutedPrice
		if summary.HighestPrice < trx.ExecutedPrice {
			summary.HighestPrice = trx.ExecutedPrice
		}
		if summary.LowestPrice > trx.ExecutedPrice {
			summary.LowestPrice = trx.ExecutedPrice
		}

		if summary.IsNewDay > 0 {
			summary.OpenPrice = trx.ExecutedPrice
			summary.IsNewDay = 0
		}

		summary.ClosePrice = trx.ExecutedPrice
	}
	if (trx.Type == "E" || trx.Type == "P") && trx.ExecutedQuantity == 0 ||
		!(trx.Type == "E" || trx.Type == "P") && trx.Quantity == 0 {
		summary.PreviousPrice = trx.Price
		summary.IsNewDay = 1
	}

	log.Default().Print(summary, " stock: ", trx.Stock, " type: ", trx.Type)
	err = rec.SetRedisSummaryLog(trx.Stock, summary)
	if err != nil {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
		return err
	}

	return nil

}

func (rec OHLC) InsertNewRecordFromKafka(msg *kafka.Message) (err error) {

	transaction, err := filereader.ConvertToStruct(msg.Value)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecordFromKafka"+string(msg.Value))
		log.Fatal(err)
		return
	}

	err = rec.InsertNewRecord(transaction)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecordFromKafka")
		log.Fatal(err)
		return
	}

	return nil
}
