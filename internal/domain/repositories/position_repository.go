package repositories

import (
	"context"
	"github.com/polymarket/solana-program/internal/domain/entities"
)

// PositionRepository defines the interface for position data operations
type PositionRepository interface {
	Create(ctx context.Context, position *entities.Position) error
	GetByID(ctx context.Context, id string) (*entities.Position, error)
	GetByMarketID(ctx context.Context, marketID string) ([]*entities.Position, error)
	GetByUserID(ctx context.Context, userID string) ([]*entities.Position, error)
	Update(ctx context.Context, position *entities.Position) error
}

