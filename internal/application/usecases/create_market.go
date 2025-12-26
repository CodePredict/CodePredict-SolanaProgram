package usecases

import (
	"context"
	"time"
	"github.com/polymarket/solana-program/internal/domain/entities"
	"github.com/polymarket/solana-program/internal/domain/repositories"
	"github.com/polymarket/solana-program/internal/domain/services"
)

// CreateMarketUseCase handles market creation
type CreateMarketUseCase struct {
	marketRepo   repositories.MarketRepository
	marketService services.MarketService
}

// NewCreateMarketUseCase creates a new CreateMarketUseCase
func NewCreateMarketUseCase(
	marketRepo repositories.MarketRepository,
	marketService services.MarketService,
) *CreateMarketUseCase {
	return &CreateMarketUseCase{
		marketRepo:   marketRepo,
		marketService: marketService,
	}
}

// CreateMarketInput represents the input for creating a market
type CreateMarketInput struct {
	Title       string
	Description string
	Category    string
	EndDate     time.Time
	Creator     string
}

// Execute creates a new market
func (uc *CreateMarketUseCase) Execute(ctx context.Context, input CreateMarketInput) (*entities.Market, error) {
	market := &entities.Market{
		ID:          generateMarketID(),
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		EndDate:     input.EndDate,
		Resolution:  entities.ResolutionPending,
		Status:      entities.StatusOpen,
		Creator:     input.Creator,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.marketService.ValidateMarket(ctx, market); err != nil {
		return nil, err
	}

	if err := uc.marketRepo.Create(ctx, market); err != nil {
		return nil, err
	}

	return market, nil
}

func generateMarketID() string {
	// In a real implementation, this would generate a unique ID
	// For Solana, this could be a PDA (Program Derived Address)
	return "market_" + time.Now().Format("20060102150405")
}

