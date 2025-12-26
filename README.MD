# Solana Program for Polymarket

A prediction market on Solana blockchain, implemented using clean architecture in Go.

## Architecture

The project follows clean architecture principles and is divided into the following layers:

### Domain Layer (`internal/domain/`)
- **entities/**: Business entities (Market, Position, Account)
- **repositories/**: Repository interfaces
- **services/**: Business logic interfaces

### Application Layer (`internal/application/`)
- **usecases/**: Use cases (CreateMarket, ResolveMarket, CreatePosition, CloseMarket)

### Infrastructure Layer (`internal/infrastructure/`)
- **solana/**: Full Solana integration:
  - `program.go` - Solana program
  - `account_manager.go` - Account management
  - `account_validator.go` - Account validation
  - `borsh_serializer.go` - Borsh serialization/deserialization
  - `instruction_builder.go` - Instruction building
  - `transaction_handler.go` - Transaction handling
  - `pda_manager.go` - PDA (Program Derived Addresses) management
  - `config.go` - Configuration
  - `rent_calculator.go` - Rent calculation
  - `error_handler.go` - Error handling
  - `logger.go` - Logging
- **repositories/**: Solana repository implementations:
  - `solana_market_repository.go` - Market repository
  - `solana_position_repository.go` - Position repository
  - `solana_account_repository.go` - Account repository
  - `solana_market_index_repository.go` - Market indexing
  - `solana_position_index_repository.go` - Position indexing
- **services/**: Service implementations

### Presentation Layer (`internal/presentation/`)
- **instructions/**: Solana instruction handlers:
  - `instruction_handler.go` - Main handler
  - `market_handlers.go` - Market handlers
  - `position_handlers.go` - Position handlers
  - `instruction_validator.go` - Instruction validation

### Shared (`pkg/`)
- **errors/**: Common error types
- **utils/**: Utilities
- **solana/**: Solana utilities (key conversion, lamports, etc.)

## Project Structure

```
.
├── cmd/
│   └── program/
│       └── main.go              # Program entry point
├── internal/
│   ├── domain/                  # Domain layer
│   │   ├── entities/            # Entities
│   │   ├── repositories/        # Repository interfaces
│   │   └── services/            # Service interfaces
│   ├── application/             # Application layer
│   │   └── usecases/            # Use cases
│   ├── infrastructure/          # Infrastructure layer
│   │   ├── solana/              # Solana integration (10+ files)
│   │   ├── repositories/        # Repository implementations (5 files)
│   │   └── services/            # Service implementations
│   └── presentation/            # Presentation layer
│       └── instructions/        # Instruction handlers (4 files)
├── pkg/                         # Shared packages
│   ├── errors/                  # Error handling
│   ├── utils/                   # Utilities
│   └── solana/                  # Solana utilities
├── go.mod                       # Dependencies
├── Makefile                     # Build commands
└── README.MD                    # Documentation
```

## Main Entities

### Market
Represents a prediction market with the ability to create, close, and resolve.

### Position
Represents a user's position on a market (YES or NO).

## Solana Integrations

### PDA (Program Derived Addresses)
- PDA management for markets, positions, and indexes
- Automatic address calculation

### Borsh Serialization
- Full account serialization/deserialization
- MarketAccount and PositionAccount support

### Validation
- Account validation (owner, size, permissions)
- PDA validation
- Instruction validation

### Transactions
- Instruction building
- Transaction sending
- Status tracking

### Indexing
- Market indexes
- User position indexes
- Market position indexes

## Solana Instructions

1. **CreateMarket**: Create a new market
2. **ResolveMarket**: Resolve a market (Yes/No)
3. **CreatePosition**: Create a position on a market
4. **CloseMarket**: Close a market

## Installation and Setup

### Requirements
- Go 1.21 or higher
- Solana CLI (for deployment)

### Install Dependencies
```bash
make deps
# or
go mod tidy
```

### Build
```bash
make build
```

### Test
```bash
make test
```

### Format Code
```bash
make fmt
```

### Lint
```bash
make lint
```

## Development

### Adding New Functionality

1. **Domain Layer**: Define entities and interfaces in `internal/domain/`
2. **Application Layer**: Create a use case in `internal/application/usecases/`
3. **Infrastructure**: Implement repositories in `internal/infrastructure/repositories/`
4. **Presentation**: Add an instruction handler in `internal/presentation/instructions/`

## Solana Components

### AccountManager
Solana account management, PDA operations.

### BorshSerializer
Data serialization and deserialization in Borsh format.

### AccountValidator
Account validation, permission checking.

### InstructionBuilder
Building Solana instructions for various operations.

### TransactionHandler
Transaction sending and tracking.

### PDAManager
Program Derived Addresses management for all entities.

### RentCalculator
Rent calculation for accounts.

## Notes

- The program uses Program Derived Addresses (PDA) for data storage
- Data serialization is performed in Borsh format (Solana standard)
- Full integration with Solana RPC for blockchain operations
- Indexing for efficient market and position search
- Validation of all operations at the instruction level

## License

MIT
