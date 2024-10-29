package simulation

import (
	"context"
	"fmt"
	"simulacra/pkg/core/agent"
	"simulacra/pkg/core/event"
	"simulacra/pkg/core/world"
	"sync"
	"time"
)

// Simulation orchestrates the interaction between world and agents
type Simulation struct {
	world    world.World
	eventBus event.Bus
	agents   map[string]agent.Agent

	// Control channels
	stopCh   chan struct{}
	pauseCh  chan struct{}
	resumeCh chan struct{}

	// Configuration
	stepInterval time.Duration

	mu sync.RWMutex
}

func New(w world.World, config Config) *Simulation {
	return &Simulation{
		world:        w,
		eventBus:     event.NewEventBus(),
		agents:       make(map[string]agent.Agent),
		stopCh:       make(chan struct{}),
		pauseCh:      make(chan struct{}),
		resumeCh:     make(chan struct{}),
		stepInterval: config.StepInterval,
	}
}

type Config struct {
	StepInterval time.Duration
}

// AddAgent adds an agent to the simulation
func (s *Simulation) AddAgent(ctx context.Context, a agent.Agent) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.agents[a.GetID()]; exists {
		return fmt.Errorf("agent %s already exists", a.GetID())
	}

	s.agents[a.GetID()] = a

	// Notify about new agent
	return s.eventBus.Publish(event.Event{
		Type:      event.TypeAgentJoined,
		Source:    "simulation",
		Target:    a.GetID(),
		Timestamp: time.Now(),
	})
}

// Start begins the simulation loop
func (s *Simulation) Start(ctx context.Context) error {
	ticker := time.NewTicker(s.stepInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-s.stopCh:
			return nil
		case <-s.pauseCh:
			<-s.resumeCh // Wait for resume signal
		case <-ticker.C:
			if err := s.step(ctx); err != nil {
				return fmt.Errorf("simulation step error: %w", err)
			}
		}
	}
}

// step advances the simulation by one tick
func (s *Simulation) step(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. Process world state
	_ = s.world.GetState()

	// 2. Process each agent in parallel
	var wg sync.WaitGroup
	errs := make(chan error, len(s.agents))

	for _, a := range s.agents {
		wg.Add(1)
		go func(agent agent.Agent) {
			defer wg.Done()

			err := agent.Think(ctx)
			if err != nil {
				errs <- fmt.Errorf("agent %s think error: %w", agent.GetID(), err)
				return
			}

			action, err := agent.DecideAction(ctx)
			if err != nil {
				errs <- fmt.Errorf("agent %s action error: %w", agent.GetID(), err)
				return
			}
			// Based on world state and action, decide outcome
			outcome, err := s.world.ApplyAction(action)
			if err != nil {
				errs <- fmt.Errorf("world apply action error: %w", err)
				return
			}
			// Send back outcome to agent
			agent.ReceiveOutcome(ctx, action, outcome)

			// Publish agent events
			s.eventBus.Publish(event.Event{
				Type:      event.TypeAgentAction,
				Source:    agent.GetID(),
				Timestamp: time.Now(),
				Data: map[string]interface{}{
					"action": action,
				},
			})
		}(a)
	}

	wg.Wait()
	close(errs)

	// Collect any errors
	var errors []error
	for err := range errs {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("multiple errors during step: %v", errors)
	}

	return nil
}

// Stop halts the simulation
func (s *Simulation) Stop() {
	close(s.stopCh)
}

// Pause temporarily halts the simulation
func (s *Simulation) Pause() {
	s.pauseCh <- struct{}{}
}

// Resume continues a paused simulation
func (s *Simulation) Resume() {
	s.resumeCh <- struct{}{}
}

// GetEventBus returns the simulation's event bus
func (s *Simulation) GetEventBus() event.Bus {
	return s.eventBus
}
