package solana

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// ErrorHandler handles Solana-specific errors
type ErrorHandler struct{}

// NewErrorHandler creates a new ErrorHandler
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// HandleRPCError handles RPC errors
func (eh *ErrorHandler) HandleRPCError(err error) error {
	if err == nil {
		return nil
	}

	// In a real implementation, this would parse RPC error codes
	// and return domain-specific errors

	return fmt.Errorf("RPC error: %w", err)
}

// HandleTransactionError handles transaction errors
func (eh *ErrorHandler) HandleTransactionError(signature solana.Signature, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("transaction %s failed: %w", signature.String(), err)
}

// HandleTransactionError handles transaction errors
func (eh *ErrorHandler) HandleWalletConnecttionError(signature solana.Wallet, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("transaction %s failed: %w", signature.String(), err)
}

// IsAccountNotFoundError checks if error is account not found
func (eh *ErrorHandler) IsAccountNotFoundError(err error) bool {
	// In a real implementation, check RPC error code
	return err != nil
}

// ParseRPCError parses RPC error response
func (eh *ErrorHandler) ParseRPCError(rpcErr *rpc.RPCError) error {
	if rpcErr == nil {
		return nil
	}

	// Map RPC error codes to domain errors
	switch rpcErr.Code {
	case rpc.ErrCodeAccountNotFound:
		return ErrAccountNotFound
	case rpc.ErrCodeInvalidAccountData:
		return ErrInvalidAccountData
	default:
		return fmt.Errorf("RPC error %d: %s", rpcErr.Code, rpcErr.Message)
	}
}
