package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// TradesAPI handles trades operations
type TradesAPI struct {
	client *client.ClobClient
}

// NewTradesAPI creates a new TradesAPI instance
func NewTradesAPI(client *client.ClobClient) *TradesAPI {
	return &TradesAPI{
		client: client,
	}
}

// GetTrades gets trades for the authenticated user based on provided filters
func (t *TradesAPI) GetTrades(ctx context.Context, request types.TradesRequest) ([]types.Trade, error) {
	// Validate required L2 authentication
	if !t.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for getting trades")
	}

	// Build query parameters
	queryParams := make(map[string]string)
	if request.ID != "" {
		queryParams["id"] = request.ID
	}
	if request.Taker != "" {
		queryParams["taker"] = request.Taker
	}
	if request.Maker != "" {
		queryParams["maker"] = request.Maker
	}
	if request.Market != "" {
		queryParams["market"] = request.Market
	}
	if request.Before != "" {
		queryParams["before"] = request.Before
	}
	if request.After != "" {
		queryParams["after"] = request.After
	}

	body, err := t.client.DoGet(ctx, "/data/trades", true, queryParams)
	if err != nil {
		return nil, err
	}

	var trades []types.Trade
	if err := json.Unmarshal(body, &trades); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return trades, nil
}
