package memory

import (
	"context"
	"simulacra/pkg/core/store"
)

type MemoryScore int

var (
	MemoryScoreLow    MemoryScore = 1
	MemoryScoreMedium MemoryScore = 5
	MemoryScoreHigh   MemoryScore = 9
)

type Memory interface {
	Retrieve(ctx context.Context, query string, threshold MemoryScore) (string, error)
	Store(ctx context.Context, memory string, score MemoryScore) error
}

type MemoryStore struct {
	memories []TimestampedMemory
	capacity int
	store    store.DefaultStoreType
}

var _ Memory = &MemoryStore{}

type TimestampedMemory struct {
	Timestamp int64
	Content   interface{}
	Type      string
	Metadata  map[string]interface{}
}

func NewMemoryStore(capacity int, store store.DefaultStoreType) *MemoryStore {
	return &MemoryStore{
		memories: make([]TimestampedMemory, 0, capacity),
		capacity: capacity,
		store:    store,
	}
}

func (m *MemoryStore) Retrieve(ctx context.Context, query string, threshold MemoryScore) (string, error) {
	return "", nil
}

func (m *MemoryStore) Store(ctx context.Context, memory string, score MemoryScore) error {
	return nil
}
