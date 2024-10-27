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
	Content     string
	Type        string // "fast" or "slow"
	Timestamp   time.Time
	RelatedMems []string // IDs of related memories
}

// Agent defines the core interface for an agent in the system
type Agent interface {
	// Core identity and state
	GetID() string
	GetName() string
	GetState() map[string]interface{}

	// Memory operations
	AddMemory(ctx context.Context, mem Memory) error
	RetrieveMemories(ctx context.Context, query string, k int) ([]Memory, error)

	// Thought processes
	Think(ctx context.Context, input string) (*Thought, error)
	DecideAction(ctx context.Context) (action.Action, error)

	// Plugin system hooks
	RegisterPlugin(plugin AgentPlugin) error
	GetPlugins() []AgentPlugin

	// Interaction capabilities
	Interact(ctx context.Context, target Agent, action action.Action) error
	ReceiveInteraction(ctx context.Context, source Agent, action action.Action) error
}
