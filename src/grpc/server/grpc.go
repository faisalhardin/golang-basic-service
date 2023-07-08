package server

import (
	"context"
	summary_proto "task1/entity/summaryservice"
	"task1/src/calculation"

	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

type GRPCServiceHandler struct {
	OHLCUsecase calculation.OHLC
}

func NewGRPCServiceHandler(handler *GRPCServiceHandler) *GRPCServiceHandler {
	return handler
}

func (s GRPCServiceHandler) GetSummary(ctx context.Context, summary *summary_proto.GetStockSummaryRequest) (resp *summary_proto.GetStockSummaryResponse, err error) {

	ohlcSummary, err := s.OHLCUsecase.GetRedisSummaryLog(summary.Stock)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		err = errors.Wrap(err, "GRPCServiceHandler.GetSummary")
		return
	}
	err = nil

	resp = &summary_proto.GetStockSummaryResponse{}
	var average int64
	if ohlcSummary.Volume > 0 {
		average = ohlcSummary.Value / ohlcSummary.Volume
	}
	resp.Summary = &summary_proto.Summary{
		Stock:         summary.Stock,
		Previousprice: ohlcSummary.PreviousPrice,
		Openprice:     ohlcSummary.OpenPrice,
		Highestprice:  ohlcSummary.HighestPrice,
		Lowestprice:   ohlcSummary.LowestPrice,
		Closeprice:    ohlcSummary.ClosePrice,
		Value:         ohlcSummary.Value,
		Volume:        ohlcSummary.Volume,
		Average:       average,
	}

	return
}
