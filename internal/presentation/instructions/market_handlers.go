package instructions

import (
	"context"
	"encoding/binary"
	"time"
	"github.com/gagliardetto/solana-go"
	"github.com/polymarket/solana-program/internal/application/usecases"
	"github.com/polymarket/solana-program/internal/domain/entities"
)

// handleCreateMarket handles the create market instruction
func (h *InstructionHandler) handleCreateMarket(ctx context.Context, data []byte, accounts []*solana.AccountMeta) error {
	if len(accounts) < 3 {
		return ErrInvalidAccounts
	}

	// Parse instruction data
	// Format: [title_len(4)][title][desc_len(4)][desc][category_len(4)][category][end_date(8)]
	offset := 0

	// Read title
	titleLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	title := string(data[offset : offset+int(titleLen)])
	offset += int(titleLen)

	// Read description
	descLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	description := string(data[offset : offset+int(descLen)])
	offset += int(descLen)

	// Read category
	categoryLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	category := string(data[offset : offset+int(categoryLen)])
	offset += int(categoryLen)

	// Read end date
	endDateUnix := binary.LittleEndian.Uint64(data[offset:])
	endDate := time.Unix(int64(endDateUnix), 0)

	// Get creator from accounts
	creator := accounts[0].PublicKey.String()

	// Create market input
	input := usecases.CreateMarketInput{
		Title:       title,
		Description: description,
		Category:    category,
		EndDate:     endDate,
		Creator:     creator,
	}

	_, err := h.createMarketUseCase.Execute(ctx, input)
	return err
}

// handleResolveMarket handles the resolve market instruction
func (h *InstructionHandler) handleResolveMarket(ctx context.Context, data []byte, accounts []*solana.AccountMeta) error {
	if len(data) < 2 {
		return ErrInvalidInstructionData
	}

	// Parse instruction data
	// Format: [market_id_len(4)][market_id][resolution(1)]
	offset := 0

	marketIDLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	marketID := string(data[offset : offset+int(marketIDLen)])
	offset += int(marketIDLen)

	resolution := entities.MarketResolution(data[offset])
	resolver := accounts[0].PublicKey.String()

	input := usecases.ResolveMarketInput{
		MarketID:   marketID,
		Resolution: resolution,
		Resolver:   resolver,
	}

	return h.resolveMarketUseCase.Execute(ctx, input)
}

// handleCloseMarket handles the close market instruction
func (h *InstructionHandler) handleCloseMarket(ctx context.Context, data []byte, accounts []*solana.AccountMeta) error {
	if len(data) < 5 || len(accounts) < 1 {
		return ErrInvalidInstructionData
	}

	// Parse instruction data
	// Format: [market_id_len(4)][market_id]
	offset := 0

	marketIDLen := binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	marketID := string(data[offset : offset+int(marketIDLen)])

	closer := accounts[0].PublicKey.String()

	input := usecases.CloseMarketInput{
		MarketID: marketID,
		Closer:   closer,
	}

	return h.closeMarketUseCase.Execute(ctx, input)
}

var (
	ErrInvalidAccounts      = &InstructionError{Message: "invalid accounts"}
	ErrInvalidInstructionData = &InstructionError{Message: "invalid instruction data"}
)

