package world

import (
	"encoding/json"
	"log/slog"
	"simulacra/pkg/core/logger"
	"simulacra/pkg/core/store"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

type defaultWorld struct {
	state store.DefaultStoreType
	mu    sync.RWMutex
	log   *slog.Logger
}

func NewDefaultWorld(log *slog.Logger) *defaultWorld {
	return &defaultWorld{
		state: store.DefaultStore(),
		log:   log.With(logger.CategoryKey, logger.CategoryWorld),
	}
}

func (w *defaultWorld) GetState() map[string]interface{} {
	w.mu.RLock()
	defer w.mu.RUnlock()

	w.log.Debug("Getting world state")
	s, err := w.state.Get([]byte(StatePrefix), nil)
	if err != nil {
		w.log.Error("Failed to get state", "error", err)
		return nil
	}

	var ret map[string]interface{}
	err = json.Unmarshal(s, &ret)
	if err != nil {
		w.log.Error("Failed to unmarshal state", "error", err)
		return nil
	}

	w.log.Debug("Successfully retrieved world state")
	return ret
}

func (w *defaultWorld) SetState(state map[string]interface{}) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	b, err := json.Marshal(state)
	if err != nil {
		w.log.Error("Failed to marshal state", "error", err)
		return err
	}

	if err := w.state.Put([]byte(StatePrefix), b, nil); err != nil {
		w.log.Error("Failed to set state", "error", err)
		return err
	}

	w.log.Debug("Successfully set world state", "state", spew.Sdump(state))
	return nil
}

func (w *defaultWorld) IsValidAction(action interface{}) bool {
	// TODO: Implement action validation logic
	return true
}

func (w *defaultWorld) ApplyAction(action interface{}) (string, error) {
	// TODO: Implement action application logic
	return "", nil
}
