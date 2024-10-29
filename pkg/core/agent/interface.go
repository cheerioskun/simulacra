package agent

import (
	"context"
	"simulacra/pkg/core/action"
	"time"
)

// Memory represents a single memory entry
type Memory struct {
	ID        string
	Content   string
	Timestamp time.Time
	Embedding []float32
	Metadata  map[string]interface{}
}

// Thought represents an agent's thought process
type Thought struct {
	Content   string
	Type      string // "fast" or "slow"
	Timestamp time.Time
}

// Agent defines the core interface for an agent in the system
type Agent interface {
	// Core identity and state
	GetID() string
	GetName() string
	GetState() map[string]interface{}

	// Thought processes
	Think(ctx context.Context) error
	DecideAction(ctx context.Context) (action.Action, error)
	ReceiveOutcome(ctx context.Context, action action.Action, outcome string) error

	// Plugin system hooks
	RegisterPlugin(plugin AgentPlugin) error
	GetPlugins() []AgentPlugin

	// Interaction capabilities
	Interact(ctx context.Context, target Agent, action action.Action) error
	ReceiveInteraction(ctx context.Context, source Agent, action action.Action) error
}
