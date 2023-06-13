package src

import (
	"fmt"
	"math"
	"strconv"
	"task1/entity"
	"task1/src/repo"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var (
	StoreSummaryStockAsKeyFormat = "summarystock:%s"
)

type OHLC struct {
	transactionLog map[string][]entity.MstTransaction
	summaryLog     map[string]entity.Summary
	Store          *repo.RedisOptions
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
		return summary, err
	}
	if errors.Is(err, redis.ErrNil) {
		err = nil
	}

	return summary, nil
}

func (rec OHLC) SetRedisSummaryLog(stockCode string, summary entity.Summary) (err error) {
	_, err = rec.Store.Del(stockCode)
	if err != nil {
		err = errors.Wrap(err, "GetSummaryLog")
		return err
	}

	_, err = rec.Store.HSetSummary(stockCode, summary)
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
	quantity, err := strconv.ParseInt(trx.Quantity, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecord")
		return
	}

	price, err := strconv.ParseInt(trx.Price, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecord")
		return
	}

	newEntry := entity.MstTransaction{
		Type:      trx.Type,
		Stock:     trx.Stock,
		OrderBook: trx.OrderBook,
		Quantity:  quantity,
		Price:     price,
	}

	if _, found := rec.transactionLog[trx.Stock]; !found {
		rec.transactionLog[trx.Stock] = []entity.MstTransaction{
			newEntry,
		}
	} else {
		rec.transactionLog[trx.Stock] = append(rec.transactionLog[trx.Stock], newEntry)
	}

	err = rec.CalculateRecordsByStockCode(newEntry)
	if err != nil {
		err = errors.Wrap(err, "InsertNewRecord")
	}

	return
}

func (rec OHLC) CalculateRecordsByStockCode(trx entity.MstTransaction) (err error) {
	summary, found := rec.summaryLog[trx.Stock]
	if !found {
		summary.LowestPrice = math.MaxInt64
	}

	summary, err = rec.GetRedisSummaryLog(trx.Stock)
	if err != nil {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
		return err
	}

	if trx.Type == "E" || trx.Type == "P" {
		summary.Volume += trx.Quantity
		summary.Value += trx.Quantity * trx.Price
		if summary.HighestPrice < trx.Price {
			summary.HighestPrice = trx.Price
		}
		if summary.LowestPrice > trx.Price {
			summary.LowestPrice = trx.Price
		}

		if summary.IsNewDay {
			summary.OpenPrice = trx.Price
			summary.IsNewDay = false
		}

		summary.ClosePrice = trx.Price
	}
	if trx.Quantity == 0 {
		summary.PreviousPrice = trx.Price
		summary.IsNewDay = true
	}

	rec.summaryLog[trx.Stock] = summary
	err = rec.SetRedisSummaryLog(trx.Stock, summary)
	if err != nil {
		err = errors.Wrap(err, "CalculateRecordsByStockCode")
		return err
	}

	return nil

}
