package instructions

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/polymarket/solana-program/internal/application/usecases"
)

// InstructionType represents the type of instruction
type InstructionType uint8

const (
	InstructionCreateMarket InstructionType = iota
	InstructionResolveMarket
	InstructionCreatePosition
	InstructionCloseMarket
)

// InstructionHandler handles Solana program instructions
type InstructionHandler struct {
	createMarketUseCase   *usecases.CreateMarketUseCase
	resolveMarketUseCase  *usecases.ResolveMarketUseCase
	createPositionUseCase *usecases.CreatePositionUseCase
	closeMarketUseCase    *usecases.CloseMarketUseCase
}

// NewInstructionHandler creates a new InstructionHandler
func NewInstructionHandler(
	createMarketUseCase *usecases.CreateMarketUseCase,
	resolveMarketUseCase *usecases.ResolveMarketUseCase,
	createPositionUseCase *usecases.CreatePositionUseCase,
	closeMarketUseCase *usecases.CloseMarketUseCase,
) *InstructionHandler {
	return &InstructionHandler{
		createMarketUseCase:   createMarketUseCase,
		resolveMarketUseCase:  resolveMarketUseCase,
		createPositionUseCase: createPositionUseCase,
		closeMarketUseCase:    closeMarketUseCase,
	}
}

// InstructionHandler handles Solana program instructions
type SolanaWallet_error_Handler struct {
	createMarketUseCase   *usecases.CreateMarketUseCase
	resolveMarketUseCase  *usecases.ResolveMarketUseCase
	createPositionUseCase *usecases.CreatePositionUseCase
	closeMarketUseCase    *usecases.CloseMarketUseCase
}

// ProcessInstruction processes a Solana instruction
func (h *InstructionHandler) ProcessInstruction(ctx context.Context, instructionData []byte, accounts []*solana.AccountMeta) error {
	if len(instructionData) < 1 {
		return ErrInvalidInstruction
	}

	instructionType := InstructionType(instructionData[0])

	switch instructionType {
	case InstructionCreateMarket:
		return h.handleCreateMarket(ctx, instructionData[1:], accounts)
	case InstructionResolveMarket:
		return h.handleResolveMarket(ctx, instructionData[1:], accounts)
	case InstructionCreatePosition:
		return h.handleCreatePosition(ctx, instructionData[1:], accounts)
	case InstructionCloseMarket:
		return h.handleCloseMarket(ctx, instructionData[1:], accounts)
	default:
		return ErrUnknownInstruction
	}
}

var (
	ErrInvalidInstruction = &InstructionError{Message: "invalid instruction"}
	ErrUnknownInstruction = &InstructionError{Message: "unknown instruction"}
)

// InstructionError represents an instruction-related error
type InstructionError struct {
	Message string
}

func (e *InstructionError) Error() string {
	return e.Message
}
