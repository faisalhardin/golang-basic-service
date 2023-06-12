package entity

type Transaction struct {
	Type      string `json:"type"`
	Stock     string `json:"stock_code"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	OrderBook string `json:"order_book"`
}

type MstTransaction struct {
	Type      string `json:"type"`
	Stock     string `json:"stock_code"`
	Quantity  int64  `json:"quantity"`
	Price     int64  `json:"price"`
	OrderBook string `json:"order_book"`
}

type Summary struct {
	PreviousPrice int64 `json:"previous_price"`
	OpenPrice     int64 `json:"open_price"`
	HighestPrice  int64 `json:"highest_price"`
	LowestPrice   int64 `json:"lowest_price"`
	ClosePrice    int64 `json:"close_price"`
	Volume        int64 `json:"volume"`
	Value         int64 `json:"value"`
	IsNewDay      bool
}
