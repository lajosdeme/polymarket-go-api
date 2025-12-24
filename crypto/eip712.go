package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lajosdeme/polymarket-go-api/types"
)

// EIP712Signer handles EIP-712 signature generation
type EIP712Signer struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    int64
}

// NewEIP712Signer creates a new EIP-712 signer
func NewEIP712Signer(privateKeyHex string, chainID int64) (*EIP712Signer, error) {
	if privateKeyHex == "" {
		return nil, fmt.Errorf("private key cannot be empty")
	}

	// Remove 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// Convert hex to private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	// Get address from private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &EIP712Signer{
		privateKey: privateKey,
		address:    address,
		chainID:    chainID,
	}, nil
}

// GetAddress returns the signer's address
func (s *EIP712Signer) GetAddress() string {
	return s.address.Hex()
}

// SignClobAuth signs a CLOB authentication message
func (s *EIP712Signer) SignClobAuth(timestamp string, nonce uint64) (string, error) {
	// Create EIP-712 typed data
	typedData := types.EIP712TypedData{
		Domain: types.EIP712Domain{
			Name:    "ClobAuthDomain",
			Version: "1",
			ChainID: s.chainID,
		},
		Types: map[string][]types.EIP712Type{
			"ClobAuth": {
				{Name: "address", Type: "address"},
				{Name: "timestamp", Type: "string"},
				{Name: "nonce", Type: "uint256"},
				{Name: "message", Type: "string"},
			},
		},
		Message: types.EIP712ClobAuth{
			Address:   s.address.Hex(),
			Timestamp: timestamp,
			Nonce:     nonce,
			Message:   "This message attests that I control the given wallet",
		},
	}

	return s.SignTypedData(typedData)
}

// SignTypedData signs EIP-712 typed data
func (s *EIP712Signer) SignTypedData(typedData types.EIP712TypedData) (string, error) {
	// Convert to hash
	hash, err := s.hashTypedData(typedData)
	if err != nil {
		return "", fmt.Errorf("failed to hash typed data: %w", err)
	}

	// Sign the hash
	signature, err := crypto.Sign(hash.Bytes(), s.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}

	// Adjust V parameter for EIP-712
	if signature[64] >= 27 {
		signature[64] -= 27
	}

	return hexutil.Encode(signature), nil
}

// hashTypedData creates the hash of EIP-712 typed data
func (s *EIP712Signer) hashTypedData(typedData types.EIP712TypedData) (common.Hash, error) {
	// Hash domain separator
	domainSeparator, err := s.hashDomainSeparator(typedData.Domain)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to hash domain separator: %w", err)
	}

	// Hash message data
	dataHash, err := s.hashMessageData("ClobAuth", typedData.Types, typedData.Message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to hash message data: %w", err)
	}

	// Create final hash
	// keccak256("0x1901" || domainSeparator || dataHash)
	prefix := []byte{0x19, 0x01}
	hashData := append(prefix, domainSeparator.Bytes()...)
	hashData = append(hashData, dataHash.Bytes()...)

	return crypto.Keccak256Hash(hashData), nil
}

// hashDomainSeparator creates the domain separator hash
func (s *EIP712Signer) hashDomainSeparator(domain types.EIP712Domain) (common.Hash, error) {
	// keccak256(
	//   keccak256("EIP712Domain(string name,string version,uint256 chainId)"),
	//   keccak256(name),
	//   keccak256(version),
	//   chainId
	// )

	// Type hash
	typeHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId)"))

	// Name hash
	nameHash := crypto.Keccak256Hash([]byte(domain.Name))

	// Version hash
	versionHash := crypto.Keccak256Hash([]byte(domain.Version))

	// Chain ID
	chainID := new(big.Int).SetInt64(domain.ChainID)

	// Combine and hash
	hashData := make([]byte, 0, 32+32+32+32)
	hashData = append(hashData, typeHash.Bytes()...)
	hashData = append(hashData, nameHash[:]...)
	hashData = append(hashData, versionHash[:]...)
	hashData = append(hashData, common.LeftPadBytes(chainID.Bytes(), 32)...)

	return crypto.Keccak256Hash(hashData), nil
}

// hashMessageData creates the hash of message data
func (s *EIP712Signer) hashMessageData(primaryType string, types map[string][]types.EIP712Type, message interface{}) (common.Hash, error) {
	// Get type definition
	typeDef, exists := types[primaryType]
	if !exists {
		return common.Hash{}, fmt.Errorf("type %s not found", primaryType)
	}

	// Convert message to JSON and encode
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to marshal message: %w", err)
	}

	var messageMap map[string]interface{}
	if err := json.Unmarshal(messageBytes, &messageMap); err != nil {
		return common.Hash{}, fmt.Errorf("failed to unmarshal message: %w", err)
	}

	// Encode each field
	encodedData := make([][]byte, 0, len(typeDef))
	for _, field := range typeDef {
		value, exists := messageMap[field.Name]
		if !exists {
			return common.Hash{}, fmt.Errorf("field %s not found in message", field.Name)
		}

		encoded, err := s.encodeField(field.Type, value, types)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to encode field %s: %w", field.Name, err)
		}

		encodedData = append(encodedData, encoded)
	}

	// Hash the encoded data
	hashData := make([]byte, 0, len(encodedData)*32)
	for _, data := range encodedData {
		hashData = append(hashData, data...)
	}

	return crypto.Keccak256Hash(hashData), nil
}

// encodeField encodes a field according to EIP-712 rules
func (s *EIP712Signer) encodeField(typeStr string, value interface{}, types map[string][]types.EIP712Type) ([]byte, error) {
	// Handle arrays
	if strings.HasSuffix(typeStr, "[]") {
		// TODO: Implement array encoding
		return nil, fmt.Errorf("array encoding not implemented")
	}

	// Handle custom types
	if _, isCustom := types[typeStr]; isCustom {
		hash, err := s.hashMessageData(typeStr, types, value)
		if err != nil {
			return nil, err
		}
		return hash.Bytes(), nil
	}

	// Handle primitive types
	switch typeStr {
	case "address":
		addrStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for address, got %T", value)
		}
		if !common.IsHexAddress(addrStr) {
			return nil, fmt.Errorf("invalid address: %s", addrStr)
		}
		addr := common.HexToAddress(addrStr)
		return crypto.Keccak256Hash(addr.Bytes()).Bytes(), nil

	case "string":
		str, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for string, got %T", value)
		}
		return crypto.Keccak256Hash([]byte(str)).Bytes(), nil

	case "bytes":
		bytesStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for bytes, got %T", value)
		}
		bytes, err := hex.DecodeString(strings.TrimPrefix(bytesStr, "0x"))
		if err != nil {
			return nil, fmt.Errorf("invalid hex string: %w", err)
		}
		return crypto.Keccak256Hash(bytes).Bytes(), nil

	case "bool":
		boolVal, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("expected bool for bool, got %T", value)
		}
		result := make([]byte, 32)
		if boolVal {
			result[31] = 1
		}
		return result, nil

	case "uint256", "uint":
		var bigInt *big.Int
		switch v := value.(type) {
		case uint64:
			bigInt = new(big.Int).SetUint64(v)
		case int64:
			bigInt = big.NewInt(v)
		case string:
			bigInt = new(big.Int)
			if _, ok := bigInt.SetString(v, 10); !ok {
				return nil, fmt.Errorf("invalid number string: %s", v)
			}
		case float64:
			bigInt = new(big.Int).SetInt64(int64(v))
		default:
			return nil, fmt.Errorf("unsupported type for uint256: %T", value)
		}
		return common.LeftPadBytes(bigInt.Bytes(), 32), nil

	default:
		return nil, fmt.Errorf("unsupported type: %s", typeStr)
	}
}

// GenerateSalt generates a random salt for orders
func GenerateSalt() (uint64, error) {
	salt, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 64))
	if err != nil {
		return 0, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt.Uint64(), nil
}
