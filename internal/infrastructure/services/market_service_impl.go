package services

import (
	"context"
	"errors"
	"github.com/polymarket/solana-program/internal/domain/entities"
	"github.com/polymarket/solana-program/internal/domain/repositories"
	"github.com/polymarket/solana-program/internal/domain/services"
)

// MarketServiceImpl implements MarketService
type MarketServiceImpl struct {
	marketRepo repositories.MarketRepository
	validator  *services.MarketValidator
}

// NewMarketServiceImpl creates a new MarketServiceImpl
func NewMarketServiceImpl(marketRepo repositories.MarketRepository) services.MarketService {
	return &MarketServiceImpl{
		marketRepo: marketRepo,
		validator:  &services.MarketValidator{},
	}
}

// CreateMarket creates a new market
func (s *MarketServiceImpl) CreateMarket(ctx context.Context, market *entities.Market) error {
	if err := s.validator.ValidateMarket(ctx, market); err != nil {
		return err
	}
	return s.marketRepo.Create(ctx, market)
}

// ResolveMarket resolves a market
func (s *MarketServiceImpl) ResolveMarket(ctx context.Context, marketID string, resolution entities.MarketResolution, resolver string) error {
	market, err := s.marketRepo.GetByID(ctx, marketID)
	if err != nil {
		return err
	}

	if market == nil {
		return services.ErrMarketNotFound
	}

	if market.Status != entities.StatusOpen && market.Status != entities.StatusClosed {
		return services.ErrInvalidMarketStatus
	}

	market.Resolution = resolution
	market.Status = entities.StatusResolved

	return s.marketRepo.Update(ctx, market)
}

// CloseMarket closes a market
func (s *MarketServiceImpl) CloseMarket(ctx context.Context, marketID string) error {
	market, err := s.marketRepo.GetByID(ctx, marketID)
	if err != nil {
		return err
	}

	if market == nil {
		return services.ErrMarketNotFound
	}

	market.Status = entities.StatusClosed
	return s.marketRepo.Update(ctx, market)
}

// ValidateMarket validates market business rules
func (s *MarketServiceImpl) ValidateMarket(ctx context.Context, market *entities.Market) error {
	return s.validator.ValidateMarket(ctx, market)
}

