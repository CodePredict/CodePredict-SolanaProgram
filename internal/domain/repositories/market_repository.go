package repositories

import (
	"context"
	"github.com/polymarket/solana-program/internal/domain/entities"
)

// MarketRepository defines the interface for market data operations
type MarketRepository interface {
	Create(ctx context.Context, market *entities.Market) error
	GetByID(ctx context.Context, id string) (*entities.Market, error)
	Update(ctx context.Context, market *entities.Market) error
	GetAll(ctx context.Context) ([]*entities.Market, error)
	GetByCreator(ctx context.Context, creator string) ([]*entities.Market, error)
}

