package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"simulacra/pkg/core/event"
	"simulacra/pkg/core/logger"
	"simulacra/pkg/core/simulation"
	"simulacra/pkg/core/world"
	"simulacra/pkg/llm"
	"simulacra/pkg/llm/factory"
	"syscall"
	"time"
)

func main() {
	// Setup logger
	log := logger.SetupLogger(true)

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Info("Shutting down gracefully...")
		cancel()
	}()

	// Initialize components
	if err := run(ctx, log); err != nil {
		log.Error("Error running simulation", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *slog.Logger) error {
	// 1. Initialize LLM client
	llm, err := factory.New(factory.Config{
		Provider: "openrouter",
	})
	if err != nil {
		return err
	}

	// 2. Create world
	w := world.NewDefaultWorld(log)

	// 3. Initialize simulation
	sim := simulation.New(w, simulation.Config{
		StepInterval: 5 * time.Second,
	})

	// 4. Create and add initial agents
	if err := setupInitialAgents(ctx, sim, llm, log); err != nil {
		return err
	}

	// 5. Subscribe to simulation events
	if err := setupEventHandlers(sim, log); err != nil {
		return err
	}

	// 6. Start simulation
	log.Info("Starting simulation...")
	return sim.Start(ctx)
}

func setupInitialAgents(ctx context.Context, sim *simulation.Simulation, llm llm.Provider, log *slog.Logger) error {
	// TODO: Create and add initial agents
	return nil
}

func setupEventHandlers(sim *simulation.Simulation, log *slog.Logger) error {
	eventBus := sim.GetEventBus()

	// // Log all agent actions
	// err := eventBus.Subscribe(event.TypeAgentAction, func(e event.Event) error {
	// 	log.Info("Agent action",
	// 		"agent", e.Source,
	// 		"thought", e.Data["thought"],
	// 		"action", e.Data["action"],
	// 	)
	// 	return nil
	// })
	// if err != nil {
	// 	return err
	// }

	// Log agent joins/leaves
	err := eventBus.Subscribe(event.TypeAgentJoined, func(e event.Event) error {
		log.Info("Agent joined simulation", "agent", e.Target)
		return nil
	})
	if err != nil {
		return err
	}

	err = eventBus.Subscribe(event.TypeAgentLeft, func(e event.Event) error {
		log.Info("Agent left simulation", "agent", e.Target)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
