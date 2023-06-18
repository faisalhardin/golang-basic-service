package calculation

import (
	"reflect"
	"task1/entity"
	"task1/src/repo"
	"testing"
)

var (
	mockTransaction = []entity.Transaction{
		{
			Type:     "A",
			Stock:    "BBCA",
			Quantity: "0",
			Price:    "8000",
		},
		{
			Type:     "P",
			Stock:    "BBCA",
			Quantity: "100",
			Price:    "8050",
		},
		{
			Type:     "P",
			Stock:    "BBCA",
			Quantity: "500",
			Price:    "7950",
		},
		{
			Type:     "A",
			Stock:    "BBCA",
			Quantity: "200",
			Price:    "8150",
		},
		{
			Type:     "E",
			Stock:    "BBCA",
			Quantity: "300",
			Price:    "8100",
		},
		{
			Type:     "A",
			Stock:    "BBCA",
			Quantity: "100",
			Price:    "8200",
		},
	}
)

func Test_InsertNewRecord(t *testing.T) {
	type args struct {
		records []entity.Transaction
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantSummary entity.Summary
	}{
		{
			name: "Successful and Correct InsertNewRecord",
			args: args{
				records: mockTransaction,
			},
			wantSummary: entity.Summary{
				PreviousPrice: 8000,
				OpenPrice:     8050,
				HighestPrice:  8100,
				LowestPrice:   7950,
				ClosePrice:    8100,
				Volume:        900,
				Value:         7210000,
				IsNewDay:      0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisRepo := repo.NewRedisRepo(&repo.RedisOptions{
				Address: "127.0.0.1:6379",
			})
			rec := NewOHLCRecords(&OHLC{
				Store: redisRepo,
			})

			for _, record := range tt.args.records {
				err := rec.InsertNewRecord(record)
				if (err != nil) != tt.wantErr {
					t.Errorf("InsertNewRecord() err = %v, wantErr = %v", err, tt.wantErr)
				}
			}
			got := rec.GetSummaryLog("BBCA")
			if !reflect.DeepEqual(tt.wantSummary, got) {
				t.Errorf("rec.GetSummaryLog() = %v, want %v", got, tt.wantSummary)
			}

		})
	}
}
