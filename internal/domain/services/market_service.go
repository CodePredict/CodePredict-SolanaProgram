package services

import (
	"context"
	"errors"
	"time"
	"github.com/polymarket/solana-program/internal/domain/entities"
)

var (
	ErrMarketNotFound      = errors.New("market not found")
	ErrInvalidMarketStatus = errors.New("invalid market status")
	ErrMarketClosed        = errors.New("market is closed")
	ErrUnauthorized        = errors.New("unauthorized")
)

// MarketService defines business logic for markets
type MarketService interface {
	CreateMarket(ctx context.Context, market *entities.Market) error
	ResolveMarket(ctx context.Context, marketID string, resolution entities.MarketResolution, resolver string) error
	CloseMarket(ctx context.Context, marketID string) error
	ValidateMarket(ctx context.Context, market *entities.Market) error
}

// MarketValidator validates market business rules
type MarketValidator struct{}

// ValidateMarket validates market creation rules
func (v *MarketValidator) ValidateMarket(ctx context.Context, market *entities.Market) error {
	if market.Title == "" {
		return errors.New("market title is required")
	}
	if market.EndDate.Before(time.Now()) {
		return errors.New("market end date must be in the future")
	}
	if market.Creator == "" {
		return errors.New("market creator is required")
	}
	return nil
}

