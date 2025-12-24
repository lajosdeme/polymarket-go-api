package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/lajosdeme/polymarket-go-api/types"
)

// ClobClient represents the main CLOB API client
type ClobClient struct {
	baseURL     string
	httpClient  *http.Client
	authManager *AuthManager
}

// NewClobClient creates a new CLOB client
func NewClobClient(baseURL string) *ClobClient {
	if baseURL == "" {
		baseURL = "https://clob.polymarket.com"
	}

	return &ClobClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		authManager: NewAuthManager(),
	}
}

// SetTimeout sets HTTP client timeout
func (c *ClobClient) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// GetAuthManager returns the authentication manager
func (c *ClobClient) GetAuthManager() *AuthManager {
	return c.authManager
}

// SetupL1Auth sets up L1 authentication
func (c *ClobClient) SetupL1Auth(privateKeyHex string, signatureType types.SignatureType, funder string) error {
	return c.authManager.SetupL1Auth(privateKeyHex, signatureType, funder)
}

// SetupL2Auth sets up L2 authentication
func (c *ClobClient) SetupL2Auth(apiKey, secret, passphrase string) error {
	return c.authManager.SetupL2Auth(apiKey, secret, passphrase)
}

// DoRequest performs an HTTP request with authentication
func (c *ClobClient) DoRequest(ctx context.Context, method, path string, body interface{}, requireL2Auth bool) ([]byte, error) {
	// Prepare request body
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	requestURL := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add authentication headers
	if requireL2Auth {
		bodyStr := string(reqBody)
		headers, err := c.authManager.GenerateL2Headers(method, path, bodyStr)
		if err != nil {
			return nil, fmt.Errorf("failed to generate L2 headers: %w", err)
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoRequestWithL1Auth performs an HTTP request with L1 authentication
func (c *ClobClient) DoRequestWithL1Auth(ctx context.Context, method, path string, body interface{}, nonce uint64) ([]byte, error) {
	// Prepare request body
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	requestURL := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add L1 authentication headers
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	headers, err := c.authManager.GenerateL1Headers(timestamp, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to generate L1 headers: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoGet performs a GET request with optional authentication
func (c *ClobClient) DoGet(ctx context.Context, path string, requireL2Auth bool, queryParams map[string]string) ([]byte, error) {
	// Build URL with query parameters
	requestURL := c.baseURL + path
	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Add(key, value)
		}
		requestURL += "?" + values.Encode()
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add authentication headers
	if requireL2Auth {
		headers, err := c.authManager.GenerateL2Headers("GET", path, "")
		if err != nil {
			return nil, fmt.Errorf("failed to generate L2 headers: %w", err)
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoGetWithL1Auth performs a GET request with L1 authentication
func (c *ClobClient) DoGetWithL1Auth(ctx context.Context, path string, nonce uint64, queryParams map[string]string) ([]byte, error) {
	// Build URL with query parameters
	requestURL := c.baseURL + path
	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Add(key, value)
		}
		requestURL += "?" + values.Encode()
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add L1 authentication headers
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	headers, err := c.authManager.GenerateL1Headers(timestamp, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to generate L1 headers: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoDelete performs a DELETE request with authentication
func (c *ClobClient) DoDelete(ctx context.Context, path string, body interface{}) ([]byte, error) {
	// Prepare request body
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	// Create request
	requestURL := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set content type
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add authentication headers
	bodyStr := string(reqBody)
	headers, err := c.authManager.GenerateL2Headers("DELETE", path, bodyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to generate L2 headers: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
