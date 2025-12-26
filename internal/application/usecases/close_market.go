package usecases

import (
	"context"
	"github.com/polymarket/solana-program/internal/domain/repositories"
	"github.com/polymarket/solana-program/internal/domain/services"
)

// CloseMarketUseCase handles market closing
type CloseMarketUseCase struct {
	marketRepo   repositories.MarketRepository
	marketService services.MarketService
}

// NewCloseMarketUseCase creates a new CloseMarketUseCase
func NewCloseMarketUseCase(
	marketRepo repositories.MarketRepository,
	marketService services.MarketService,
) *CloseMarketUseCase {
	return &CloseMarketUseCase{
		marketRepo:   marketRepo,
		marketService: marketService,
	}
}

// CloseMarketInput represents the input for closing a market
type CloseMarketInput struct {
	MarketID string
	Closer   string
}

// Execute closes a market
func (uc *CloseMarketUseCase) Execute(ctx context.Context, input CloseMarketInput) error {
	market, err := uc.marketRepo.GetByID(ctx, input.MarketID)
	if err != nil {
		return err
	}

	// Check if closer is authorized
	if market.Creator != input.Closer {
		return services.ErrUnauthorized
	}

	return uc.marketService.CloseMarket(ctx, input.MarketID)
}

