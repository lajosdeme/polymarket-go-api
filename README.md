# Polymarket Go API SDK

A comprehensive Go SDK for interacting with the Polymarket CLOB (Centralized Limit Order Book) REST API and WebSocket endpoints.

## Features

- **Full API Coverage**: Complete implementation of all documented Polymarket CLOB API endpoints
- **Authentication**: Support for both L1 (private key/EIP-712) and L2 (API key/HMAC) authentication
- **Order Management**: Place, cancel, and monitor orders with all supported order types
- **Market Data**: Access real-time pricing, orderbook, and historical data
- **WebSocket Integration**: Real-time market and user channel subscriptions
- **Type Safety**: Comprehensive type definitions for all API entities
- **Error Handling**: Detailed error types with retry logic support
- **Thread Safety**: Safe concurrent operations

## Installation

```bash
go get github.com/lajosdeme/polymarket-go-api
```

## Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/lajosdeme/polymarket-go-api/api"
    "github.com/lajosdeme/polymarket-go-api/client"
    "github.com/lajosdeme/polymarket-go-api/types"
)

func main() {
    // Initialize client
    c := client.NewClobClient("")
    
    // Setup L2 authentication with API credentials
    err := c.SetupL2Auth("your-api-key", "your-secret", "your-passphrase")
    if err != nil {
        log.Fatal(err)
    }

    // Create API instances
    pricingAPI := api.NewPricingAPI(c)
    ordersAPI := api.NewOrdersAPI(c)
    
    // Get market price
    price, err := pricingAPI.GetPrice(context.Background(), "token-id", types.BUY)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Current BUY price: %s\n", price.Price)
}
```

### L1 Authentication (Private Key)

```go
// Setup L1 authentication for API key creation
err := c.SetupL1Auth("your-private-key", types.EOA, "funder-address")
if err != nil {
    log.Fatal(err)
}

// Create API credentials
authAPI := api.NewAuthAPI(c)
credentials, err := authAPI.CreateAPIKey(context.Background(), nonce)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("API Key: %s\n", credentials.APIKey)
```

### WebSocket Connection

```go
// Create WebSocket client
wsClient := client.NewWebSocketClient("", c.GetAuthManager())

// Set up event handlers
wsClient.SetBookMessageHandler(func(event *types.WebSocketBookEvent) {
    fmt.Printf("Book update: %s\n", event.AssetID)
})

wsClient.SetPriceChangeMessageHandler(func(event *types.WebSocketPriceChangeEvent) {
    fmt.Printf("Price change: %s\n", event.Market)
})

// Connect to market channel
err := wsClient.ConnectMarketChannel([]string{"token-id-1", "token-id-2"})
if err != nil {
    log.Fatal(err)
}

// Keep connection alive
select {}
```

## Authentication

### L1 Authentication (Private Key)

Used for creating API credentials and signing orders locally.

```go
// EOA (MetaMask) wallet
err := c.SetupL1Auth(privateKey, types.EOA, "")

// Polymarket Proxy (Magic Link) wallet
err := c.SetupL1Auth(privateKey, types.POLY_PROXY, proxyAddress)

// Gnosis Safe wallet
err := c.SetupL1Auth(privateKey, types.GNOSIS_SAFE, safeAddress)
```

### L2 Authentication (API Credentials)

Used for trading operations and accessing user data.

```go
err := c.SetupL2Auth(apiKey, secret, passphrase)
```

## API Endpoints

### Authentication API

```go
authAPI := api.NewAuthAPI(c)

// Create new API credentials
credentials, err := authAPI.CreateAPIKey(ctx, nonce)

// Derive existing credentials
credentials, err := authAPI.DeriveAPIKey(ctx, nonce)
```

### Orders API

```go
ordersAPI := api.NewOrdersAPI(c)

// Place single order
response, err := ordersAPI.PlaceOrder(ctx, order)

// Place multiple orders
responses, err := ordersAPI.PlaceOrders(ctx, orders)

// Get order information
order, err := ordersAPI.GetOrder(ctx, orderID)

// Get active orders
orders, err := ordersAPI.GetActiveOrders(ctx, market, assetID)

// Cancel orders
response, err := ordersAPI.CancelOrder(ctx, orderID)
responses, err := ordersAPI.CancelOrders(ctx, orderIDs)
response, err := ordersAPI.CancelAllOrders(ctx)

// Check order scoring
scoring, err := ordersAPI.CheckOrderScoring(ctx, orderID)
```

### Pricing API

```go
pricingAPI := api.NewPricingAPI(c)

// Get market price
price, err := pricingAPI.GetPrice(ctx, tokenID, types.BUY)

// Get multiple prices
prices, err := pricingAPI.GetPrices(ctx)
prices, err := pricingAPI.GetPricesByRequest(ctx, requests)

// Get midpoint price
midpoint, err := pricingAPI.GetMidpointPrice(ctx, tokenID)

// Get price history
history, err := pricingAPI.GetPriceHistory(ctx, request)

// Get spreads
spreads, err := pricingAPI.GetSpreads(ctx, requests)
```

### Orderbook API

```go
orderbookAPI := api.NewOrderbookAPI(c)

// Get single orderbook
orderbook, err := orderbookAPI.GetOrderbook(ctx, tokenID)

// Get multiple orderbooks
orderbooks, err := orderbookAPI.GetOrderbooks(ctx, requests)
```

### Trades API

```go
tradesAPI := api.NewTradesAPI(c)

// Get user trades
trades, err := tradesAPI.GetTrades(ctx, types.TradesRequest{
    Market: market,
    Before: timestamp,
})
```

## Order Types

The SDK supports all order types:

- **GTC** - Good-Til-Cancelled
- **GTD** - Good-Til-Date  
- **FOK** - Fill-Or-Kill
- **FAK** - Fill-And-Kill

## WebSocket Events

### Market Channel Events

- **Book Events**: Orderbook snapshots and updates
- **Price Change Events**: Real-time price changes
- **Tick Size Change Events**: Market tick size updates
- **Last Trade Price Events**: Recent trade price updates

### User Channel Events

- **Trade Events**: User trade confirmations and updates
- **Order Events**: Order placement, updates, and cancellations

## Error Handling

The SDK provides comprehensive error handling:

```go
import "github.com/lajosdeme/polymarket-go-api/api"

if err != nil {
    if clobErr := api.NewClobError(0, []byte(err.Error())); api.IsClobError(clobErr) {
        if clobErr.IsAuthenticationError() {
            log.Fatal("Authentication failed")
        }
        if clobErr.IsOrderValidationError() {
            log.Printf("Order validation error: %v", clobErr)
        }
        if clobErr.IsRetryable() {
            // Retry the request
            return
        }
    }
}
```

## Order Creation

Orders require EIP-712 signatures. Here's the basic structure:

```go
order := types.Order{
    Salt:          "random-salt",
    Maker:         makerAddress,
    Signer:        signerAddress,
    Taker:         takerAddress,
    TokenID:       tokenId,
    MakerAmount:   "1000000",    // USDC amount (6 decimals)
    TakerAmount:   "1000000",    // Token amount
    Expiration:    expirationTimestamp,
    Nonce:         "0",
    FeeRateBps:    "0",
    Side:          types.BUY,
    SignatureType:  int(types.EOA),
    Signature:      eip712Signature, // Must be signed
}

postOrder := types.PostOrder{
    Order:     order,
    OrderType: types.GTC,
    Owner:     apiKey,
}
```

## WebSocket Usage

### Market Channel

```go
// Connect to market channel for real-time data
assetIDs := []string{"token-1", "token-2"}
err := wsClient.ConnectMarketChannel(assetIDs)

// Subscribe to additional assets
wsClient.SubscribeToAssets([]string{"token-3"})

// Unsubscribe from assets
wsClient.UnsubscribeFromAssets([]string{"token-1"})
```

### User Channel

```go
// Connect to user channel for personal events
markets := []string{"market-1", "market-2"}
err := wsClient.ConnectUserChannel(markets)
```

## Configuration

### Client Options

```go
c := client.NewClobClient("https://clob.polymarket.com")
c.SetTimeout(60 * time.Second)  // Set HTTP timeout
```

### WebSocket Options

```go
wsClient := client.NewWebSocketClient("", authManager)
wsClient.SetPingInterval(30 * time.Second)  // Set ping interval
```

## Development

### Building

```bash
go build ./...
```

### Testing

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run tests: `go test ./...`
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions:

- GitHub Issues: [Create an issue](https://github.com/lajosdeme/polymarket-go-api/issues)
- Documentation: [Polymarket CLOB API](https://docs.polymarket.com/developers/CLOB)

## Security

- Never commit private keys or API credentials
- Use secure storage for sensitive data
- Validate all inputs before submission
- Use HTTPS connections only
- Implement proper error handling