package entity

import (
	"fmt"
	"strconv"
)

var (
	ErrNotEvent = fmt.Errorf("parameter needs to be in event number consistiong of pair of key and value")
)

const (
	IsNewDayTrue  int64 = 1
	IsNewDayFalse int64 = 0
)

type Transaction struct {
	Type             string `json:"type"`
	Stock            string `json:"stock_code"`
	Quantity         string `json:"quantity"`
	ExecutedQuantity string `json:"executed_quantity"`
	Price            string `json:"price"`
	ExecutedPrice    string `json:"execution_price"`
}

type MstTransaction struct {
	Type             string `json:"type"`
	Stock            string `json:"stock_code"`
	Quantity         int64  `json:"quantity"`
	ExecutedQuantity int64  `json:"executed_quantity"`
	Price            int64  `json:"price"`
	ExecutedPrice    int64  `json:"execution_price"`
}

type Summary struct {
	PreviousPrice int64 `json:"previous_price"`
	OpenPrice     int64 `json:"open_price"`
	HighestPrice  int64 `json:"highest_price"`
	LowestPrice   int64 `json:"lowest_price"`
	ClosePrice    int64 `json:"close_price"`
	Volume        int64 `json:"volume"`
	Value         int64 `json:"value"`
	IsNewDay      int64 `json:"is_new_day"`
}

func (s *Summary) IsCurrentlyNewDay() bool {
	return s.IsNewDay == IsNewDayTrue
}

func (s *Summary) ConvertFromHGetAllToStruct(redisResponse []string) (err error) {

	if len(redisResponse)%2 > 0 {
		return ErrNotEvent
	}
	for i := 0; i < len(redisResponse)/2; i++ {
		key := redisResponse[i*2]
		value, e := strconv.ParseInt(redisResponse[i*2+1], 10, 64)
		if e != nil {
			err = e
			return
		}
		switch key {
		case "previous_price":
			s.PreviousPrice = value
		case "open_price":
			s.OpenPrice = value
		case "highest_price":
			s.HighestPrice = value
		case "lowest_price":
			s.LowestPrice = value
		case "close_price":
			s.ClosePrice = value
		case "volume":
			s.Volume = value
		case "value":
			s.Value = value
		case "is_new_day":
			s.IsNewDay = value
		}
	}

	return
}

type StockCodeToTransactionLogKeyValue struct {
	StockCode      string
	TransactionLog string
}
