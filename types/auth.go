package types

// SignatureType represents the type of signature used for authentication
type SignatureType int

const (
	// EOA - Standard Ethereum wallet (MetaMask)
	EOA SignatureType = 0
	// POLY_PROXY - Custom proxy wallet for Magic Link email/Google users
	POLY_PROXY SignatureType = 1
	// GNOSIS_SAFE - Gnosis Safe multisig proxy wallet
	GNOSIS_SAFE SignatureType = 2
)

// String returns the string representation of SignatureType
func (st SignatureType) String() string {
	switch st {
	case EOA:
		return "EOA"
	case POLY_PROXY:
		return "POLY_PROXY"
	case GNOSIS_SAFE:
		return "GNOSIS_SAFE"
	default:
		return "UNKNOWN"
	}
}

// APICredentials represents L2 authentication credentials
type APICredentials struct {
	APIKey     string `json:"apiKey"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

// EIP712Domain represents the EIP-712 domain for CLOB authentication
type EIP712Domain struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	ChainID int64  `json:"chainId"`
}

// EIP712ClobAuth represents the EIP-712 ClobAuth type
type EIP712ClobAuth struct {
	Address   string `json:"address"`
	Timestamp string `json:"timestamp"`
	Nonce     uint64 `json:"nonce"`
	Message   string `json:"message"`
}

// EIP712TypedData represents the complete EIP-712 typed data structure
type EIP712TypedData struct {
	Domain  EIP712Domain            `json:"domain"`
	Types   map[string][]EIP712Type `json:"types"`
	Message EIP712ClobAuth          `json:"message"`
}

// EIP712Type represents a single type in EIP-712 typed data
type EIP712Type struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
