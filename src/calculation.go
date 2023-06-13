package src

import (
	"math"
	"strconv"
	"task1/entity"
)

type OHLC struct {
	transactionLog map[string][]entity.MstTransaction
	summaryLog     map[string]entity.Summary
}

func NewOHLCRecords(records *OHLC) *OHLC {
	newLogs := make(map[string][]entity.MstTransaction)
	newSummary := make(map[string]entity.Summary)
	return &OHLC{
		transactionLog: newLogs,
		summaryLog:     newSummary,
	}
}

func (rec OHLC) GetTransactionLog(stockCode string) []entity.MstTransaction {
	return rec.transactionLog[stockCode]
}

func (rec OHLC) GetSummaryLog(stockCode string) entity.Summary {
	return rec.summaryLog[stockCode]
}

func (rec OHLC) InsertNewRecord(trx entity.Transaction) (err error) {
	quantity, e := strconv.ParseInt(trx.Quantity, 10, 64)
	if e != nil {
		err = e
		return
	}

	price, e := strconv.ParseInt(trx.Price, 10, 64)
	if e != nil {
		err = e
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

	rec.CalculateRecordsByStockCode(newEntry)

	return
}

func (rec OHLC) CalculateRecordsByStockCode(trx entity.MstTransaction) {
	summary, found := rec.summaryLog[trx.Stock]
	if !found {
		summary.LowestPrice = math.MaxInt64
	}

	// switch {
	// case trx.Quantity == 0:
	// 	summary.PreviousPrice = trx.Price
	// 	summary.IsNewDay = true
	// case trx.Type == "E" || trx.Type == "P":
	// 	summary.Volume += trx.Quantity
	// 	summary.Value += trx.Quantity * trx.Price

	// }

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

}
