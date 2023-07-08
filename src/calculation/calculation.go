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

	WrapMsgCalculateRecordsByStockCode = "CalculateRecordsByStockCode"
)

type OHLC struct {
	Store entity.StorageInterface
}

func NewOHLCRecords(records *OHLC) *OHLC {

	records = &OHLC{
		Store: records.Store,
	}

	filereaderConvertToStruct = filereader.ConvertToStruct
	ohlcInsertNewRecord = records.InsertNewRecord

	return records
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
	var summaryOperation SummaryOperation
	summary, err := rec.GetRedisSummaryLog(trx.Stock)
	if err == nil {
		summaryOperation = &StoredSummary{
			stockCode: trx.Stock,
			summary:   summary,
		}
	}
	if err != nil && !errors.Is(err, redis.ErrNil) {
		err = errors.Wrap(err, WrapMsgCalculateRecordsByStockCode)
		return err
	}
	if err != nil && errors.Is(err, redis.ErrNil) {
		summary.LowestPrice = math.MaxInt32
		summaryOperation = &UnstoredSummary{
			stockCode: trx.Stock,
			summary:   summary,
		}
		err = nil
	}

	if trx.Type == "E" || trx.Type == "P" {
		err = summaryOperation.SetSummaryVolume(rec, summary.Volume+trx.ExecutedQuantity)
		if err != nil {
			return errors.Wrap(err, "CalculateRecordsByStockCode")
		}
		err = summaryOperation.SetSummaryValue(rec, summary.Value+(trx.ExecutedQuantity*trx.ExecutedPrice))
		if err != nil {
			return errors.Wrap(err, "CalculateRecordsByStockCode")
		}

		if summary.HighestPrice < trx.ExecutedPrice {
			err = summaryOperation.SetSummaryHighesPrice(rec, trx.ExecutedPrice)
			if err != nil {
				return errors.Wrap(err, "CalculateRecordsByStockCode")
			}
		}

		if summary.LowestPrice > trx.ExecutedPrice {
			err = summaryOperation.SetSummaryLowestPrice(rec, trx.ExecutedPrice)
			if err != nil {
				return errors.Wrap(err, "CalculateRecordsByStockCode")
			}
		}

		if summary.IsCurrentlyNewDay() {
			err = summaryOperation.SetSummaryOpenPrice(rec, trx.ExecutedPrice)
			if err != nil {
				return errors.Wrap(err, "CalculateRecordsByStockCode")
			}
			err = summaryOperation.SetSummaryIsNewDay(rec, entity.IsNewDayFalse)
			if err != nil {
				return errors.Wrap(err, "CalculateRecordsByStockCode")
			}
		}

		err = summaryOperation.SetSummaryClosePrice(rec, trx.ExecutedPrice)
		if err != nil {
			return errors.Wrap(err, "CalculateRecordsByStockCode")
		}
	}
	if (trx.Type == "E" || trx.Type == "P") && trx.ExecutedQuantity == 0 ||
		(trx.Type != "E" && trx.Type != "P") && trx.Quantity == 0 {
		err = summaryOperation.SetSummaryPreviousPrice(rec, trx.Price)
		if err != nil {
			return errors.Wrap(err, "CalculateRecordsByStockCode")
		}

		err = summaryOperation.SetSummaryIsNewDay(rec, entity.IsNewDayTrue)
		if err != nil {
			return errors.Wrap(err, "CalculateRecordsByStockCode")
		}
	}

	log.Default().Print(summary, " stock: ", trx.Stock, " type: ", trx.Type)
	err = summaryOperation.SetSummaryWhole(rec)
	if err != nil {
		return errors.Wrap(err, "CalculateRecordsByStockCode")
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

type StoredSummary struct {
	stockCode string
	summary   entity.Summary
}

func (s *StoredSummary) GetSummary() entity.Summary {
	return s.summary
}

func (s *StoredSummary) SetSummaryVolume(ohlc OHLC, value int64) (err error) {
	s.summary.Volume = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "volume", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryVolume")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryValue(ohlc OHLC, value int64) (err error) {
	s.summary.Value = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "value", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryValue")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryHighesPrice(ohlc OHLC, value int64) (err error) {
	s.summary.HighestPrice = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "highest_price", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryHighesPrice")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryLowestPrice(ohlc OHLC, value int64) (err error) {
	s.summary.LowestPrice = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "lowest_price", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryLowestPrice")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryOpenPrice(ohlc OHLC, value int64) (err error) {
	s.summary.OpenPrice = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "open_price", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryOpenPrice")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryIsNewDay(ohlc OHLC, value int64) (err error) {
	s.summary.IsNewDay = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "is_new_day", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryIsNewDay")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryClosePrice(ohlc OHLC, value int64) (err error) {
	s.summary.ClosePrice = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "close_price", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryClosePrice")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryPreviousPrice(ohlc OHLC, value int64) (err error) {
	s.summary.PreviousPrice = value
	err = ohlc.SetRedisSummaryField(s.stockCode, "previous_price", value)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryPreviousPrice")
		return err
	}

	return nil
}

func (s *StoredSummary) SetSummaryWhole(ohlc OHLC) (err error) {
	return nil
}

type UnstoredSummary struct {
	stockCode string
	summary   entity.Summary
}

func (s *UnstoredSummary) GetSummary() entity.Summary {
	return s.summary
}

func (s *UnstoredSummary) SetSummaryVolume(ohlc OHLC, value int64) (err error) {
	s.summary.Volume = value

	return nil
}

func (s *UnstoredSummary) SetSummaryValue(ohlc OHLC, value int64) (err error) {
	s.summary.Value = value

	return nil
}

func (s *UnstoredSummary) SetSummaryHighesPrice(ohlc OHLC, value int64) (err error) {
	s.summary.HighestPrice = value

	return nil
}

func (s *UnstoredSummary) SetSummaryLowestPrice(ohlc OHLC, value int64) (err error) {
	s.summary.LowestPrice = value

	return nil
}

func (s *UnstoredSummary) SetSummaryOpenPrice(ohlc OHLC, value int64) (err error) {
	s.summary.OpenPrice = value

	return nil
}

func (s *UnstoredSummary) SetSummaryIsNewDay(ohlc OHLC, value int64) (err error) {
	s.summary.IsNewDay = value

	return nil
}

func (s *UnstoredSummary) SetSummaryClosePrice(ohlc OHLC, value int64) (err error) {
	s.summary.ClosePrice = value

	return nil
}

func (s *UnstoredSummary) SetSummaryPreviousPrice(ohlc OHLC, value int64) (err error) {
	s.summary.PreviousPrice = value

	return nil
}

func (s *UnstoredSummary) SetSummaryWhole(ohlc OHLC) (err error) {
	err = ohlc.SetRedisSummaryLog(s.stockCode, s.summary)
	if err != nil {
		err = errors.Wrap(err, "SetSummaryWhole")
		return err
	}

	return nil
}

type SummaryOperation interface {
	GetSummary() entity.Summary
	SetSummaryVolume(ohlc OHLC, value int64) (err error)
	SetSummaryValue(ohlc OHLC, value int64) (err error)
	SetSummaryHighesPrice(ohlc OHLC, value int64) (err error)
	SetSummaryLowestPrice(ohlc OHLC, value int64) (err error)
	SetSummaryOpenPrice(ohlc OHLC, value int64) (err error)
	SetSummaryIsNewDay(ohlc OHLC, value int64) (err error)
	SetSummaryClosePrice(ohlc OHLC, value int64) (err error)
	SetSummaryPreviousPrice(ohlc OHLC, value int64) (err error)
	SetSummaryWhole(ohlc OHLC) (err error)
}
