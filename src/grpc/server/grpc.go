package server

import (
	"context"
	"log"
	summary_proto "task1/entity/summaryservice"
)

type Server struct {
}

func (s Server) GetSummary(ctx context.Context, summary *summary_proto.GetStockSummaryRequest) (resp *summary_proto.GetStockSummaryResponse, err error) {
	log.Printf("Received %s", summary.Stock)
	resp = &summary_proto.GetStockSummaryResponse{}
	resp.Summary = &summary_proto.Summary{
		Stock: summary.Stock,
		Value: 69,
	}
	return
}
