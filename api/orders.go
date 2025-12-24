package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// OrdersAPI handles order-related operations
type OrdersAPI struct {
	client *client.ClobClient
}

// NewOrdersAPI creates a new OrdersAPI instance
func NewOrdersAPI(client *client.ClobClient) *OrdersAPI {
	return &OrdersAPI{
		client: client,
	}
}

// PlaceOrder places a single order
func (o *OrdersAPI) PlaceOrder(ctx context.Context, order types.PostOrder) (*types.OrderResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for placing orders")
	}

	body, err := o.client.DoRequest(ctx, "POST", "/order", order, true)
	if err != nil {
		return nil, err
	}

	var response types.OrderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// PlaceOrders places multiple orders (batch)
func (o *OrdersAPI) PlaceOrders(ctx context.Context, orders []types.PostOrder) ([]types.OrderResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for placing orders")
	}

	body, err := o.client.DoRequest(ctx, "POST", "/orders", orders, true)
	if err != nil {
		return nil, err
	}

	var responses []types.OrderResponse
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return responses, nil
}

// GetOrder gets information about an existing order
func (o *OrdersAPI) GetOrder(ctx context.Context, orderID string) (*types.OpenOrder, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for getting order information")
	}

	body, err := o.client.DoGet(ctx, "/data/order/"+orderID, true, nil)
	if err != nil {
		return nil, err
	}

	var order types.OpenOrder
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &order, nil
}

// GetActiveOrders gets active orders for specific filters
func (o *OrdersAPI) GetActiveOrders(ctx context.Context, id, market, assetID string) ([]types.OpenOrder, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for getting active orders")
	}

	// Build query parameters
	queryParams := make(map[string]string)
	if id != "" {
		queryParams["id"] = id
	}
	if market != "" {
		queryParams["market"] = market
	}
	if assetID != "" {
		queryParams["asset_id"] = assetID
	}

	body, err := o.client.DoGet(ctx, "/data/orders", true, queryParams)
	if err != nil {
		return nil, err
	}

	var orders []types.OpenOrder
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return orders, nil
}

// CancelOrder cancels a single order
func (o *OrdersAPI) CancelOrder(ctx context.Context, orderID string) (*types.CancelResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for canceling orders")
	}

	request := types.CancelOrderRequest{
		OrderID: orderID,
	}

	body, err := o.client.DoDelete(ctx, "/order", request)
	if err != nil {
		return nil, err
	}

	var response types.CancelResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelOrders cancels multiple orders
func (o *OrdersAPI) CancelOrders(ctx context.Context, orderIDs []string) (*types.CancelResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for canceling orders")
	}

	request := types.CancelOrdersRequest{
		OrderIDs: orderIDs,
	}

	body, err := o.client.DoDelete(ctx, "/orders", request)
	if err != nil {
		return nil, err
	}

	var response types.CancelResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelAllOrders cancels all open orders
func (o *OrdersAPI) CancelAllOrders(ctx context.Context) (*types.CancelResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for canceling orders")
	}

	body, err := o.client.DoDelete(ctx, "/cancel-all", nil)
	if err != nil {
		return nil, err
	}

	var response types.CancelResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// CancelMarketOrders cancels orders from a specific market
func (o *OrdersAPI) CancelMarketOrders(ctx context.Context, market, assetID string) (*types.CancelResponse, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for canceling orders")
	}

	request := types.CancelMarketOrdersRequest{
		Market:  market,
		AssetID: assetID,
	}

	body, err := o.client.DoDelete(ctx, "/cancel-market-orders", request)
	if err != nil {
		return nil, err
	}

	var response types.CancelResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// CheckOrderScoring checks if an order is eligible for rewards scoring
func (o *OrdersAPI) CheckOrderScoring(ctx context.Context, orderID string) (*types.OrderScoring, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for checking order scoring")
	}

	queryParams := map[string]string{
		"order_id": orderID,
	}

	body, err := o.client.DoGet(ctx, "/order-scoring", true, queryParams)
	if err != nil {
		return nil, err
	}

	var scoring types.OrderScoring
	if err := json.Unmarshal(body, &scoring); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &scoring, nil
}

// CheckOrdersScoring checks if multiple orders are eligible for rewards scoring
func (o *OrdersAPI) CheckOrdersScoring(ctx context.Context, orderIDs []string) (*types.OrdersScoring, error) {
	// Validate required L2 authentication
	if !o.client.GetAuthManager().HasL2Auth() {
		return nil, fmt.Errorf("L2 authentication required for checking order scoring")
	}

	request := types.OrdersScoringRequest{
		OrderIDs: orderIDs,
	}

	body, err := o.client.DoRequest(ctx, "POST", "/orders-scoring", request, true)
	if err != nil {
		return nil, err
	}

	var scoring types.OrdersScoring
	if err := json.Unmarshal(body, &scoring); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &scoring, nil
}
