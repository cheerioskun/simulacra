package event

import (
	"fmt"
	"sync"
)

type defaultBus struct {
	handlers map[Type][]Handler
	mu       sync.RWMutex
}

// NewEventBus creates a new event bus instance
func NewEventBus() Bus {
	return &defaultBus{
		handlers: make(map[Type][]Handler),
	}
}

func (b *defaultBus) Publish(event Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	handlers, exists := b.handlers[event.Type]
	if !exists {
		return nil // No handlers for this event type
	}

	var errs []error
	for _, handler := range handlers {
		if err := handler(event); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors publishing event: %v", errs)
	}
	return nil
}

func (b *defaultBus) Subscribe(eventType Type, handler Handler) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.handlers[eventType]; !exists {
		b.handlers[eventType] = []Handler{}
	}
	b.handlers[eventType] = append(b.handlers[eventType], handler)
	return nil
}

func (b *defaultBus) Unsubscribe(eventType Type, handler Handler) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if handlers, exists := b.handlers[eventType]; exists {
		for i, h := range handlers {
			if fmt.Sprintf("%p", h) == fmt.Sprintf("%p", handler) {
				b.handlers[eventType] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (b *defaultBus) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Clear all handlers
	b.handlers = make(map[Type][]Handler)
	return nil
}
