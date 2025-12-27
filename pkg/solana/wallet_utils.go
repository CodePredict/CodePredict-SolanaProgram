package solana

import (
	"encoding/hex"
	"fmt"
	"github.com/gagliardetto/solana-go"
)

// ValidatePublicKey validates a Solana public key
func ValidatePublicKey(pubKeyStr string) (solana.PublicKey, error) {
	if pubKeyStr == "" {
		return solana.PublicKey{}, fmt.Errorf("empty public key")
	}

	pubKey, err := solana.PublicKeyFromBase58(pubKeyStr)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("invalid public key: %w", err)
	}

	return pubKey, nil
}

// ValidatePrivateKey validates a Solana private key
func ValidatePrivateKey(privKeyStr string) error {
	if privKeyStr == "" {
		return fmt.Errorf("empty private key")
	}

	_, err := solana.PrivateKeyFromBase58(privKeyStr)
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	return nil
}

// PublicKeyToHex converts public key to hex string
func PublicKeyToHex(pubKey solana.PublicKey) string {
	return hex.EncodeToString(pubKey[:])
}

// HexToPublicKey converts hex string to public key
func HexToPublicKey(hexStr string) (solana.PublicKey, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return solana.PublicKey{}, err
	}

	if len(bytes) != 32 {
		return solana.PublicKey{}, fmt.Errorf("invalid public key length: expected 32 bytes, got %d", len(bytes))
	}

	var pubKey solana.PublicKey
	copy(pubKey[:], bytes)
	return pubKey, nil
}

// ShortPublicKey returns a shortened version of public key (first 4 + last 4 chars)
func ShortPublicKey(pubKey solana.PublicKey) string {
	keyStr := pubKey.String()
	if len(keyStr) <= 8 {
		return keyStr
	}
	return keyStr[:4] + "..." + keyStr[len(keyStr)-4:]
}


