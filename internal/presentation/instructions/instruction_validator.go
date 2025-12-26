package instructions

import (
	"errors"
	"github.com/gagliardetto/solana-go"
	"github.com/polymarket/solana-program/internal/infrastructure/solana"
)

// InstructionValidator validates Solana instructions
type InstructionValidator struct {
	accountValidator *solana.AccountValidator
}

// NewInstructionValidator creates a new InstructionValidator
func NewInstructionValidator(accountValidator *solana.AccountValidator) *InstructionValidator {
	return &InstructionValidator{
		accountValidator: accountValidator,
	}
}

// ValidateCreateMarket validates create market instruction
func (iv *InstructionValidator) ValidateCreateMarket(
	accounts []*solana.AccountMeta,
	creator solana.PublicKey,
	marketPDA solana.PublicKey,
) error {
	if len(accounts) < 3 {
		return ErrInvalidAccounts
	}
	
	// Validate creator is signer
	signers := make([]solana.PublicKey, 0)
	for _, acc := range accounts {
		if acc.IsSigner {
			signers = append(signers, acc.PublicKey)
		}
	}
	
	if err := iv.accountValidator.ValidateSigner(creator, signers); err != nil {
		return err
	}
	
	// Validate market PDA is writable
	for _, acc := range accounts {
		if acc.PublicKey.Equals(marketPDA) {
			if !acc.IsWritable {
				return solana.ErrAccountNotWritable
			}
			break
		}
	}
	
	return nil
}

// ValidateResolveMarket validates resolve market instruction
func (iv *InstructionValidator) ValidateResolveMarket(
	accounts []*solana.AccountMeta,
	resolver solana.PublicKey,
	marketPDA solana.PublicKey,
) error {
	if len(accounts) < 2 {
		return ErrInvalidAccounts
	}
	
	// Validate resolver is signer
	signers := make([]solana.PublicKey, 0)
	for _, acc := range accounts {
		if acc.IsSigner {
			signers = append(signers, acc.PublicKey)
		}
	}
	
	return iv.accountValidator.ValidateSigner(resolver, signers)
}

// ValidateCreatePosition validates create position instruction
func (iv *InstructionValidator) ValidateCreatePosition(
	accounts []*solana.AccountMeta,
	user solana.PublicKey,
	positionPDA solana.PublicKey,
) error {
	if len(accounts) < 3 {
		return ErrInvalidAccounts
	}
	
	// Validate user is signer
	signers := make([]solana.PublicKey, 0)
	for _, acc := range accounts {
		if acc.IsSigner {
			signers = append(signers, acc.PublicKey)
		}
	}
	
	if err := iv.accountValidator.ValidateSigner(user, signers); err != nil {
		return err
	}
	
	// Validate position PDA is writable
	for _, acc := range accounts {
		if acc.PublicKey.Equals(positionPDA) {
			if !acc.IsWritable {
				return solana.ErrAccountNotWritable
			}
			break
		}
	}
	
	return nil
}

