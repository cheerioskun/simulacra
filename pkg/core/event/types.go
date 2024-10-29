package event

import (
	"time"
)

// Type represents the type of event
type Type string

const (
	TypeAgentJoined      Type = "agent_joined"
	TypeAgentLeft        Type = "agent_left"
	TypeAgentAction      Type = "agent_action"
	TypeWorldStateChange Type = "world_state_change"
	TypeAgentInteraction Type = "agent_interaction"
)

// Event represents a basic event in the system
type Event struct {
	ID        string                 // Unique identifier for the event
	Type      Type                   // Type of the event
	Timestamp time.Time              // When the event occurred
	Source    string                 // ID of the source (agent/world/system)
	Target    string                 // ID of the target (if applicable)
	Data      map[string]interface{} // Event payload
}

// Handler represents an event handler function
type Handler func(Event) error

// Bus defines the interface for the event system
type Bus interface {
	// Publish sends an event to all subscribers
	Publish(Event) error

	// Subscribe registers a handler for specific event types
	Subscribe(Type, Handler) error

	// Unsubscribe removes a handler for an event type
	Unsubscribe(Type, Handler) error

	// Close shuts down the event bus
	Close() error
}
