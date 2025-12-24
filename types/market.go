package types

// PriceLevel represents a price level in the orderbook
type PriceLevel struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

// Orderbook represents an orderbook for a token
type Orderbook struct {
	Market       string       `json:"market"`
	AssetID      string       `json:"asset_id"`
	Timestamp    string       `json:"timestamp"`
	Hash         string       `json:"hash"`
	Bids         []PriceLevel `json:"bids"`
	Asks         []PriceLevel `json:"asks"`
	MinOrderSize string       `json:"min_order_size"`
	TickSize     string       `json:"tick_size"`
	NegRisk      bool         `json:"neg_risk"`
}

// OrderbookRequest represents a request to get orderbook data
type OrderbookRequest struct {
	TokenID string `json:"token_id"`
}

// OrderbooksRequest represents a request to get multiple orderbooks
type OrderbooksRequest struct {
	TokenIDs []string `json:"token_ids"`
}

// Price represents a price for a token and side
type Price struct {
	Price string `json:"price"`
}

// PricesRequest represents a request to get prices for multiple tokens
type PricesRequest struct {
	TokenID string    `json:"token_id"`
	Side    OrderSide `json:"side"`
}

// PricesResponse represents the response from prices endpoint
type PricesResponse map[string]map[string]string

// MidpointResponse represents the response from midpoint endpoint
type MidpointResponse struct {
	Mid string `json:"mid"`
}

// PriceHistoryPoint represents a single point in price history
type PriceHistoryPoint struct {
	T int64   `json:"t"` // timestamp
	P float64 `json:"p"` // price
}

// PriceHistoryResponse represents the response from prices-history endpoint
type PriceHistoryResponse struct {
	History []PriceHistoryPoint `json:"history"`
}

// PriceHistoryRequest represents a request to get price history
type PriceHistoryRequest struct {
	Market   string `json:"market"`
	StartTs  *int64 `json:"startTs,omitempty"`
	EndTs    *int64 `json:"endTs,omitempty"`
	Interval string `json:"interval,omitempty"`
	Fidelity *int   `json:"fidelity,omitempty"`
}

// SpreadsRequest represents a request to get spreads
type SpreadsRequest struct {
	TokenID string    `json:"token_id"`
	Side    OrderSide `json:"side,omitempty"`
}

// SpreadsResponse represents the response from spreads endpoint
type SpreadsResponse map[string]string
