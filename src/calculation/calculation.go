package calculation

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"task1/entity"
	"task1/src/filereader"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var (
	StoreSummaryStockAsKeyFormat = "summarystock:%s"

	filereaderConvertToStruct func(line []byte) (transaction entity.Transaction, err error)

	ohlcInsertNewRecord func(trx entity.Transaction) (err error)
)

type OHLC struct {
	transactionLog map[string][]entity.MstTransaction
	summaryLog     map[string]entity.Summary
	Store          entity.StorageInterface
}

func NewOHLCRecords(records *OHLC) *OHLC {
	newLogs := make(map[string][]entity.MstTransaction)
	newSummary := make(map[string]entity.Summary)

	records = &OHLC{
		transactionLog: newLogs,
		summaryLog:     newSummary,
		Store:          records.Store,
	}

	filereaderConvertToStruct = filereader.ConvertToStruct
	ohlcInsertNewRecord = records.InsertNewRecord

	return records
}

func (rec OHLC) GetTransactionLog(stockCode string) []entity.MstTransaction {
	return rec.transactionLog[stockCode]
}

func (rec OHLC) GetRedisSummaryLog(stockCode string) (summary entity.Summary, err error) {
	summary, err = rec.Store.HGetSummary(fmt.Sprintf(StoreSummaryStockAsKeyFormat, stockCode))
	if err != nil {
		err = errors.Wrap(err, "GetSummaryLog")
	}

	return summary, err
}

func (rec OHLC) SetRedisSummaryField(stockCode string, fieldName string, value interface{}) (err error) {
	_, err = rec.Store.HSet(fmt.Sprintf(StoreSummaryStockAsKeyFormat, stockCode), fieldName, value)
	if err != nil {
		err = errors.Wrap(err, "SetRedisSummaryField")
	}

	return err
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

func (rec OHLC) GetSummaryLog(stockCode string) (summary entity.Summary, err error) {
	summary, err = rec.GetRedisSummaryLog(stockCode)
	if err != nil {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
	}

	return
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
	isFound := true
	summary, err := rec.GetSummaryLog(trx.Stock)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
		return err
	}
	if err != nil && errors.Is(err, redis.ErrNil) {
		isFound = false
		summary.LowestPrice = math.MaxInt32
		err = nil
	}

	if trx.Type == "E" || trx.Type == "P" {
		summary.Volume += trx.ExecutedQuantity
		summary.Value += trx.ExecutedQuantity * trx.ExecutedPrice
		if isFound {
			err = rec.SetRedisSummaryField(trx.Stock, "volume", summary.Volume)
			if err != nil {
				err = errors.Wrap(err, "CalculateRecordsByStockCode")
				return err
			}
			err = rec.SetRedisSummaryField(trx.Stock, "value", summary.Value)
			if err != nil {
				err = errors.Wrap(err, "CalculateRecordsByStockCode")
				return err
			}
		}
		if summary.HighestPrice < trx.ExecutedPrice {
			summary.HighestPrice = trx.ExecutedPrice
			if isFound {
				err = rec.SetRedisSummaryField(trx.Stock, "highest_price", summary.HighestPrice)
				if err != nil {
					err = errors.Wrap(err, "CalculateRecordsByStockCode")
					return err
				}
			}
		}
		if summary.LowestPrice > trx.ExecutedPrice {
			summary.LowestPrice = trx.ExecutedPrice
			if isFound {
				err = rec.SetRedisSummaryField(trx.Stock, "lowest_price", summary.LowestPrice)
				if err != nil {
					err = errors.Wrap(err, "CalculateRecordsByStockCode")
					return err
				}
			}
		}

		if summary.IsNewDay > 0 {
			summary.OpenPrice = trx.ExecutedPrice
			summary.IsNewDay = 0
			if isFound {
				err = rec.SetRedisSummaryField(trx.Stock, "open_price", summary.OpenPrice)
				if err != nil {
					err = errors.Wrap(err, "CalculateRecordsByStockCode")
					return err
				}
				err = rec.SetRedisSummaryField(trx.Stock, "is_new_day", summary.IsNewDay)
				if err != nil {
					err = errors.Wrap(err, "CalculateRecordsByStockCode")
					return err
				}
			}
		}

		summary.ClosePrice = trx.ExecutedPrice
		if isFound {
			err = rec.SetRedisSummaryField(trx.Stock, "close_price", summary.ClosePrice)
			if err != nil {
				err = errors.Wrap(err, "CalculateRecordsByStockCode")
				return err
			}
		}
	}
	if (trx.Type == "E" || trx.Type == "P") && trx.ExecutedQuantity == 0 ||
		(trx.Type != "E" && trx.Type != "P") && trx.Quantity == 0 {
		summary.PreviousPrice = trx.Price
		summary.IsNewDay = 1
		if isFound {
			err = rec.SetRedisSummaryField(trx.Stock, "previous_price", summary.PreviousPrice)
			if err != nil {
				err = errors.Wrap(err, "CalculateRecordsByStockCode")
				return err
			}
			err = rec.SetRedisSummaryField(trx.Stock, "is_new_day", summary.IsNewDay)
			if err != nil {
				err = errors.Wrap(err, "CalculateRecordsByStockCode")
				return err
			}
		}
	}

	log.Default().Print(summary, " stock: ", trx.Stock, " type: ", trx.Type)
	if !isFound {
		err = rec.SetRedisSummaryLog(trx.Stock, summary)
		if err != nil {
			err = errors.Wrap(err, "CalculateRecordsByStockCode")
			return err
		}
	}

	return nil

}

func (rec OHLC) InsertNewRecordFromKafka(msg *kafka.Message) (err error) {

	transaction, err := filereaderConvertToStruct(msg.Value)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecordFromKafka"+string(msg.Value))
		log.Fatal(err)
		return
	}

	err = ohlcInsertNewRecord(transaction)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecordFromKafka")
		log.Fatal(err)
		return
	}

	return nil
}
