package calculation

import (
	"task1/entity"
	"testing"

	mockrepo "task1/mock/repo"

	"github.com/golang/mock/gomock"
)

var (
	mockInsertTransaction = entity.Transaction{
		Type:             "P",
		Stock:            "BBCA",
		ExecutedQuantity: "100",
		ExecutedPrice:    "8050",
	}
	mockGetSummary = entity.Summary{
		PreviousPrice: 8000,
		HighestPrice:  8000,
		LowestPrice:   7000,
		Volume:        100,
		Value:         750000,
	}
	mockWantedSummary = entity.Summary{
		PreviousPrice: 8000,
		HighestPrice:  8050,
		LowestPrice:   7000,
		Volume:        200,
		Value:         1555000,
	}
)

var (
	mockRedisHandler *mockrepo.MockHandler
	mockStorage      *mockrepo.MockStorageInterface
)

func initMocks(t *testing.T) *gomock.Controller {
	ctrl := gomock.NewController(t)
	mockRedisHandler = mockrepo.NewMockHandler(ctrl)
	mockStorage = mockrepo.NewMockStorageInterface(ctrl)

	return ctrl
}

func Test_InsertNewRecord(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		records entity.Transaction
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantSummary entity.Summary
		patch       func()
	}{
		{
			name: "Successful and Correct InsertNewRecord",
			args: args{
				records: mockInsertTransaction,
			},
			wantSummary: mockWantedSummary,
			wantErr:     false,
			patch: func() {
				mockStorage.EXPECT().HGetSummary(gomock.Any()).
					Return(mockGetSummary, nil).Times(1)
				mockStorage.EXPECT().Del(gomock.Any()).Return(int64(1), nil).Times(1)
				mockStorage.EXPECT().HSetSummary(gomock.Any(), gomock.Any()).
					Return(int64(1), nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &OHLC{
				Store: mockStorage,
			}
			tt.patch()
			err := rec.InsertNewRecord(mockInsertTransaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertNewRecord() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
