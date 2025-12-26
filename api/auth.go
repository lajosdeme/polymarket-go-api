package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// AuthAPI handles authentication operations
type AuthAPI struct {
	client *client.ClobClient
}

// NewAuthAPI creates a new AuthAPI instance
func NewAuthAPI(client *client.ClobClient) *AuthAPI {
	return &AuthAPI{
		client: client,
	}
}

// CreateAPIKey creates new API credentials for user
func (a *AuthAPI) CreateAPIKey(ctx context.Context, nonce uint64) (*types.APICredentials, error) {
	// Validate required L1 authentication
	if !a.client.GetAuthManager().HasL1Auth() {
		return nil, fmt.Errorf("L1 authentication required for creating API keys")
	}

	// Get server timestamp
	timestamp, err := a.GetServerTime(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get server time: %w", err)
	}

	body, err := a.client.DoRequestWithL1Auth(ctx, "POST", "/auth/api-key", nil, nonce, timestamp)
	if err != nil {
		return nil, err
	}

	var credentials types.APICredentials
	if err := json.Unmarshal(body, &credentials); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &credentials, nil
}

// DeriveAPIKey derives existing API credentials for user
func (a *AuthAPI) DeriveAPIKey(ctx context.Context, nonce uint64) (*types.APICredentials, error) {
	// Validate required L1 authentication
	if !a.client.GetAuthManager().HasL1Auth() {
		return nil, fmt.Errorf("L1 authentication required for deriving API keys")
	}

	// Get server timestamp
	timestamp, err := a.GetServerTime(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get server time: %w", err)
	}

	body, err := a.client.DoGetWithL1Auth(ctx, "/auth/derive-api-key", nonce, timestamp, nil)
	if err != nil {
		return nil, err
	}

	var credentials types.APICredentials
	if err := json.Unmarshal(body, &credentials); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &credentials, nil
}

// GetServerTime gets the current server timestamp
func (a *AuthAPI) GetServerTime(ctx context.Context) (int64, error) {
	body, err := a.client.DoGet(ctx, "/time", false, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get server time: %w", err)
	}

	// The response is just a number as string
	timestampStr := string(body)
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse timestamp: %w", err)
	}

	return timestamp, nil
}
