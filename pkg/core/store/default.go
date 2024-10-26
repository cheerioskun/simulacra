package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

var DefaultStore *leveldb.DB

type DefaultStoreType = *leveldb.DB

const InMemory = false

func InitDefaultStore() {
	if InMemory {
		DefaultStore, _ = leveldb.Open(storage.NewMemStorage(), nil)
	} else {
		DefaultStore, _ = leveldb.OpenFile("store.db", nil)
	}
}
