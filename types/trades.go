package types

// TradeStatus represents the status of a trade
type TradeStatus string

const (
	// TradeStatusMatched - Trade has been matched and sent to executor
	TradeStatusMatched TradeStatus = "MATCHED"
	// TradeStatusMined - Trade is mined into the chain
	TradeStatusMined TradeStatus = "MINED"
	// TradeStatusConfirmed - Trade has achieved finality and was successful
	TradeStatusConfirmed TradeStatus = "CONFIRMED"
	// TradeStatusRetrying - Trade transaction failed and is being retried
	TradeStatusRetrying TradeStatus = "RETRYING"
	// TradeStatusFailed - Trade has failed and is not being retried
	TradeStatusFailed TradeStatus = "FAILED"
)

// TradeType represents the type of trade (TAKER or MAKER)
type TradeType string

const (
	// TradeTypeTaker - Taker trade
	TradeTypeTaker TradeType = "TAKER"
	// TradeTypeMaker - Maker trade
	TradeTypeMaker TradeType = "MAKER"
)

// Trade represents a CLOB trade
type Trade struct {
	ID              string       `json:"id"`
	TakerOrderID    string       `json:"taker_order_id"`
	Market          string       `json:"market"`
	AssetID         string       `json:"asset_id"`
	Side            OrderSide    `json:"side"`
	Size            string       `json:"size"`
	FeeRateBps      string       `json:"fee_rate_bps"`
	Price           string       `json:"price"`
	Status          TradeStatus  `json:"status"`
	MatchTime       string       `json:"match_time"`
	LastUpdate      string       `json:"last_update"`
	Outcome         string       `json:"outcome"`
	MakerAddress    string       `json:"maker_address"`
	Owner           string       `json:"owner"`
	TransactionHash string       `json:"transaction_hash"`
	BucketIndex     int          `json:"bucket_index"`
	MakerOrders     []MakerOrder `json:"maker_orders"`
	Type            TradeType    `json:"type"`
}

// MakerOrder represents a maker order within a trade
type MakerOrder struct {
	OrderID       string    `json:"order_id"`
	MakerAddress  string    `json:"maker_address"`
	Owner         string    `json:"owner"`
	MatchedAmount string    `json:"matched_amount"`
	FeeRateBps    string    `json:"fee_rate_bps"`
	Price         string    `json:"price"`
	AssetID       string    `json:"asset_id"`
	Outcome       string    `json:"outcome"`
	Side          OrderSide `json:"side"`
}

// TradesRequest represents a request to get trades
type TradesRequest struct {
	ID     string `json:"id,omitempty"`
	Taker  string `json:"taker,omitempty"`
	Maker  string `json:"maker,omitempty"`
	Market string `json:"market,omitempty"`
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}
