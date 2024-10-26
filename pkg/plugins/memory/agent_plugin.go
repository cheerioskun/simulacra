package memory

import (
	"context"
	"llm-simulation/pkg/core/action"
	"llm-simulation/pkg/core/agent"
	"llm-simulation/pkg/core/logger"
	"log/slog"
)

type AgentMemoryPlugin struct {
	memory Memory
	log    *slog.Logger
}

var _ agent.AgentPlugin = &AgentMemoryPlugin{}

func NewAgentMemoryPlugin(ctx context.Context) *AgentMemoryPlugin {
	return &AgentMemoryPlugin{
		memory: NewMemoryStore(100, nil),
		log:    ctx.Value(logger.Key).(*slog.Logger).With(logger.CategoryKey, logger.CategoryPlugin),
	}
}

func (p *AgentMemoryPlugin) GetDescription() string {
	return "AgentMemoryPlugin is a plugin that allows an agent to store and retrieve memories."
}

func (p *AgentMemoryPlugin) GetID() string {
	return "AgentMemoryPlugin"
}

func (p *AgentMemoryPlugin) GetName() string {
	return "Agent Memory Plugin"
}

func (p *AgentMemoryPlugin) OnLoad(agent agent.Agent) error {
	p.log.Info("Loading AgentMemoryPlugin")
	return nil
}

func (p *AgentMemoryPlugin) OnUnload() error {
	p.log.Info("Unloading AgentMemoryPlugin")
	return nil
}

func (p *AgentMemoryPlugin) OnAction(ctx context.Context, action action.Action) error {
	return nil
}

func (p *AgentMemoryPlugin) PreThink(ctx context.Context, input string) (string, error) {
	return input, nil
}

func (p *AgentMemoryPlugin) PostThink(ctx context.Context, thought *agent.Thought) error {
	return nil
}

func (p *AgentMemoryPlugin) PreAction(ctx context.Context, action action.Action) error {
	return nil
}

func (p *AgentMemoryPlugin) PostAction(ctx context.Context, action action.Action) error {
	return nil
}
