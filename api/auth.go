package api

import (
	"context"
	"encoding/json"
	"fmt"

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

	body, err := a.client.DoRequestWithL1Auth(ctx, "POST", "/auth/api-key", nil, nonce)
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

	body, err := a.client.DoGetWithL1Auth(ctx, "/auth/derive-api-key", nonce, nil)
	if err != nil {
		return nil, err
	}

	var credentials types.APICredentials
	if err := json.Unmarshal(body, &credentials); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &credentials, nil
}
