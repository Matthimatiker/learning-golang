package key_value_store

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

type storeHandler struct {
	store KeyValueStore
}

// Returns a HTTP handler that operates on the given key-value store.
func NewStoreHandler(store KeyValueStore) http.Handler {
	return &storeHandler{
		store: store,
	}
}

// Handles a HTTP request and maps it to the store.
func (handler *storeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	switch method := r.Method; method {
	case http.MethodGet:
		value := handler.store.Get(key);
		if (value == "") {
			w.WriteHeader(404);
		} else {
			w.WriteHeader(200)
			fmt.Fprint(w, value)
		}
	case http.MethodPost:
		value, err := ioutil.ReadAll(r.Body)
		if (err != nil) {
			http.Error(w, err.Error(), 500)
			return;
		}
		handler.store.Set(key, string(value))
		w.WriteHeader(201)
	default:
		http.Error(w, "Unsupported method '" + method + "'", 400)
	}

}
