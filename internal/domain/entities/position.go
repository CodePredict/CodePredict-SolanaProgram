package entities

import "time"

// Position represents a user's position in a market
type Position struct {
	ID        string
	MarketID  string
	UserID    string // Public key of the user
	Side      PositionSide
	Amount    uint64 // Amount in lamports
	Price     uint64 // Price per share in lamports
	CreatedAt time.Time
}

// PositionSide represents whether the position is YES or NO
type PositionSide string

const (
	SideYes PositionSide = "yes"
	SideNo  PositionSide = "no"
)

