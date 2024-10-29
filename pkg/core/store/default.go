package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

var defaultStore *leveldb.DB

type DefaultStoreType = *leveldb.DB

const InMemory = false

func InitDefaultStore() {
	if InMemory {
		defaultStore, _ = leveldb.Open(storage.NewMemStorage(), nil)
	} else {
		defaultStore, _ = leveldb.OpenFile("store.db", nil)
	}
}

func DefaultStore() DefaultStoreType {
	return defaultStore
}
