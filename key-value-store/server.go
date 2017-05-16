package key_value_store

import "net/http"

type storeHandler struct {
	store KeyValueStore
}

// Returns a HTTP handler that operates on the given key-value store.
func NewStoreHandler(store KeyValueStore) http.Handler {
	return &storeHandler{
		store: store,
	}
}

func (handler *storeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
