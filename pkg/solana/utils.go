package solana

import (
	"encoding/base64"
	"errors"
	"github.com/gagliardetto/solana-go"
)

// PublicKeyFromString converts a string to PublicKey
func PublicKeyFromString(s string) (solana.PublicKey, error) {
	if s == "" {
		return solana.PublicKey{}, errors.New("empty public key string")
	}
	return solana.PublicKeyFromBase58(s)
}

// PublicKeyToString converts PublicKey to string
func PublicKeyToString(pk solana.PublicKey) string {
	return pk.String()
}

// SignatureFromString converts a string to Signature
func SignatureFromString(s string) (solana.Signature, error) {
	if s == "" {
		return solana.Signature{}, errors.New("empty signature string")
	}
	
	// Try base58 first
	sig, err := solana.SignatureFromBase58(s)
	if err == nil {
		return sig, nil
	}
	
	// Try base64
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return solana.Signature{}, err
	}
	
	if len(bytes) != 64 {
		return solana.Signature{}, errors.New("invalid signature length")
	}
	
	var sigBytes [64]byte
	copy(sigBytes[:], bytes)
	return solana.Signature(sigBytes), nil
}

// LamportsToSOL converts lamports to SOL
func LamportsToSOL(lamports uint64) float64 {
	return float64(lamports) / 1e9
}

// SOLToLamports converts SOL to lamports
func SOLToLamports(sol float64) uint64 {
	return uint64(sol * 1e9)
}

