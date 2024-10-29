package agent

import (
	"context"
	"fmt"
	"log/slog"
	"simulacra/pkg/core/action"
	"simulacra/pkg/core/logger"
	"simulacra/pkg/llm"
	"sync"
	"time"
)

type DefaultAgent struct {
	id      string
	name    string
	state   map[string]interface{}
	plugins []AgentPlugin
	llm     llm.Provider
	log     *slog.Logger
	mu      sync.RWMutex
}

type Config struct {
	ID      string
	Name    string
	LLM     llm.Provider
	Logger  *slog.Logger
	Plugins []AgentPlugin
}

func NewDefaultAgent(cfg Config) (*DefaultAgent, error) {
	if cfg.ID == "" {
		return nil, fmt.Errorf("agent ID is required")
	}
	if cfg.Name == "" {
		return nil, fmt.Errorf("agent name is required")
	}
	if cfg.LLM == nil {
		return nil, fmt.Errorf("LLM provider is required")
	}

	log := cfg.Logger
	if log == nil {
		log = slog.Default()
	}
	log = log.With(
		logger.CategoryKey, logger.CategoryAgent,
		"agent_id", cfg.ID,
		"agent_name", cfg.Name,
	)

	return &DefaultAgent{
		id:      cfg.ID,
		name:    cfg.Name,
		state:   make(map[string]interface{}),
		plugins: cfg.Plugins,
		llm:     cfg.LLM,
		log:     log,
	}, nil
}

// Core identity and state
func (a *DefaultAgent) GetID() string {
	return a.id
}

func (a *DefaultAgent) GetName() string {
	return a.name
}

func (a *DefaultAgent) GetState() map[string]interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.state
}

// Thought processes
func (a *DefaultAgent) Think(ctx context.Context) error {

	a.log.Info("Agent thinking", "ID", a.id, "Name", a.name)

	thought := &Thought{
		Content:   "Thinking...", // Replace with actual LLM response
		Type:      "fast",        // Default to fast thought
		Timestamp: time.Now(),
	}

	// Run the pre-thought plugins
	for _, p := range a.plugins {
		err := p.PreThink(ctx, thought)
		if err != nil {
			return fmt.Errorf("plugin pre-thought error: %w", err)
		}
	}

	// TODO: Format memories and input for LLM prompt
	// TODO: Call LLM for thought generation

	return nil
}

func (a *DefaultAgent) DecideAction(ctx context.Context) (action.Action, error) {
	// TODO: Implement action decision logic with LLM
	a.log.Info("Agent deciding action", "ID", a.id, "Name", a.name)
	return &action.SimpleAction{
		Type: action.ActionTypeNoop,
	}, nil
}

func (a *DefaultAgent) ReceiveOutcome(ctx context.Context, action action.Action, outcome string) error {
	// Run postaction hooks
	a.log.Info("Agent receiving outcome", "ID", a.id, "Name", a.name, "Action", action.GetType(), "Outcome", outcome)
	for _, p := range a.plugins {
		err := p.PostAction(ctx, action)
		if err != nil {
			return fmt.Errorf("plugin post-action error: %w", err)
		}
	}
	return nil
}

// Plugin system
func (a *DefaultAgent) RegisterPlugin(plugin AgentPlugin) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.log.Info("Registering plugin", "agent", a.name, "name", plugin.GetName())
	// Check for duplicate plugins
	for _, p := range a.plugins {
		if fmt.Sprintf("%T", p) == fmt.Sprintf("%T", plugin) {
			return fmt.Errorf("plugin type %T already registered", plugin)
		}
	}

	a.plugins = append(a.plugins, plugin)
	a.log.Info("Registered plugin", "type", fmt.Sprintf("%T", plugin))
	return nil
}

func (a *DefaultAgent) GetPlugins() []AgentPlugin {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.plugins
}

// Interaction capabilities
func (a *DefaultAgent) Interact(ctx context.Context, target Agent, action action.Action) error {
	a.log.Info("Interacting with agent",
		"target_id", target.GetID(),
		"action_type", action.GetType(),
	)
	return target.ReceiveInteraction(ctx, a, action)
}

func (a *DefaultAgent) ReceiveInteraction(ctx context.Context, source Agent, action action.Action) error {
	a.log.Info("Received interaction",
		"source_id", source.GetID(),
		"action_type", action.GetType(),
	)
	// TODO: Process received interaction
	return nil
}
