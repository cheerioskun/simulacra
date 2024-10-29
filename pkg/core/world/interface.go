package world

// World defines the interface for the simulation world
type World interface {
	// Core world operations
	GetState() map[string]interface{}
	SetState(state map[string]interface{}) error

	// World-specific logic
	IsValidAction(action interface{}) bool
	ApplyAction(action interface{}) (outcome string, err error)
}
