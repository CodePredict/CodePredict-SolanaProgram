package instructions

import (
	"context"
	"encoding/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/polymarket/solana-program/internal/application/usecases"
	"github.com/polymarket/solana-program/internal/domain/entities"
)

// handleCreatePosition handles the create position instruction
func (h *InstructionHandler) handleCreatePosition(ctx context.Context, data []byte, accounts []*solana.AccountMeta) error {
	if len(data) < 25 || len(accounts) < 2 {
		return ErrInvalidInstructionData
	}

	// Parse instruction data
	// Format: [market_id_len(4)][market_id][side(1)][amount(8)][price(8)]
	offset := 0

	marketIDLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	marketID := string(data[offset : offset+int(marketIDLen)])
	offset += int(marketIDLen)

	side := entities.PositionSide(data[offset])
	offset++

	amount := binary.LittleEndian.Uint64(data[offset:])
	offset += 8

	price := binary.LittleEndian.Uint64(data[offset:])

	userID := accounts[0].PublicKey.String()

	input := usecases.CreatePositionInput{
		MarketID: marketID,
		UserID:   userID,
		Side:     side,
		Amount:   amount,
		Price:    price,
	}

	_, err := h.createPositionUseCase.Execute(ctx, input)
	return err
}

