package storage

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

type ZerologGormAdapter struct {
	logger zerolog.Logger
}

// Compile time interface implementation enforcement
var _ logger.Interface = &ZerologGormAdapter{}

// Not worth testing as just a wrapper for putting into a struct
func NewGormLogger(zerologger zerolog.Logger) *ZerologGormAdapter {
	return &ZerologGormAdapter{zerologger}
}

// Not testable since it focuses entirely on the side effect of
// configuring loglevel
func (g *ZerologGormAdapter) LogMode(newLevel logger.LogLevel) logger.Interface {
	switch newLevel {
	case logger.Error:
		g.logger = g.logger.Level(zerolog.ErrorLevel)
	case logger.Warn:
		g.logger = g.logger.Level(zerolog.WarnLevel)
	case logger.Info:
		g.logger = g.logger.Level(zerolog.InfoLevel)
	case logger.Silent:
		g.logger = g.logger.Level(zerolog.Disabled)
	}
	return g
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) Info(ctx context.Context, format string, args ...any) {
	g.logger.Info().Ctx(ctx).Msgf(format, args...)
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) Warn(ctx context.Context, format string, args ...any) {
	g.logger.Warn().Ctx(ctx).Msgf(format, args...)
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) Error(ctx context.Context, format string, args ...any) {
	g.logger.Error().Ctx(ctx).Msgf(format, args...)
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	sql, rowsAffected := fc()
	g.logger.Trace().
		Ctx(ctx).
		Time("gorm-begin", begin).
		Err(err).
		Str("gorm-query", sql).
		Int64("gorm-rows-affected", rowsAffected).
		Send()
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) OverwriteLoggingLevel(new zerolog.Level) {
	g.logger = g.logger.Level(new)
}

// Not worth testing since only a wrapper around another function call
func (g *ZerologGormAdapter) OverwriteLogger(new zerolog.Logger) {
	g.logger = new
}
