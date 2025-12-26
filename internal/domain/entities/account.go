package entities

import (
	"github.com/gagliardetto/solana-go"
)

// Account represents a Solana account
type Account struct {
	PublicKey  solana.PublicKey
	Data       []byte
	Owner      solana.PublicKey
	Lamports   uint64
	Executable bool
}

// MarketAccount represents the on-chain state of a market
type MarketAccount struct {
	MarketID   string
	Title      string
	EndDate    int64
	Status     uint8
	Resolution uint8
	Creator    [32]byte
}

// PositionAccount represents the on-chain state of a position
type PositionAccount struct {
	MarketID string
	UserID   [32]byte
	Side     uint8
	Amount   uint64
	Price    uint64
}

