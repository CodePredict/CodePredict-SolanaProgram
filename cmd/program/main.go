package main

import (
	"context"
	"log"
	solanago "github.com/gagliardetto/solana-go"
	"github.com/polymarket/solana-program/internal/application/usecases"
	"github.com/polymarket/solana-program/internal/infrastructure/repositories"
	"github.com/polymarket/solana-program/internal/infrastructure/services"
	"github.com/polymarket/solana-program/internal/infrastructure/solana"
	"github.com/polymarket/solana-program/internal/presentation/instructions"
)

func main() {
	// Initialize program ID (in production, this would be the deployed program ID)
	programID := solanago.MustPublicKeyFromBase58("11111111111111111111111111111111") // Placeholder

	// Initialize infrastructure
	program := solana.NewProgram(programID)
	accountManager := solana.NewAccountManager(program)

	// Initialize repositories
	marketRepo := repositories.NewSolanaMarketRepository(accountManager, program)
	positionRepo := repositories.NewSolanaPositionRepository(accountManager, program)

	// Initialize services
	marketService := services.NewMarketServiceImpl(marketRepo)

	// Initialize use cases
	createMarketUseCase := usecases.NewCreateMarketUseCase(marketRepo, marketService)
	resolveMarketUseCase := usecases.NewResolveMarketUseCase(marketRepo, marketService)
	createPositionUseCase := usecases.NewCreatePositionUseCase(positionRepo, marketRepo)
	closeMarketUseCase := usecases.NewCloseMarketUseCase(marketRepo, marketService)

	// Initialize instruction handler
	instructionHandler := instructions.NewInstructionHandler(
		createMarketUseCase,
		resolveMarketUseCase,
		createPositionUseCase,
		closeMarketUseCase,
	)

	// This is where the Solana program entry point would be
	// In a real Solana program, this would be called by the Solana runtime
	ctx := context.Background()

	// Example: Process an instruction
	// In production, this would receive instruction data from Solana runtime
	_ = instructionHandler

	log.Println("Solana program initialized")
	log.Printf("Program ID: %s\n", programID.String())

	// Keep the program running (in production, Solana runtime handles this)
	select {}
}

