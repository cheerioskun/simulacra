package factory

import (
	"fmt"

	"simulacra/pkg/llm"
	"simulacra/pkg/llm/openrouter"
)

// Config holds the configuration for LLM providers
type Config struct {
	Provider string
	APIKey   string
}

// New creates a new LLM provider based on the configuration
func New(config Config) (llm.Provider, error) {
	switch config.Provider {
	case "openrouter":
		return openrouter.New(config.APIKey), nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", config.Provider)
	}
}
