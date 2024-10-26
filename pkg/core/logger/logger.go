package logger

import (
	"log/slog"
	"os"
)

// Categories for our simulation
const (
	CategoryAgent    = "agent"
	CategoryWorld    = "world"
	CategoryPlugin   = "plugin"
	CategorySystem   = "system"
	CategoryResearch = "research"
	CategoryPerf     = "performance"
)

// SetupLogger configures slog for our simulation
func SetupLogger(isDevelopment bool) *slog.Logger {
	var handler slog.Handler

	if isDevelopment {
		// Pretty text output for development
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		// JSON output for production
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	return slog.New(handler)
}

// WithAgentContext adds common agent fields to the logger
func WithAgentContext(logger *slog.Logger, agentID string) *slog.Logger {
	return logger.With(
		slog.String("category", CategoryAgent),
		slog.String("agent_id", agentID),
	)
}
