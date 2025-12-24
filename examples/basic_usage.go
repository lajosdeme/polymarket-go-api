package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lajosdeme/polymarket-go-api/api"
	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/crypto"
	"github.com/lajosdeme/polymarket-go-api/types"
)

func main() {
	// Example 1: Setting up L1 authentication
	fmt.Println("=== L1 Authentication Example ===")
	err := l1AuthenticationExample()
	if err != nil {
		log.Printf("L1 Auth Example Error: %v", err)
	}

	// Example 2: Setting up L2 authentication
	fmt.Println("\n=== L2 Authentication Example ===")
	err = l2AuthenticationExample()
	if err != nil {
		log.Printf("L2 Auth Example Error: %v", err)
	}

	// Example 3: Market data without authentication
	fmt.Println("\n=== Market Data Example ===")
	err = marketDataExample()
	if err != nil {
		log.Printf("Market Data Example Error: %v", err)
	}

	// Example 4: Order management with L2 authentication
	fmt.Println("\n=== Order Management Example ===")
	err = orderManagementExample()
	if err != nil {
		log.Printf("Order Management Example Error: %v", err)
	}

	// Example 5: WebSocket connection
	fmt.Println("\n=== WebSocket Example ===")
	err = webSocketExample()
	if err != nil {
		log.Printf("WebSocket Example Error: %v", err)
	}
}

// l1AuthenticationExample demonstrates L1 authentication setup and API key creation
func l1AuthenticationExample() error {
	// Initialize client with L1 authentication
	c := client.NewClobClient("")

	// Replace with your actual private key (in production, use secure storage)
	privateKey := "0x..." // Your private key here

	// Setup L1 authentication (EOA wallet type)
	err := c.SetupL1Auth(privateKey, types.EOA, "")
	if err != nil {
		return fmt.Errorf("failed to setup L1 auth: %w", err)
	}

	// Create API credentials
	authAPI := api.NewAuthAPI(c)

	// Generate a random nonce
	salt, err := crypto.GenerateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	ctx := context.Background()
	credentials, err := authAPI.CreateAPIKey(ctx, salt)
	if err != nil {
		return fmt.Errorf("failed to create API key: %w", err)
	}

	fmt.Printf("Created API Credentials:\n")
	fmt.Printf("  API Key: %s\n", credentials.APIKey)
	fmt.Printf("  Secret: %s\n", credentials.Secret)
	fmt.Printf("  Passphrase: %s\n", credentials.Passphrase)
	fmt.Printf("  Address: %s\n", c.GetAuthManager().GetAddress())

	return nil
}

// l2AuthenticationExample demonstrates L2 authentication setup
func l2AuthenticationExample() error {
	// Initialize client
	c := client.NewClobClient("")

	// Setup L2 authentication (replace with your actual credentials)
	err := c.SetupL2Auth(
		"your-api-key",
		"your-secret",
		"your-passphrase",
	)
	if err != nil {
		return fmt.Errorf("failed to setup L2 auth: %w", err)
	}

	fmt.Printf("L2 Authentication Setup Complete\n")
	fmt.Printf("Address: %s\n", c.GetAuthManager().GetAddress())
	fmt.Printf("Signature Type: %s\n", c.GetAuthManager().GetSignatureType())

	return nil
}

// marketDataExample demonstrates fetching market data without authentication
func marketDataExample() error {
	c := client.NewClobClient("")
	pricingAPI := api.NewPricingAPI(c)
	orderbookAPI := api.NewOrderbookAPI(c)

	ctx := context.Background()
	tokenID := "1234567890" // Example token ID

	// Get price
	price, err := pricingAPI.GetPrice(ctx, tokenID, types.BUY)
	if err != nil {
		return fmt.Errorf("failed to get price: %w", err)
	}
	fmt.Printf("Current BUY price: %s\n", price.Price)

	// Get midpoint price
	midpoint, err := pricingAPI.GetMidpointPrice(ctx, tokenID)
	if err != nil {
		return fmt.Errorf("failed to get midpoint price: %w", err)
	}
	fmt.Printf("Midpoint price: %s\n", midpoint.Mid)

	// Get orderbook
	orderbook, err := orderbookAPI.GetOrderbook(ctx, tokenID)
	if err != nil {
		return fmt.Errorf("failed to get orderbook: %w", err)
	}
	fmt.Printf("Orderbook:\n")
	fmt.Printf("  Market: %s\n", orderbook.Market)
	fmt.Printf("  Asset ID: %s\n", orderbook.AssetID)
	fmt.Printf("  Bids: %d levels\n", len(orderbook.Bids))
	fmt.Printf("  Asks: %d levels\n", len(orderbook.Asks))
	fmt.Printf("  Min Order Size: %s\n", orderbook.MinOrderSize)
	fmt.Printf("  Tick Size: %s\n", orderbook.TickSize)

	// Get price history
	now := time.Now().Unix()
	startTime := now - 86400 // 24 hours ago
	priceHistoryReq := types.PriceHistoryRequest{
		Market:   "example-market-condition-id",
		StartTs:  &startTime,
		EndTs:    &now,
		Interval: "1h",
		Fidelity: intPtr(60), // 60 minute intervals
	}

	history, err := pricingAPI.GetPriceHistory(ctx, priceHistoryReq)
	if err != nil {
		return fmt.Errorf("failed to get price history: %w", err)
	}
	fmt.Printf("Price History Points: %d\n", len(history.History))

	return nil
}

// orderManagementExample demonstrates order operations
func orderManagementExample() error {
	c := client.NewClobClient("")

	// Setup L2 authentication
	err := c.SetupL2Auth(
		"your-api-key",
		"your-secret",
		"your-passphrase",
	)
	if err != nil {
		return fmt.Errorf("failed to setup L2 auth: %w", err)
	}

	ordersAPI := api.NewOrdersAPI(c)
	tradesAPI := api.NewTradesAPI(c)

	ctx := context.Background()

	// Note: Real order creation requires proper EIP-712 signing of the order
	// This is just showing the structure - see documentation for complete implementation

	// Example order structure (requires proper signing before submission):
	// salt, err := crypto.GenerateSalt()
	// if err != nil {
	//     return fmt.Errorf("failed to generate salt: %w", err)
	// }
	//
	// postOrder := types.PostOrder{
	//     Order: types.Order{
	//         Salt:          fmt.Sprintf("%d", salt),
	//         Maker:         c.GetAuthManager().GetAddress(),
	//         Signer:        c.GetAuthManager().GetAddress(),
	//         Taker:         "0x0000000000000000000000000000000000000",
	//         TokenID:       "1234567890",
	//         MakerAmount:   "1000000", // 1 USDC in wei (6 decimals)
	//         TakerAmount:   "1000000", // 1 token
	//         Expiration:    fmt.Sprintf("%d", time.Now().Add(24*time.Hour).Unix()),
	//         Nonce:         "0",
	//         FeeRateBps:    "0",
	//         Side:          types.BUY,
	//         SignatureType: int(types.EOA),
	//         Signature:     "0x...", // Proper signature needed here
	//     },
	//     OrderType: types.GTC,
	//     Owner:     c.GetAuthManager().GetAPICredentials().APIKey,
	// }

	// Place order (commented out as it requires proper signing)
	/*
		response, err := ordersAPI.PlaceOrder(ctx, postOrder)
		if err != nil {
			return fmt.Errorf("failed to place order: %w", err)
		}
		fmt.Printf("Order placed successfully:\n")
		fmt.Printf("  Order ID: %s\n", response.OrderID)
		fmt.Printf("  Success: %t\n", response.Success)
	*/

	// Get active orders
	orders, err := ordersAPI.GetActiveOrders(ctx, "", "", "")
	if err != nil {
		return fmt.Errorf("failed to get active orders: %w", err)
	}
	fmt.Printf("Active Orders: %d\n", len(orders))

	// Get trades
	trades, err := tradesAPI.GetTrades(ctx, types.TradesRequest{})
	if err != nil {
		return fmt.Errorf("failed to get trades: %w", err)
	}
	fmt.Printf("Trades: %d\n", len(trades))

	return nil
}

// webSocketExample demonstrates WebSocket connection
func webSocketExample() error {
	c := client.NewClobClient("")

	// Setup L2 authentication for user channel
	err := c.SetupL2Auth(
		"your-api-key",
		"your-secret",
		"your-passphrase",
	)
	if err != nil {
		return fmt.Errorf("failed to setup L2 auth: %w", err)
	}

	wsClient := client.NewWebSocketClient("", c.GetAuthManager())

	// Set up event handlers
	wsClient.SetBookMessageHandler(func(event *types.WebSocketBookEvent) {
		fmt.Printf("Book Event - Asset: %s, Bids: %d, Asks: %d\n",
			event.AssetID, len(event.Bids), len(event.Asks))
	})

	wsClient.SetPriceChangeMessageHandler(func(event *types.WebSocketPriceChangeEvent) {
		fmt.Printf("Price Change - Market: %s, Changes: %d\n",
			event.Market, len(event.PriceChanges))
	})

	wsClient.SetErrorHandler(func(err error) {
		fmt.Printf("WebSocket Error: %v\n", err)
	})

	// Connect to market channel
	assetIDs := []string{"1234567890", "0987654321"}
	err = wsClient.ConnectMarketChannel(assetIDs)
	if err != nil {
		return fmt.Errorf("failed to connect to market channel: %w", err)
	}

	fmt.Printf("Connected to market channel for assets: %v\n", assetIDs)

	// Keep connection alive for demo
	time.Sleep(30 * time.Second)

	// Close connection
	err = wsClient.Close()
	if err != nil {
		return fmt.Errorf("failed to close WebSocket: %w", err)
	}

	return nil
}

// Helper function to create int pointer
func intPtr(i int) *int {
	return &i
}
