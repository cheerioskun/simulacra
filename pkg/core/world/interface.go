package world

import (
	"context"
	"simulacra/pkg/core/action"
	"simulacra/pkg/core/agent"
)

// World defines the interface for the simulation world
type World interface {
	// Core world operations
	Initialize(ctx context.Context, config map[string]interface{}) error
	// Step advances the world by one timestep
	Step(ctx context.Context) error

	// Action system
	ActionChan() <-chan action.Action

	// Agent management
	AddAgent(ctx context.Context, agent agent.Agent) error
	RemoveAgent(ctx context.Context, agentID string) error
	GetAgent(agentID string) (agent.Agent, error)

	// Plugin system
	RegisterPlugin(plugin WorldPlugin) error
	GetPlugins() []WorldPlugin

	// Environment operations
	GetState() map[string]interface{}
	SetState(key string, value interface{}) error
}
