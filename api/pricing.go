package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// PricingAPI handles pricing and market data operations
type PricingAPI struct {
	client *client.ClobClient
}

// NewPricingAPI creates a new PricingAPI instance
func NewPricingAPI(client *client.ClobClient) *PricingAPI {
	return &PricingAPI{
		client: client,
	}
}

// GetPrice gets the market price for a specific token and side
func (p *PricingAPI) GetPrice(ctx context.Context, tokenID string, side types.OrderSide) (*types.Price, error) {
	queryParams := map[string]string{
		"token_id": tokenID,
		"side":     string(side),
	}

	body, err := p.client.DoGet(ctx, "/price", false, queryParams)
	if err != nil {
		return nil, err
	}

	var price types.Price
	if err := json.Unmarshal(body, &price); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &price, nil
}

// GetPrices gets market prices for multiple tokens
func (p *PricingAPI) GetPrices(ctx context.Context) (*types.PricesResponse, error) {
	body, err := p.client.DoGet(ctx, "/prices", false, nil)
	if err != nil {
		return nil, err
	}

	var prices types.PricesResponse
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &prices, nil
}

// GetPricesByRequest gets market prices for specified tokens and sides via POST request
func (p *PricingAPI) GetPricesByRequest(ctx context.Context, requests []types.PricesRequest) (*types.PricesResponse, error) {
	body, err := p.client.DoRequest(ctx, "POST", "/prices", requests, false)
	if err != nil {
		return nil, err
	}

	var prices types.PricesResponse
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &prices, nil
}

// GetMidpointPrice gets the midpoint price for a specific token
func (p *PricingAPI) GetMidpointPrice(ctx context.Context, tokenID string) (*types.MidpointResponse, error) {
	queryParams := map[string]string{
		"token_id": tokenID,
	}

	body, err := p.client.DoGet(ctx, "/midpoint", false, queryParams)
	if err != nil {
		return nil, err
	}

	var midpoint types.MidpointResponse
	if err := json.Unmarshal(body, &midpoint); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &midpoint, nil
}

// GetPriceHistory fetches historical price data for a specified market token
func (p *PricingAPI) GetPriceHistory(ctx context.Context, request types.PriceHistoryRequest) (*types.PriceHistoryResponse, error) {
	queryParams := make(map[string]string)
	queryParams["market"] = request.Market

	if request.StartTs != nil {
		queryParams["startTs"] = fmt.Sprintf("%d", *request.StartTs)
	}
	if request.EndTs != nil {
		queryParams["endTs"] = fmt.Sprintf("%d", *request.EndTs)
	}
	if request.Interval != "" {
		queryParams["interval"] = request.Interval
	}
	if request.Fidelity != nil {
		queryParams["fidelity"] = fmt.Sprintf("%d", *request.Fidelity)
	}

	body, err := p.client.DoGet(ctx, "/prices-history", false, queryParams)
	if err != nil {
		return nil, err
	}

	var history types.PriceHistoryResponse
	if err := json.Unmarshal(body, &history); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &history, nil
}

// GetSpreads retrieves bid-ask spreads for multiple tokens
func (p *PricingAPI) GetSpreads(ctx context.Context, requests []types.SpreadsRequest) (*types.SpreadsResponse, error) {
	body, err := p.client.DoRequest(ctx, "POST", "/spreads", requests, false)
	if err != nil {
		return nil, err
	}

	var spreads types.SpreadsResponse
	if err := json.Unmarshal(body, &spreads); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &spreads, nil
}
