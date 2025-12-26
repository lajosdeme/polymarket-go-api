package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
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
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
			"ClobAuth": {
				{Name: "address", Type: "address"},
				{Name: "timestamp", Type: "string"},
				{Name: "nonce", Type: "uint256"},
				{Name: "message", Type: "string"},
			},
		},
		PrimaryType: "ClobAuth",
		Domain: apitypes.TypedDataDomain{
			Name:    "ClobAuthDomain",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(137),
		},
		Message: apitypes.TypedDataMessage{
			"address":   s.address.Hex(),
			"timestamp": timestamp,
			"nonce":     fmt.Sprintf("%d", nonce),
			"message":   "This message attests that I control the given wallet",
		},
	}

	sig, err := s.SignTypedData(typedData)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(sig), nil
}

// SignTypedData signs typed data and returns the signature
func (s *EIP712Signer) SignTypedData(typedData apitypes.TypedData) ([]byte, error) {
	hash, err := s.EncodeForSigning(typedData)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(hash.Bytes(), s.privateKey)
	if err != nil {
		return nil, err
	}

	sig[64] += 27
	return sig, nil
}

// EncodeForSigning encodes the typed data for signing
func (s *EIP712Signer) EncodeForSigning(typedData apitypes.TypedData) (common.Hash, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := common.BytesToHash(crypto.Keccak256(rawData))
	return hash, nil
}

// VerifySig verifies signature with recovered address
func VerifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	fmt.Printf("the recovered address: %v \n", recoveredAddr)
	return from == recoveredAddr.Hex()
}

// GenerateSalt generates a random salt for orders
func GenerateSalt() (uint64, error) {
	salt, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 64))
	if err != nil {
		return 0, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt.Uint64(), nil
}
