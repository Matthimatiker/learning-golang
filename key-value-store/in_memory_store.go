package key_value_store

type inMemoryStore struct {
	data map[string]string
}

// Creates an empty in-memory key-value store.
func NewInMemoryStore() (KeyValueStore) {
	return &inMemoryStore{
		data: map[string]string{},
	}
}

func (store *inMemoryStore) Get(key string) string {
	return store.data[key];
}

func (store *inMemoryStore) Set(key string, value string) {
	store.data[key] = value;
}

func (store *inMemoryStore) All() (map[string]string) {
	return store.data;
}
