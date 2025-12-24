package client

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lajosdeme/polymarket-go-api/crypto"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// AuthLevel represents the authentication level
type AuthLevel int

const (
	// AuthLevelNone - No authentication
	AuthLevelNone AuthLevel = iota
	// AuthLevelL1 - L1 authentication (private key)
	AuthLevelL1 AuthLevel = iota
	// AuthLevelL2 - L2 authentication (API credentials)
	AuthLevelL2 AuthLevel = iota
)

// AuthManager handles authentication for the CLOB API
type AuthManager struct {
	authLevel      AuthLevel
	signer         *crypto.EIP712Signer
	apiCredentials *types.APICredentials
	address        string
	signatureType  types.SignatureType
	funder         string
}

// NewAuthManager creates a new authentication manager
func NewAuthManager() *AuthManager {
	return &AuthManager{
		authLevel: AuthLevelNone,
	}
}

// SetupL1Auth sets up L1 authentication with private key
func (am *AuthManager) SetupL1Auth(privateKeyHex string, signatureType types.SignatureType, funder string) error {
	if privateKeyHex == "" {
		return fmt.Errorf("private key cannot be empty")
	}

	// Create EIP-712 signer
	signer, err := crypto.NewEIP712Signer(privateKeyHex, 137) // Polygon chain ID
	if err != nil {
		return fmt.Errorf("failed to create signer: %w", err)
	}

	am.authLevel = AuthLevelL1
	am.signer = signer
	am.address = signer.GetAddress()
	am.signatureType = signatureType
	am.funder = funder

	return nil
}

// SetupL2Auth sets up L2 authentication with API credentials
func (am *AuthManager) SetupL2Auth(apiKey, secret, passphrase string) error {
	if apiKey == "" || secret == "" || passphrase == "" {
		return fmt.Errorf("API credentials cannot be empty")
	}

	am.authLevel = AuthLevelL2
	am.apiCredentials = &types.APICredentials{
		APIKey:     apiKey,
		Secret:     secret,
		Passphrase: passphrase,
	}

	return nil
}

// GetAddress returns the authenticated address
func (am *AuthManager) GetAddress() string {
	return am.address
}

// GetSignatureType returns the signature type
func (am *AuthManager) GetSignatureType() types.SignatureType {
	return am.signatureType
}

// GetFunder returns the funder address
func (am *AuthManager) GetFunder() string {
	return am.funder
}

// GetAPICredentials returns the API credentials
func (am *AuthManager) GetAPICredentials() *types.APICredentials {
	return am.apiCredentials
}

// GetAuthLevel returns the current authentication level
func (am *AuthManager) GetAuthLevel() AuthLevel {
	return am.authLevel
}

// SignL1Message signs a message using L1 authentication
func (am *AuthManager) SignL1Message(timestamp string, nonce uint64) (string, error) {
	if am.authLevel < AuthLevelL1 {
		return "", fmt.Errorf("L1 authentication required")
	}

	if am.signer == nil {
		return "", fmt.Errorf("signer not initialized")
	}

	return am.signer.SignClobAuth(timestamp, nonce)
}

// GenerateL1Headers generates L1 authentication headers
func (am *AuthManager) GenerateL1Headers(timestamp string, nonce uint64) (map[string]string, error) {
	if am.authLevel < AuthLevelL1 {
		return nil, fmt.Errorf("L1 authentication required")
	}

	signature, err := am.SignL1Message(timestamp, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	headers := map[string]string{
		"POLY_ADDRESS":   am.address,
		"POLY_SIGNATURE": signature,
		"POLY_TIMESTAMP": timestamp,
		"POLY_NONCE":     strconv.FormatUint(nonce, 10),
	}

	return headers, nil
}

// GenerateL2Headers generates L2 authentication headers
func (am *AuthManager) GenerateL2Headers(method, path, body string) (map[string]string, error) {
	if am.authLevel < AuthLevelL2 {
		return nil, fmt.Errorf("L2 authentication required")
	}

	if am.apiCredentials == nil {
		return nil, fmt.Errorf("API credentials not initialized")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Generate HMAC signature
	signature, err := crypto.SignRequest(am.apiCredentials.Secret, method, path, body, time.Now().Unix())
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	headers := map[string]string{
		"POLY_ADDRESS":    am.address,
		"POLY_SIGNATURE":  signature,
		"POLY_TIMESTAMP":  timestamp,
		"POLY_API_KEY":    am.apiCredentials.APIKey,
		"POLY_PASSPHRASE": am.apiCredentials.Passphrase,
	}

	return headers, nil
}

// IsAuthenticated returns true if authenticated at any level
func (am *AuthManager) IsAuthenticated() bool {
	return am.authLevel > AuthLevelNone
}

// HasL1Auth returns true if L1 authentication is set up
func (am *AuthManager) HasL1Auth() bool {
	return am.authLevel >= AuthLevelL1
}

// HasL2Auth returns true if L2 authentication is set up
func (am *AuthManager) HasL2Auth() bool {
	return am.authLevel >= AuthLevelL2
}

// Clear clears all authentication
func (am *AuthManager) Clear() {
	am.authLevel = AuthLevelNone
	am.signer = nil
	am.apiCredentials = nil
	am.address = ""
	am.signatureType = types.EOA
	am.funder = ""
}
