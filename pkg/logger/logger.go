package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a global logger interface
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	With(fields ...zap.Field) Logger
	Sync() error
}

// NewProduction creates a new production logger
func NewProduction() (Logger, error) {
	logger, err := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}
	return &zapLogger{logger: logger}, nil
}

// NewDevelopment creates a new development logger
func NewDevelopment() (Logger, error) {
	logger, err := zap.NewDevelopment(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}
	return &zapLogger{logger: logger}, nil
}

// NewNop creates a no-op logger
func NewNop() Logger {
	return &zapLogger{logger: zap.NewNop()}
}

// zapLogger wraps zap.Logger to implement Logger interface
type zapLogger struct {
	logger *zap.Logger
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) With(fields ...zap.Field) Logger {
	return &zapLogger{logger: l.logger.With(fields...)}
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}

