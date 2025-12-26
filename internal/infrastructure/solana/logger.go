package solana

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger handles logging for Solana operations using zap
type Logger struct {
	logger  *zap.Logger
	sugar   *zap.SugaredLogger
	enabled bool
}

// NewLogger creates a new Logger with zap
func NewLogger(enabled bool) *Logger {
	if !enabled {
		nop := zap.NewNop()
		return &Logger{
			logger:  nop,
			sugar:   nop.Sugar(),
			enabled: false,
		}
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.FunctionKey = "function"

	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		// Fallback to production logger if development config fails
		logger, _ = zap.NewProduction()
	}

	return &Logger{
		logger:  logger.With(zap.String("component", "solana")),
		sugar:   logger.With(zap.String("component", "solana")).Sugar(),
		enabled: true,
	}
}

// NewProductionLogger creates a new Logger with production configuration
func NewProductionLogger() *Logger {
	logger, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		logger = zap.NewNop()
	}

	return &Logger{
		logger:  logger.With(zap.String("component", "solana")),
		sugar:   logger.With(zap.String("component", "solana")).Sugar(),
		enabled: true,
	}
}

// NewDevelopmentLogger creates a new Logger with development configuration
func NewDevelopmentLogger() *Logger {
	logger, err := zap.NewDevelopment(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		logger = zap.NewNop()
	}

	return &Logger{
		logger:  logger.With(zap.String("component", "solana")),
		sugar:   logger.With(zap.String("component", "solana")).Sugar(),
		enabled: true,
	}
}

// LogInstruction logs instruction processing
func (l *Logger) LogInstruction(instructionType string, accounts []string) {
	if !l.enabled {
		return
	}
	l.sugar.Infow("Processing instruction",
		"instruction_type", instructionType,
		"accounts", accounts,
	)
}

// LogTransaction logs transaction
func (l *Logger) LogTransaction(signature string, status string) {
	if !l.enabled {
		return
	}
	l.sugar.Infow("Transaction",
		"signature", signature,
		"status", status,
	)
}

// LogAccount logs account operation
func (l *Logger) LogAccount(operation string, publicKey string) {
	if !l.enabled {
		return
	}
	l.sugar.Infow("Account operation",
		"operation", operation,
		"public_key", publicKey,
	)
}

// LogError logs error
func (l *Logger) LogError(operation string, err error) {
	if !l.enabled || err == nil {
		return
	}
	l.sugar.Errorw("Error in operation",
		"operation", operation,
		"error", err,
	)
}

// Info logs info level message
func (l *Logger) Info(msg string, fields ...zap.Field) {
	if !l.enabled {
		return
	}
	l.logger.Info(msg, fields...)
}

// Error logs error level message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	if !l.enabled {
		return
	}
	l.logger.Error(msg, fields...)
}

// Debug logs debug level message
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if !l.enabled {
		return
	}
	l.logger.Debug(msg, fields...)
}

// Warn logs warn level message
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if !l.enabled {
		return
	}
	l.logger.Warn(msg, fields...)
}

// With creates a child logger with additional fields
func (l *Logger) With(fields ...zap.Field) *Logger {
	if !l.enabled {
		return l
	}
	return &Logger{
		logger:  l.logger.With(fields...),
		sugar:   l.logger.With(fields...).Sugar(),
		enabled: l.enabled,
	}
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	if !l.enabled {
		return nil
	}
	return l.logger.Sync()
}

// GetLogger returns the underlying zap logger
func (l *Logger) GetLogger() *zap.Logger {
	return l.logger
}

// GetSugaredLogger returns the sugared zap logger
func (l *Logger) GetSugaredLogger() *zap.SugaredLogger {
	return l.sugar
}
