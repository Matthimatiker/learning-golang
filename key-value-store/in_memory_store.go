package key_value_store

import "sync"

type inMemoryStore struct {
	data map[string]string
	lock *sync.RWMutex
}

// Creates an empty in-memory key-value store.
func NewInMemoryStore() (KeyValueStore) {
	return &inMemoryStore{
		data: map[string]string{},
		lock: &sync.RWMutex{},
	}
}

func (store *inMemoryStore) Get(key string) string {
	store.lock.RLock()
	defer store.lock.RUnlock()
	return store.data[key];
}

func (store *inMemoryStore) Set(key string, value string) {
	store.lock.Lock()
	defer store.lock.Unlock()
	store.data[key] = value;
}

func (store *inMemoryStore) All() (map[string]string) {
	return store.data;
}
