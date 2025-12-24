package api

import (
	"fmt"
	"strings"
)

// ErrorCode represents API error codes
type ErrorCode string

const (
	// Invalid order errors
	ErrInvalidOrderMinTickSize ErrorCode = "INVALID_ORDER_MIN_TICK_SIZE"
	ErrInvalidOrderMinSize     ErrorCode = "INVALID_ORDER_MIN_SIZE"
	ErrInvalidOrderDuplicated  ErrorCode = "INVALID_ORDER_DUPLICATED"
	ErrInvalidOrderBalance     ErrorCode = "INVALID_ORDER_NOT_ENOUGH_BALANCE"
	ErrInvalidOrderExpiration  ErrorCode = "INVALID_ORDER_EXPIRATION"
	ErrInvalidOrderError       ErrorCode = "INVALID_ORDER_ERROR"

	// Execution errors
	ErrExecutionError     ErrorCode = "EXECUTION_ERROR"
	ErrOrderDelayed       ErrorCode = "ORDER_DELAYED"
	ErrDelayingOrderError ErrorCode = "DELAYING_ORDER_ERROR"
	ErrFOKOrderNotFilled  ErrorCode = "FOK_ORDER_NOT_FILLED_ERROR"
	ErrMarketNotReady     ErrorCode = "MARKET_NOT_READY"

	// Authentication errors
	ErrInvalidSignature     ErrorCode = "INVALID_SIGNATURE"
	ErrNonceAlreadyUsed     ErrorCode = "NONCE_ALREADY_USED"
	ErrInvalidFunderAddress ErrorCode = "INVALID_FUNDER_ADDRESS"

	// General errors
	ErrInternalError ErrorCode = "INTERNAL_ERROR"
	ErrRateLimited   ErrorCode = "RATE_LIMITED"
	ErrUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrForbidden     ErrorCode = "FORBIDDEN"
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrBadRequest    ErrorCode = "BAD_REQUEST"
)

// ClobError represents a CLOB API error
type ClobError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Success    bool      `json:"success"`
	StatusCode int       `json:"status_code,omitempty"`
	Details    string    `json:"details,omitempty"`
}

// Error implements the error interface
func (e *ClobError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("CLOB API Error [%s]: %s - %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("CLOB API Error [%s]: %s", e.Code, e.Message)
}

// IsRetryable returns true if the error is retryable
func (e *ClobError) IsRetryable() bool {
	switch e.Code {
	case ErrInternalError, ErrRateLimited, ErrOrderDelayed, ErrExecutionError:
		return true
	default:
		return false
	}
}

// IsAuthenticationError returns true if the error is authentication related
func (e *ClobError) IsAuthenticationError() bool {
	switch e.Code {
	case ErrInvalidSignature, ErrNonceAlreadyUsed, ErrInvalidFunderAddress, ErrUnauthorized, ErrForbidden:
		return true
	default:
		return false
	}
}

// IsOrderValidationError returns true if the error is order validation related
func (e *ClobError) IsOrderValidationError() bool {
	switch e.Code {
	case ErrInvalidOrderMinTickSize, ErrInvalidOrderMinSize, ErrInvalidOrderDuplicated,
		ErrInvalidOrderBalance, ErrInvalidOrderExpiration, ErrInvalidOrderError:
		return true
	default:
		return false
	}
}

// NewClobError creates a new ClobError from HTTP status and response body
func NewClobError(statusCode int, responseBody []byte) *ClobError {
	bodyStr := string(responseBody)

	// Try to parse known error codes
	for _, errCode := range []ErrorCode{
		ErrInvalidOrderMinTickSize, ErrInvalidOrderMinSize, ErrInvalidOrderDuplicated,
		ErrInvalidOrderBalance, ErrInvalidOrderExpiration, ErrInvalidOrderError,
		ErrExecutionError, ErrOrderDelayed, ErrDelayingOrderError,
		ErrFOKOrderNotFilled, ErrMarketNotReady, ErrInvalidSignature,
		ErrNonceAlreadyUsed, ErrInvalidFunderAddress,
	} {
		if strings.Contains(bodyStr, string(errCode)) {
			return &ClobError{
				Code:       errCode,
				Message:    extractErrorMessage(bodyStr),
				Success:    false,
				StatusCode: statusCode,
				Details:    bodyStr,
			}
		}
	}

	// Default error based on HTTP status
	var code ErrorCode
	var message string

	switch {
	case statusCode >= 400 && statusCode < 500:
		code = ErrBadRequest
		if statusCode == 401 {
			code = ErrUnauthorized
		}
		if statusCode == 403 {
			code = ErrForbidden
		}
		if statusCode == 404 {
			code = ErrNotFound
		}
		message = fmt.Sprintf("Client error: %d", statusCode)
	case statusCode >= 500:
		code = ErrInternalError
		message = fmt.Sprintf("Server error: %d", statusCode)
	default:
		code = ErrInternalError
		message = fmt.Sprintf("HTTP error: %d", statusCode)
	}

	return &ClobError{
		Code:       code,
		Message:    message,
		Success:    false,
		StatusCode: statusCode,
		Details:    bodyStr,
	}
}

// extractErrorMessage attempts to extract a human-readable error message from response body
func extractErrorMessage(bodyStr string) string {
	// Look for common error message patterns
	if strings.Contains(bodyStr, "not enough balance") {
		return "Insufficient balance or allowance for order"
	}
	if strings.Contains(bodyStr, "breaks minimum tick size") {
		return "Order price breaks minimum tick size rules"
	}
	if strings.Contains(bodyStr, "lower than the minimum") {
		return "Order size below minimum threshold"
	}
	if strings.Contains(bodyStr, "Duplicated") {
		return "Duplicate order already exists"
	}
	if strings.Contains(bodyStr, "before now") {
		return "Order expiration time is in the past"
	}
	if strings.Contains(bodyStr, "INVALID_SIGNATURE") {
		return "Invalid wallet signature"
	}
	if strings.Contains(bodyStr, "NONCE_ALREADY_USED") {
		return "Nonce has already been used"
	}
	if strings.Contains(bodyStr, "Invalid Funder Address") {
		return "Invalid funder address"
	}

	// Return a generic message if no specific pattern found
	return "API error occurred"
}

// IsClobError checks if an error is a ClobError
func IsClobError(err error) bool {
	_, ok := err.(*ClobError)
	return ok
}
