package timemanager

import (
	"context"
	"log/slog"
	"simulacra/pkg/core/logger"
	"sync"
	"time"
)

type SimulationSpeed float64 // 1.0 = real time, 2.0 = 2x speed, 0.5 = half speed

type TimeManager struct {
	mu                  sync.RWMutex
	simulationSpeed     SimulationSpeed
	startTime           time.Time // Real world start time
	simStartTime        time.Time // Simulation start time
	isPaused            bool
	pausedAt            time.Time
	totalPausedDuration time.Duration
	log                 *slog.Logger
}

func NewTimeManager(ctx context.Context) *TimeManager {
	now := time.Now()
	return &TimeManager{
		simulationSpeed: 1.0,
		startTime:       now,
		simStartTime:    now,
		isPaused:        false,
		log:             ctx.Value(logger.Key).(*slog.Logger).With(logger.CategoryKey, logger.CategoryTimeManager),
	}
}

func (tm *TimeManager) GetSimulationTime() time.Time {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if tm.isPaused {
		elapsed := tm.pausedAt.Sub(tm.startTime)
		simElapsed := time.Duration(float64(elapsed) * float64(tm.simulationSpeed))
		return tm.simStartTime.Add(simElapsed)
	}

	elapsed := time.Since(tm.startTime) - tm.totalPausedDuration
	simElapsed := time.Duration(float64(elapsed) * float64(tm.simulationSpeed))
	return tm.simStartTime.Add(simElapsed)
}

func (tm *TimeManager) SetSpeed(speed SimulationSpeed) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Store current simulation time before changing speed
	currentSimTime := tm.GetSimulationTime()

	// Reset start times with new speed
	tm.startTime = time.Now()
	tm.simStartTime = currentSimTime
	tm.simulationSpeed = speed
}

func (tm *TimeManager) Pause() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if !tm.isPaused {
		tm.isPaused = true
		tm.pausedAt = time.Now()
	}
}

func (tm *TimeManager) Resume() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.isPaused {
		tm.totalPausedDuration += time.Since(tm.pausedAt)
		tm.isPaused = false
	}
}

func (tm *TimeManager) WaitSimulationTime(duration time.Duration) {
	adjustedDuration := time.Duration(float64(duration) / float64(tm.simulationSpeed))
	time.Sleep(adjustedDuration)
}

// Convert real time duration to simulation duration
func (tm *TimeManager) ToSimulationDuration(realDuration time.Duration) time.Duration {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return time.Duration(float64(realDuration) * float64(tm.simulationSpeed))
}

// Convert simulation duration to real time duration
func (tm *TimeManager) ToRealDuration(simDuration time.Duration) time.Duration {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return time.Duration(float64(simDuration) / float64(tm.simulationSpeed))
}
