package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// OrderbookAPI handles orderbook operations
type OrderbookAPI struct {
	client *client.ClobClient
}

// NewOrderbookAPI creates a new OrderbookAPI instance
func NewOrderbookAPI(client *client.ClobClient) *OrderbookAPI {
	return &OrderbookAPI{
		client: client,
	}
}

// GetOrderbook retrieves the order book summary for a specific token
func (o *OrderbookAPI) GetOrderbook(ctx context.Context, tokenID string) (*types.Orderbook, error) {
	queryParams := map[string]string{
		"token_id": tokenID,
	}

	body, err := o.client.DoGet(ctx, "/book", false, queryParams)
	if err != nil {
		return nil, err
	}

	var orderbook types.Orderbook
	if err := json.Unmarshal(body, &orderbook); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &orderbook, nil
}

// GetOrderbooks retrieves order book summaries for specified tokens via POST request
func (o *OrderbookAPI) GetOrderbooks(ctx context.Context, requests []types.OrderbooksRequest) ([]types.Orderbook, error) {
	body, err := o.client.DoRequest(ctx, "POST", "/books", requests, false)
	if err != nil {
		return nil, err
	}

	var orderbooks []types.Orderbook
	if err := json.Unmarshal(body, &orderbooks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return orderbooks, nil
}
