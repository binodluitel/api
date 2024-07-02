package log

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// ContextLoggerKey is the type of context key used to store the logger in the context
type ContextLoggerKey string

// ContextLogger is the key used to store the logger in the context
const ContextLogger ContextLoggerKey = "ContextLogger"

// Logging constants and defaults
const (
	EnvLevel             = "LOG_LEVEL"
	EnvDevelopment       = "LOG_DEVELOPMENT"
	EnvEncoding          = "LOG_ENCODING"
	EnvTracerName        = "LOG_TRACER_NAME"
	SamplingInitial      = 100
	SamplingThereafter   = 100
	UnknownLogTracerName = "unknown-log-tracer"
)

// trace for tracer to play tandem with logger
type trace struct {
	span oteltrace.Span
}

// Logger instance for logging
type Logger struct {
	*zap.Logger
	trace trace
}

// New initializes a new logger
func New(_ context.Context, opts ...zap.Option) (Logger, error) {
	// set env vars defaults to prod config if they are not set
	logConfig := zap.NewProductionConfig()
	if os.Getenv(EnvLevel) == "" {
		if err := os.Setenv(EnvLevel, logConfig.Level.String()); err != nil {
			panic(err.Error())
		}
	}

	if os.Getenv(EnvEncoding) == "" {
		if err := os.Setenv(EnvEncoding, logConfig.Encoding); err != nil {
			panic(err.Error())
		}
	}

	if os.Getenv(EnvTracerName) == "" {
		if err := os.Setenv(EnvTracerName, UnknownLogTracerName); err != nil {
			panic(err.Error())
		}
	}

	level, err := zap.ParseAtomicLevel(os.Getenv(EnvLevel))
	if err != nil {
		return Logger{}, err
	}

	// Start with production encoder and update as required
	encoderConfig := zap.NewProductionEncoderConfig()
	zapConfig := zap.Config{
		Level:       level,
		Development: os.Getenv(EnvDevelopment) == "true",
		Sampling: &zap.SamplingConfig{
			Initial:    SamplingInitial,
			Thereafter: SamplingThereafter,
		},
		Encoding:         os.Getenv(EnvEncoding),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build(opts...)
	if err != nil {
		return Logger{}, err
	}
	return Logger{Logger: logger}, nil
}

// Get returns logger if already initialized or
// creates and returns logger if not initialized
func Get(ctx context.Context, opts ...zap.Option) (context.Context, Logger) {
	if logger, ok := ctx.Value(ContextLogger).(Logger); ok {
		return ctx, logger
	}
	// Create and return new logger
	logger, err := New(ctx, opts...)
	if err != nil {
		panic(err.Error())
	}
	ctx = context.WithValue(ctx, ContextLogger, logger)
	return ctx, logger
}

// WithTrace returns a logger with trace span initialized with an operation name
func WithTrace(ctx context.Context, operation string) (context.Context, Logger) {
	ctx, span := otel.Tracer(os.Getenv(EnvTracerName)).Start(ctx, operation)
	ctx, logger := Get(ctx)
	logger.trace.span = span
	return ctx, logger
}

func (l Logger) Sync() {
	if l.trace.span != nil {
		l.trace.span.End()
	}
	l.Logger.Sync()
}

// With returns a new logger with additional fields
func (l Logger) With(fields ...zap.Field) Logger {
	return Logger{l.Logger.With(fields...), l.trace}
}
