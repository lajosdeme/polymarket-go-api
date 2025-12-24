package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// WebSocketClient handles WebSocket connections to Polymarket
type WebSocketClient struct {
	baseURL      string
	conn         *websocket.Conn
	authManager  *AuthManager
	pingInterval time.Duration
	pingTicker   *time.Ticker
	writeMutex   sync.Mutex
	stopChan     chan struct{}

	// Event handlers
	onBookMessage           func(*types.WebSocketBookEvent)
	onPriceChangeMessage    func(*types.WebSocketPriceChangeEvent)
	onTickSizeChangeMessage func(*types.WebSocketTickSizeChangeEvent)
	onLastTradePriceMessage func(*types.WebSocketLastTradePriceEvent)
	onTradeMessage          func(*types.WebSocketTradeEvent)
	onOrderMessage          func(*types.WebSocketOrderEvent)
	onError                 func(error)
	onClose                 func()
}

// NewWebSocketClient creates a new WebSocket client
func NewWebSocketClient(baseURL string, authManager *AuthManager) *WebSocketClient {
	if baseURL == "" {
		baseURL = "wss://ws-subscriptions-clob.polymarket.com"
	}

	return &WebSocketClient{
		baseURL:      baseURL,
		authManager:  authManager,
		pingInterval: 10 * time.Second,
		stopChan:     make(chan struct{}),
	}
}

// SetPingInterval sets the WebSocket ping interval
func (w *WebSocketClient) SetPingInterval(interval time.Duration) {
	w.pingInterval = interval
}

// SetBookMessageHandler sets handler for book events
func (w *WebSocketClient) SetBookMessageHandler(handler func(*types.WebSocketBookEvent)) {
	w.onBookMessage = handler
}

// SetPriceChangeMessageHandler sets handler for price change events
func (w *WebSocketClient) SetPriceChangeMessageHandler(handler func(*types.WebSocketPriceChangeEvent)) {
	w.onPriceChangeMessage = handler
}

// SetTickSizeChangeMessageHandler sets handler for tick size change events
func (w *WebSocketClient) SetTickSizeChangeMessageHandler(handler func(*types.WebSocketTickSizeChangeEvent)) {
	w.onTickSizeChangeMessage = handler
}

// SetLastTradePriceMessageHandler sets handler for last trade price events
func (w *WebSocketClient) SetLastTradePriceMessageHandler(handler func(*types.WebSocketLastTradePriceEvent)) {
	w.onLastTradePriceMessage = handler
}

// SetTradeMessageHandler sets handler for trade events (user channel)
func (w *WebSocketClient) SetTradeMessageHandler(handler func(*types.WebSocketTradeEvent)) {
	w.onTradeMessage = handler
}

// SetOrderMessageHandler sets handler for order events (user channel)
func (w *WebSocketClient) SetOrderMessageHandler(handler func(*types.WebSocketOrderEvent)) {
	w.onOrderMessage = handler
}

// SetErrorHandler sets handler for connection errors
func (w *WebSocketClient) SetErrorHandler(handler func(error)) {
	w.onError = handler
}

// SetCloseHandler sets handler for connection close
func (w *WebSocketClient) SetCloseHandler(handler func()) {
	w.onClose = handler
}

// Connect connects to the market channel
func (w *WebSocketClient) ConnectMarketChannel(assetIDs []string) error {
	u, err := url.Parse(w.baseURL + "/ws/market")
	if err != nil {
		return fmt.Errorf("failed to parse WebSocket URL: %w", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	w.conn = conn

	// Send subscription message
	subscribeMsg := types.WebSocketSubscribeRequest{
		AssetIDs: assetIDs,
		Type:     types.WSChannelMarket,
	}

	if err := w.writeMessage(subscribeMsg); err != nil {
		conn.Close()
		return fmt.Errorf("failed to send subscription message: %w", err)
	}

	// Start message handler
	go w.messageHandler()

	// Start ping routine
	go w.startPing()

	return nil
}

// Connect connects to the user channel
func (w *WebSocketClient) ConnectUserChannel(markets []string) error {
	// Validate L2 authentication
	if !w.authManager.HasL2Auth() {
		return fmt.Errorf("L2 authentication required for user channel")
	}

	creds := w.authManager.GetAPICredentials()
	if creds == nil {
		return fmt.Errorf("API credentials not found")
	}

	u, err := url.Parse(w.baseURL + "/ws/user")
	if err != nil {
		return fmt.Errorf("failed to parse WebSocket URL: %w", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	w.conn = conn

	// Send subscription message with authentication
	subscribeMsg := types.WebSocketSubscribeRequest{
		Auth: &types.WebSocketAuth{
			APIKey:     creds.APIKey,
			Secret:     creds.Secret,
			Passphrase: creds.Passphrase,
		},
		Markets: markets,
		Type:    types.WSChannelUser,
	}

	if err := w.writeMessage(subscribeMsg); err != nil {
		conn.Close()
		return fmt.Errorf("failed to send subscription message: %w", err)
	}

	// Start message handler
	go w.messageHandler()

	// Start ping routine
	go w.startPing()

	return nil
}

// SubscribeToAssets subscribes to additional asset IDs (market channel only)
func (w *WebSocketClient) SubscribeToAssets(assetIDs []string) error {
	if w.conn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}

	updateMsg := types.WebSocketSubscribeUpdate{
		AssetIDs:  assetIDs,
		Operation: "subscribe",
	}

	return w.writeMessage(updateMsg)
}

// UnsubscribeFromAssets unsubscribes from asset IDs (market channel only)
func (w *WebSocketClient) UnsubscribeFromAssets(assetIDs []string) error {
	if w.conn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}

	updateMsg := types.WebSocketSubscribeUpdate{
		AssetIDs:  assetIDs,
		Operation: "unsubscribe",
	}

	return w.writeMessage(updateMsg)
}

// Close closes the WebSocket connection
func (w *WebSocketClient) Close() error {
	close(w.stopChan)

	if w.pingTicker != nil {
		w.pingTicker.Stop()
	}

	if w.conn != nil {
		return w.conn.Close()
	}

	return nil
}

// writeMessage writes a message to the WebSocket connection
func (w *WebSocketClient) writeMessage(msg interface{}) error {
	w.writeMutex.Lock()
	defer w.writeMutex.Unlock()

	if w.conn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	w.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err := w.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

// startPing starts the ping routine
func (w *WebSocketClient) startPing() {
	w.pingTicker = time.NewTicker(w.pingInterval)
	defer w.pingTicker.Stop()

	for {
		select {
		case <-w.stopChan:
			return
		case <-w.pingTicker.C:
			if w.conn != nil {
				w.writeMutex.Lock()
				w.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
				w.conn.WriteMessage(websocket.TextMessage, []byte("PING"))
				w.writeMutex.Unlock()
			}
		}
	}
}

// messageHandler handles incoming WebSocket messages
func (w *WebSocketClient) messageHandler() {
	defer func() {
		if w.onClose != nil {
			w.onClose()
		}
	}()

	for {
		select {
		case <-w.stopChan:
			return
		default:
			if w.conn == nil {
				return
			}

			_, message, err := w.conn.ReadMessage()
			if err != nil {
				if w.onError != nil {
					w.onError(fmt.Errorf("WebSocket read error: %w", err))
				}
				return
			}

			// Handle ping messages
			if string(message) == "PONG" {
				continue
			}

			w.handleMessage(message)
		}
	}
}

// handleMessage processes incoming WebSocket messages
func (w *WebSocketClient) handleMessage(message []byte) {
	var baseMsg map[string]interface{}
	if err := json.Unmarshal(message, &baseMsg); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal message: %w", err))
		}
		return
	}

	eventType, ok := baseMsg["event_type"].(string)
	if !ok {
		if w.onError != nil {
			w.onError(fmt.Errorf("message missing event_type"))
		}
		return
	}

	switch types.WebSocketEventType(eventType) {
	case types.WSEventTypeBook:
		w.handleBookMessage(message)
	case types.WSEventTypePriceChange:
		w.handlePriceChangeMessage(message)
	case types.WSEventTypeTickSizeChange:
		w.handleTickSizeChangeMessage(message)
	case types.WSEventTypeLastTradePrice:
		w.handleLastTradePriceMessage(message)
	case types.WSEventTypeTrade:
		w.handleTradeMessage(message)
	case types.WSEventTypeOrder:
		w.handleOrderMessage(message)
	default:
		if w.onError != nil {
			w.onError(fmt.Errorf("unknown event type: %s", eventType))
		}
	}
}

func (w *WebSocketClient) handleBookMessage(message []byte) {
	if w.onBookMessage == nil {
		return
	}

	var event types.WebSocketBookEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal book event: %w", err))
		}
		return
	}

	w.onBookMessage(&event)
}

func (w *WebSocketClient) handlePriceChangeMessage(message []byte) {
	if w.onPriceChangeMessage == nil {
		return
	}

	var event types.WebSocketPriceChangeEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal price change event: %w", err))
		}
		return
	}

	w.onPriceChangeMessage(&event)
}

func (w *WebSocketClient) handleTickSizeChangeMessage(message []byte) {
	if w.onTickSizeChangeMessage == nil {
		return
	}

	var event types.WebSocketTickSizeChangeEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal tick size change event: %w", err))
		}
		return
	}

	w.onTickSizeChangeMessage(&event)
}

func (w *WebSocketClient) handleLastTradePriceMessage(message []byte) {
	if w.onLastTradePriceMessage == nil {
		return
	}

	var event types.WebSocketLastTradePriceEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal last trade price event: %w", err))
		}
		return
	}

	w.onLastTradePriceMessage(&event)
}

func (w *WebSocketClient) handleTradeMessage(message []byte) {
	if w.onTradeMessage == nil {
		return
	}

	var event types.WebSocketTradeEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal trade event: %w", err))
		}
		return
	}

	w.onTradeMessage(&event)
}

func (w *WebSocketClient) handleOrderMessage(message []byte) {
	if w.onOrderMessage == nil {
		return
	}

	var event types.WebSocketOrderEvent
	if err := json.Unmarshal(message, &event); err != nil {
		if w.onError != nil {
			w.onError(fmt.Errorf("failed to unmarshal order event: %w", err))
		}
		return
	}

	w.onOrderMessage(&event)
}
