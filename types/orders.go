package types

// OrderSide represents the side of an order (BUY or SELL)
type OrderSide string

const (
	// BUY - Buy order
	BUY OrderSide = "BUY"
	// SELL - Sell order
	SELL OrderSide = "SELL"
)

// OrderType represents the type of order
type OrderType string

const (
	// FOK - Fill-Or-Kill order
	FOK OrderType = "FOK"
	// FAK - Fill-And-Kill order
	FAK OrderType = "FAK"
	// GTC - Good-Til-Cancelled order
	GTC OrderType = "GTC"
	// GTD - Good-Til-Date order
	GTD OrderType = "GTD"
)

// OrderStatus represents the current status of an order
type OrderStatus string

const (
	// OrderStatusLive - Order is live on the book
	OrderStatusLive OrderStatus = "live"
	// OrderStatusMatched - Order has been matched
	OrderStatusMatched OrderStatus = "matched"
	// OrderStatusDelayed - Order is delayed
	OrderStatusDelayed OrderStatus = "delayed"
	// OrderStatusUnmatched - Order was not matched
	OrderStatusUnmatched OrderStatus = "unmatched"
)

// Order represents a CLOB order
type Order struct {
	Salt          string    `json:"salt"`
	Maker         string    `json:"maker"`
	Signer        string    `json:"signer"`
	Taker         string    `json:"taker"`
	TokenID       string    `json:"tokenId"`
	MakerAmount   string    `json:"makerAmount"`
	TakerAmount   string    `json:"takerAmount"`
	Expiration    string    `json:"expiration"`
	Nonce         string    `json:"nonce"`
	FeeRateBps    string    `json:"feeRateBps"`
	Side          OrderSide `json:"side"`
	SignatureType int       `json:"signatureType"`
	Signature     string    `json:"signature"`
}

// PostOrder represents an order with additional metadata for posting
type PostOrder struct {
	Order     Order     `json:"order"`
	OrderType OrderType `json:"orderType"`
	Owner     string    `json:"owner"`
}

// OrderResponse represents the response from placing an order
type OrderResponse struct {
	Success     bool     `json:"success"`
	ErrorMsg    string   `json:"errorMsg"`
	OrderID     string   `json:"orderId"`
	OrderHashes []string `json:"orderHashes"`
	Status      string   `json:"status,omitempty"`
}

// OpenOrder represents an open order on the book
type OpenOrder struct {
	AssociateTrades []string  `json:"associate_trades"`
	ID              string    `json:"id"`
	Status          string    `json:"status"`
	Market          string    `json:"market"`
	OriginalSize    string    `json:"original_size"`
	Outcome         string    `json:"outcome"`
	MakerAddress    string    `json:"maker_address"`
	Owner           string    `json:"owner"`
	Price           string    `json:"price"`
	Side            OrderSide `json:"side"`
	SizeMatched     string    `json:"size_matched"`
	AssetID         string    `json:"asset_id"`
	Expiration      string    `json:"expiration"`
	Type            string    `json:"type"`
	CreatedAt       string    `json:"created_at"`
}

// CancelOrderRequest represents a request to cancel an order
type CancelOrderRequest struct {
	OrderID string `json:"orderID"`
}

// CancelOrdersRequest represents a request to cancel multiple orders
type CancelOrdersRequest struct {
	OrderIDs []string `json:"orderIDs"`
}

// CancelMarketOrdersRequest represents a request to cancel orders from a market
type CancelMarketOrdersRequest struct {
	Market  string `json:"market,omitempty"`
	AssetID string `json:"asset_id,omitempty"`
}

// CancelResponse represents the response from canceling orders
type CancelResponse struct {
	Canceled    []string          `json:"canceled"`
	NotCanceled map[string]string `json:"not_canceled"`
}

// OrderScoring represents the scoring information for an order
type OrderScoring struct {
	Scoring bool `json:"scoring"`
}

// OrdersScoringRequest represents a request to check scoring for multiple orders
type OrdersScoringRequest struct {
	OrderIDs []string `json:"orderIds"`
}

// OrdersScoring represents the scoring information for multiple orders
type OrdersScoring map[string]bool
