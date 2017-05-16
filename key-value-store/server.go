package key_value_store

import "net/http"

type StoreHandler http.Handler

// Returns a HTTP handler that operates on the given key-value store.
func NewStoreHandler(store KeyValueStore) StoreHandler {
	return func(w http.ResponseWriter, r *http.Request) {

	};
}
