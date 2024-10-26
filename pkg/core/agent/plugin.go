package agent

import (
	"context"
	"llm-simulation/pkg/core/action"
)

// AgentPlugin defines the interface for agent plugins
type AgentPlugin interface {
	// Plugin identity
	GetID() string
	GetName() string
	GetDescription() string

	// Lifecycle hooks
	OnLoad(agent Agent) error
	OnUnload() error

	// Core hooks
	PreThink(ctx context.Context, input string) (string, error)
	PostThink(ctx context.Context, thought *Thought) error
	PreAction(ctx context.Context, action action.Action) error
	PostAction(ctx context.Context, action action.Action) error
}
