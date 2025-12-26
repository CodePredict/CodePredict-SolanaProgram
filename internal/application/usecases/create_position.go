package usecases

import (
	"context"
	"time"

	"github.com/polymarket/solana-program/internal/domain/entities"
	"github.com/polymarket/solana-program/internal/domain/repositories"
	"github.com/polymarket/solana-program/internal/domain/services"
)

// CreatePositionUseCase handles position creation
type CreatePositionUseCase struct {
	positionRepo repositories.PositionRepository
	marketRepo   repositories.MarketRepository
}

// NewCreatePositionUseCase creates a new CreatePositionUseCase
func NewCreatePositionUseCase(
	positionRepo repositories.PositionRepository,
	marketRepo repositories.MarketRepository,
) *CreatePositionUseCase {
	return &CreatePositionUseCase{
		positionRepo: positionRepo,
		marketRepo:   marketRepo,
	}
}

// CreatePositionInput represents the input for creating a position
type CreatePositionInput struct {
	MarketID string
	UserID   string
	Side     entities.PositionSide
	Amount   uint64
	Price    uint64
}

// Execute creates a new position
func (uc *CreatePositionUseCase) Execute(ctx context.Context, input CreatePositionInput) (*entities.Position, error) {
	// Validate market exists and is open
	market, err := uc.marketRepo.GetByID(ctx, input.MarketID)
	if err != nil {
		return nil, err
	}

	if market.Status != entities.StatusOpen {
		return nil, services.ErrMarketClosed
	}

	position := &entities.Position{
		ID:        generatePositionID(),
		MarketID:  input.MarketID,
		UserID:    input.UserID,
		Side:      input.Side,
		Amount:    input.Amount,
		Price:     input.Price,
		CreatedAt: time.Now(),
	}

	if err := uc.positionRepo.Create(ctx, position); err != nil {
		return nil, err
	}

	return position, nil
}

func generatePositionID() string {
	return "position_" + time.Now().Format("20060102150405")
}
