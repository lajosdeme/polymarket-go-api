package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

// HMACSHA256 generates an HMAC-SHA256 signature
func HMACSHA256(secret, message string) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("secret cannot be empty")
	}
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}

	// Convert secret to bytes
	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		// If not hex, treat as raw string
		secretBytes = []byte(secret)
	}

	// Create HMAC-SHA256 hash
	h := hmac.New(sha256.New, secretBytes)
	h.Write([]byte(message))

	// Return hex encoded signature
	return hex.EncodeToString(h.Sum(nil)), nil
}

// hmacNew creates a new HMAC hash with the given secret
func hmacNew(secret string) (hash.Hash, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}

	// Try to decode as hex first
	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		// If not hex, use raw string
		secretBytes = []byte(secret)
	}

	return hmac.New(sha256.New, secretBytes), nil
}

// SignRequest signs a request with HMAC-SHA256 using the provided secret
func SignRequest(secret, method, path, body string, timestamp int64) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("secret cannot be empty")
	}

	// Create message to sign
	// Format: method + path + body + timestamp
	message := fmt.Sprintf("%s%s%s%d", method, path, body, timestamp)

	return HMACSHA256(secret, message)
}
