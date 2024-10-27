package world

import (
	"context"
	"log/slog"
	"simulacra/pkg/core/action"
	"simulacra/pkg/core/agent"
	"simulacra/pkg/core/logger"
	"simulacra/pkg/core/timemanager"
	"sync"
)

type SimpleWorld struct {
	World
	log        *slog.Logger
	config     map[string]interface{}
	state      map[string]interface{}
	agents     []agent.Agent
	plugins    []WorldPlugin
	tm         timemanager.TimeManager
	mutex      sync.Mutex
	actionChan <-chan action.Action
}

func DefaultWorld() *SimpleWorld {
	return &SimpleWorld{}
}

func (w *SimpleWorld) Initialize(ctx context.Context, config map[string]interface{}) error {
	w.log = ctx.Value(logger.Key).(*slog.Logger).With(logger.CategoryKey, logger.CategoryWorld)
	w.config = config
	return nil
}

func (w *SimpleWorld) Step(ctx context.Context) error {
	return nil
}

func (w *SimpleWorld) GetState() map[string]interface{} {
	return w.state
}

func (w *SimpleWorld) SetState(key string, value interface{}) error {
	w.state[key] = value
	return nil
}

func (w *SimpleWorld) AddAgent(ctx context.Context, agent agent.Agent) error {
	w.agents = append(w.agents, agent)
	return nil
}

func (w *SimpleWorld) RemoveAgent(ctx context.Context, agentID string) error {
	return nil
}

func (w *SimpleWorld) GetAgent(agentID string) (agent.Agent, error) {
	return nil, nil
}

func (w *SimpleWorld) GetPlugins() []WorldPlugin {
	return w.plugins
}

func (w *SimpleWorld) RegisterPlugin(plugin WorldPlugin) error {
	w.plugins = append(w.plugins, plugin)
	return nil
}
