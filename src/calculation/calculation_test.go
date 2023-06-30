package calculation

import (
	"reflect"
	"task1/entity"
	"testing"

	mockrepo "task1/mock/repo"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/golang/mock/gomock"
)

var (
	mockInsertHisghestPriceTransaction = entity.Transaction{
		Type:             "P",
		Stock:            "BBCA",
		ExecutedQuantity: "100",
		ExecutedPrice:    "8050",
	}
	mockInsertLowestPriceTransaction = entity.Transaction{
		Type:             "P",
		Stock:            "BBCA",
		ExecutedQuantity: "100",
		ExecutedPrice:    "6000",
	}
	mockInsertOpenPriceTransaction = entity.Transaction{
		Type:     "E",
		Stock:    "BBCA",
		Quantity: "0",
		Price:    "6000",
	}
	mockGetNewDayPriceSummary = entity.Summary{
		PreviousPrice: 8000,
		HighestPrice:  8000,
		LowestPrice:   7000,
		Volume:        100,
		Value:         750000,
		IsNewDay:      1,
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
	mockKafkaMessage = &kafka.Message{}
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

func Test_NewOHLCRecords(t *testing.T) {
	mockOHLC := OHLC{
		transactionLog: make(map[string][]entity.MstTransaction),
		summaryLog:     make(map[string]entity.Summary),
	}
	type args struct {
		ohlc *OHLC
	}
	tests := []struct {
		name     string
		args     args
		wantOHLC OHLC
	}{
		{
			name: "Succesful",
			args: args{
				ohlc: &OHLC{},
			},
			wantOHLC: mockOHLC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ohlc := NewOHLCRecords(tt.args.ohlc)
			if !reflect.DeepEqual(*ohlc, mockOHLC) {
				t.Errorf("NewOHLCRecords() = %v, want = %v", *ohlc, tt.wantOHLC)
			}
		})
	}

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
			name: "Successful highest price InsertNewRecord",
			args: args{
				records: mockInsertHisghestPriceTransaction,
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
		{
			name: "Successful lowest price InsertNewRecord",
			args: args{
				records: mockInsertLowestPriceTransaction,
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
		{
			name: "Successful new day InsertNewRecord",
			args: args{
				records: mockInsertOpenPriceTransaction,
			},
			wantSummary: mockWantedSummary,
			wantErr:     false,
			patch: func() {
				mockStorage.EXPECT().HGetSummary(gomock.Any()).
					Return(mockGetNewDayPriceSummary, nil).Times(1)
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
			err := rec.InsertNewRecord(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertNewRecord() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func Test_InsertNewRecordFromKafka(t *testing.T) {
	ctrl := initMocks(t)
	defer ctrl.Finish()

	type args struct {
		records *kafka.Message
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantSummary *kafka.Message
		patch       func()
	}{
		{
			name: "Successful and Correct InsertNewRecord",
			args: args{
				records: mockKafkaMessage,
			},
			wantSummary: mockKafkaMessage,
			wantErr:     false,
			patch: func() {
				filereaderConvertToStruct = func(line []byte) (transaction entity.Transaction, err error) {
					return entity.Transaction{}, nil
				}

				ohlcInsertNewRecord = func(trx entity.Transaction) (err error) {
					return nil
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &OHLC{}
			tt.patch()
			err := rec.InsertNewRecordFromKafka(mockKafkaMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertNewRecordFromKafka() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
