# Polymarket Go API SDK

A comprehensive Go SDK for interacting with the Polymarket CLOB (Centralized Limit Order Book) REST API and WebSocket endpoints.

## Features

- **Full API Coverage**: Complete implementation of all documented Polymarket CLOB API endpoints
- **Gamma API Support**: Access market metadata, events, tags, and search functionality
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
    
    // Create Gamma API client for market metadata
    gammaClient := client.NewGammaClient("")
    gammaAPI := api.NewGammaAPI(gammaClient)
    
    // Get market price
    price, err := pricingAPI.GetPrice(context.Background(), "token-id", types.BUY)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Current BUY price: %s\n", price.Price)
    
    // Get active markets from Gamma API
    markets, err := gammaAPI.GetActiveMarkets(context.Background(), intPtr(10), intPtr(0))
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d active markets\n", len(markets))
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

### Gamma API

```go
gammaClient := client.NewGammaClient("")
gammaAPI := api.NewGammaAPI(gammaClient)

// Get markets with optional filtering
markets, err := gammaAPI.GetMarkets(ctx, &types.MarketFilters{
    Closed: boolPtr(false),
    Limit:  intPtr(50),
    Offset: intPtr(0),
})

// Get market by slug
market, err := gammaAPI.GetMarketBySlug(ctx, "market-slug", boolPtr(true))

// Get events with filtering
events, err := gammaAPI.GetEvents(ctx, &types.EventFilters{
    TagID:   intPtr(123),
    Closed:  boolPtr(false),
    Limit:   intPtr(25),
})

// Get event by slug
event, err := gammaAPI.GetEventBySlug(ctx, "event-slug", nil, nil)

// Get all tags
tags, err := gammaAPI.GetTags(ctx, &types.TagFilters{})

// Get tag by slug
tag, err := gammaAPI.GetTagBySlug(ctx, "politics", boolPtr(true))

// Search across markets, events, and profiles
searchResult, err := gammaAPI.Search(ctx, &types.SearchFilters{
    Query: "election",
    Limit: intPtr(10),
})

// Convenience methods
activeMarkets, err := gammaAPI.GetActiveMarkets(ctx, intPtr(100), intPtr(0))
activeEvents, err := gammaAPI.GetActiveEvents(ctx, intPtr(50), intPtr(0))
marketsByTag, err := gammaAPI.GetMarketsByTag(ctx, 123, intPtr(25), intPtr(0))
```

## Order Types

The SDK supports all order types:

- **GTC** - Good-Til-Cancelled
- **GTD** - Good-Til-Date  
- **FOK** - Fill-Or-Kill
- **FAK** - Fill-And-Kill

## Gamma API Overview

The Gamma API provides comprehensive market metadata, events, and search functionality for Polymarket. It organizes markets into events and provides tagging for categorization.

### Key Concepts

- **Markets**: Individual trading markets that map to CLOB token IDs
- **Events**: Collections of related markets (e.g., "Who will win the election?" with multiple outcome markets)
- **Tags**: Categorization system for filtering markets by topic, sport, etc.

### Common Use Cases

1. **Fetch by Slug**: Get specific markets/events using URLs from Polymarket frontend
2. **Tag Filtering**: Filter markets by category or sport using tag IDs
3. **Active Markets**: Retrieve all currently active trading opportunities
4. **Search**: Search across markets, events, and user profiles

### Pagination

Most Gamma API endpoints support pagination with `limit` and `offset` parameters for handling large datasets efficiently.

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

### Gamma Client Options

```go
gammaClient := client.NewGammaClient("https://gamma-api.polymarket.com")
// Gamma API endpoints are public and don't require authentication
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