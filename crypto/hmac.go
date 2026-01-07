package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

// replaceAll replaces all occurrences of search with replace in s
func replaceAll(s, search, replace string) string {
	return strings.ReplaceAll(s, search, replace)
}

// base64ToBytes converts a base64 string to bytes
// Handles base64url encoding and sanitizes input like the TypeScript version
func base64ToBytes(b64 string) ([]byte, error) {
	// Convert base64url to base64
	sanitized := b64
	sanitized = strings.ReplaceAll(sanitized, "-", "+")
	sanitized = strings.ReplaceAll(sanitized, "_", "/")

	// Remove any non-base64 characters for backwards compatibility
	var cleaned strings.Builder
	for _, ch := range sanitized {
		if (ch >= 'A' && ch <= 'Z') ||
			(ch >= 'a' && ch <= 'z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '+' || ch == '/' || ch == '=' {
			cleaned.WriteRune(ch)
		}
	}

	return base64.StdEncoding.DecodeString(cleaned.String())
}

// bytesToBase64 converts bytes to base64 string
func bytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// BuildPolyHmacSignature builds the canonical Polymarket CLOB HMAC signature
// This matches the TypeScript implementation exactly
func BuildPolyHmacSignature(
	secret string,
	timestamp int64,
	method string,
	requestPath string,
	body *string,
) (string, error) {
	// Build message: timestamp + method + requestPath + [body]
	message := fmt.Sprintf("%d%s%s", timestamp, method, requestPath)
	if body != nil {
		message += *body
	}

	// Import the secret key from base64
	keyData, err := base64ToBytes(secret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret: %w", err)
	}

	// Create HMAC-SHA256
	h := hmac.New(sha256.New, keyData)
	h.Write([]byte(message))
	signatureBytes := h.Sum(nil)

	// Convert to base64
	sig := bytesToBase64(signatureBytes)

	// Convert to URL-safe base64 (but keep '=' padding)
	// Convert '+' to '-'
	// Convert '/' to '_'
	sigUrlSafe := replaceAll(replaceAll(sig, "+", "-"), "/", "_")

	return sigUrlSafe, nil
}
