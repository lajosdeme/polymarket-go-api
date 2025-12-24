package types

// WebSocketEventType represents the type of WebSocket event
type WebSocketEventType string

const (
	// WSEventTypeBook - Orderbook update event
	WSEventTypeBook WebSocketEventType = "book"
	// WSEventTypePriceChange - Price change event
	WSEventTypePriceChange WebSocketEventType = "price_change"
	// WSEventTypeTickSizeChange - Tick size change event
	WSEventTypeTickSizeChange WebSocketEventType = "tick_size_change"
	// WSEventTypeLastTradePrice - Last trade price event
	WSEventTypeLastTradePrice WebSocketEventType = "last_trade_price"
	// WSEventTypeTrade - Trade event (user channel)
	WSEventTypeTrade WebSocketEventType = "trade"
	// WSEventTypeOrder - Order event (user channel)
	WSEventTypeOrder WebSocketEventType = "order"
)

// WebSocketChannelType represents the type of WebSocket channel
type WebSocketChannelType string

const (
	// WSChannelMarket - Market channel
	WSChannelMarket WebSocketChannelType = "market"
	// WSChannelUser - User channel
	WSChannelUser WebSocketChannelType = "user"
)

// WebSocketAuth represents WebSocket authentication
type WebSocketAuth struct {
	APIKey     string `json:"apiKey"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

// WebSocketSubscribeRequest represents a WebSocket subscription request
type WebSocketSubscribeRequest struct {
	Auth     *WebSocketAuth       `json:"auth,omitempty"`
	Markets  []string             `json:"markets,omitempty"`
	AssetIDs []string             `json:"assets_ids,omitempty"`
	Type     WebSocketChannelType `json:"type"`
}

// WebSocketSubscribeUpdate represents a subscription update request
type WebSocketSubscribeUpdate struct {
	AssetIDs  []string `json:"assets_ids"`
	Operation string   `json:"operation"` // "subscribe" or "unsubscribe"
}

// WebSocketBookEvent represents a book event from the market channel
type WebSocketBookEvent struct {
	EventType WebSocketEventType `json:"event_type"`
	AssetID   string             `json:"asset_id"`
	Market    string             `json:"market"`
	Timestamp string             `json:"timestamp"`
	Hash      string             `json:"hash"`
	Bids      []PriceLevel       `json:"bids"`
	Asks      []PriceLevel       `json:"asks"`
}

// WebSocketPriceChangeEvent represents a price change event
type WebSocketPriceChangeEvent struct {
	EventType    WebSocketEventType  `json:"event_type"`
	Market       string              `json:"market"`
	PriceChanges []PriceChangeDetail `json:"price_changes"`
	Timestamp    string              `json:"timestamp"`
}

// PriceChangeDetail represents a single price change detail
type PriceChangeDetail struct {
	AssetID string    `json:"asset_id"`
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Side    OrderSide `json:"side"`
	Hash    string    `json:"hash"`
	BestBid string    `json:"best_bid"`
	BestAsk string    `json:"best_ask"`
}

// WebSocketTickSizeChangeEvent represents a tick size change event
type WebSocketTickSizeChangeEvent struct {
	EventType   WebSocketEventType `json:"event_type"`
	AssetID     string             `json:"asset_id"`
	Market      string             `json:"market"`
	OldTickSize string             `json:"old_tick_size"`
	NewTickSize string             `json:"new_tick_size"`
	Side        OrderSide          `json:"side"`
	Timestamp   string             `json:"timestamp"`
}

// WebSocketLastTradePriceEvent represents a last trade price event
type WebSocketLastTradePriceEvent struct {
	EventType  WebSocketEventType `json:"event_type"`
	AssetID    string             `json:"asset_id"`
	FeeRateBps string             `json:"fee_rate_bps"`
	Market     string             `json:"market"`
	Price      string             `json:"price"`
	Side       OrderSide          `json:"side"`
	Size       string             `json:"size"`
	Timestamp  string             `json:"timestamp"`
}

// WebSocketTradeEvent represents a trade event from the user channel
type WebSocketTradeEvent struct {
	EventType    WebSocketEventType `json:"event_type"`
	AssetID      string             `json:"asset_id"`
	ID           string             `json:"id"`
	LastUpdate   string             `json:"last_update"`
	MakerOrders  []MakerOrder       `json:"maker_orders"`
	Market       string             `json:"market"`
	MatchTime    string             `json:"matchtime"`
	Outcome      string             `json:"outcome"`
	Owner        string             `json:"owner"`
	Price        string             `json:"price"`
	Side         OrderSide          `json:"side"`
	Size         string             `json:"size"`
	Status       TradeStatus        `json:"status"`
	TakerOrderID string             `json:"taker_order_id"`
	Timestamp    string             `json:"timestamp"`
	TradeOwner   string             `json:"trade_owner"`
	Type         string             `json:"type"`
}

// WebSocketOrderEvent represents an order event from the user channel
type WebSocketOrderEvent struct {
	EventType       WebSocketEventType `json:"event_type"`
	AssetID         string             `json:"asset_id"`
	AssociateTrades []string           `json:"associate_trades"`
	ID              string             `json:"id"`
	Market          string             `json:"market"`
	OrderOwner      string             `json:"order_owner"`
	OriginalSize    string             `json:"original_size"`
	Outcome         string             `json:"outcome"`
	Owner           string             `json:"owner"`
	Price           string             `json:"price"`
	Side            OrderSide          `json:"side"`
	SizeMatched     string             `json:"size_matched"`
	Timestamp       string             `json:"timestamp"`
	Type            string             `json:"type"`
}
