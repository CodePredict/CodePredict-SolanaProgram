package usecases

import (
	"context"
	"github.com/polymarket/solana-program/internal/domain/entities"
	"github.com/polymarket/solana-program/internal/domain/repositories"
	"github.com/polymarket/solana-program/internal/domain/services"
)

// ResolveMarketUseCase handles market resolution
type ResolveMarketUseCase struct {
	marketRepo   repositories.MarketRepository
	marketService services.MarketService
}

// NewResolveMarketUseCase creates a new ResolveMarketUseCase
func NewResolveMarketUseCase(
	marketRepo repositories.MarketRepository,
	marketService services.MarketService,
) *ResolveMarketUseCase {
	return &ResolveMarketUseCase{
		marketRepo:   marketRepo,
		marketService: marketService,
	}
}

// ResolveMarketInput represents the input for resolving a market
type ResolveMarketInput struct {
	MarketID   string
	Resolution entities.MarketResolution
	Resolver   string
}

// Execute resolves a market
func (uc *ResolveMarketUseCase) Execute(ctx context.Context, input ResolveMarketInput) error {
	market, err := uc.marketRepo.GetByID(ctx, input.MarketID)
	if err != nil {
		return err
	}

	// Check if resolver is authorized (could be creator or admin)
	if market.Creator != input.Resolver {
		return services.ErrUnauthorized
	}

	return uc.marketService.ResolveMarket(ctx, input.MarketID, input.Resolution, input.Resolver)
}

