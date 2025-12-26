package entities

import (
	"time"
)

// Market represents a prediction market entity
type Market struct {
	ID          string
	Title       string
	Description string
	Category    string
	EndDate     time.Time
	Resolution  MarketResolution
	Status      MarketStatus
	Creator     string // Public key of the creator
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// MarketResolution represents how a market is resolved
type MarketResolution string

const (
	ResolutionPending   MarketResolution = "pending"
	ResolutionYes       MarketResolution = "yes"
	ResolutionNo        MarketResolution = "no"
	ResolutionCancelled MarketResolution = "cancelled"
)

// MarketStatus represents the current status of a market
type MarketStatus string

const (
	StatusOpen      MarketStatus = "open"
	StatusClosed    MarketStatus = "closed"
	StatusResolved  MarketStatus = "resolved"
	StatusCancelled MarketStatus = "cancelled"
)

// StatusToUint8 converts MarketStatus to uint8
func (m *Market) StatusToUint8() uint8 {
	switch m.Status {
	case StatusOpen:
		return 0
	case StatusClosed:
		return 1
	case StatusResolved:
		return 2
	case StatusCancelled:
		return 3
	default:
		return 0
	}
}

// Uint8ToStatus converts uint8 to MarketStatus
func Uint8ToStatus(status uint8) MarketStatus {
	switch status {
	case 0:
		return StatusOpen
	case 1:
		return StatusClosed
	case 2:
		return StatusResolved
	case 3:
		return StatusCancelled
	default:
		return StatusOpen
	}
}

// ResolutionToUint8 converts MarketResolution to uint8
func (m *Market) ResolutionToUint8() uint8 {
	switch m.Resolution {
	case ResolutionPending:
		return 0
	case ResolutionYes:
		return 1
	case ResolutionNo:
		return 2
	case ResolutionCancelled:
		return 3
	default:
		return 0
	}
}

// Uint8ToResolution converts uint8 to MarketResolution
func Uint8ToResolution(resolution uint8) MarketResolution {
	switch resolution {
	case 0:
		return ResolutionPending
	case 1:
		return ResolutionYes
	case 2:
		return ResolutionNo
	case 3:
		return ResolutionCancelled
	default:
		return ResolutionPending
	}
}

