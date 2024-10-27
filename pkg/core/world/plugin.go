package world

import (
	"context"
	"simulacra/pkg/core/agent"
)

// WorldPlugin defines the interface for world plugins
type WorldPlugin interface {
	// Plugin identity
	GetID() string
	GetName() string
	GetDescription() string

	// Lifecycle hooks
	OnLoad(world World) error
	OnUnload() error

	// Core hooks
	PreUpdate(ctx context.Context) error
	PostUpdate(ctx context.Context) error
	OnAgentAdded(ctx context.Context, agent agent.Agent) error
	OnAgentRemoved(ctx context.Context, agentID string) error
}
