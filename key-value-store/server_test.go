package key_value_store

import (
	"testing"
	"net/http/httptest"
	"strings"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func setUp() (http.Handler, KeyValueStore) {
	store := NewInMemoryStore()
	return NewStoreHandler(store), store
}

func Test_InMemoryStoreWorks(t *testing.T) {
	_, store := setUp()

	store.Set("hello", "world")

	assert.Equal(t, "world", store.Get("hello"))
}

func Test_ReturnsCode201OnCreation(t *testing.T) {
	handler, _ := setUp()

	req := httptest.NewRequest("POST", "/my-key", strings.NewReader("my-value"))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func Test_AddsKeyToStore(t *testing.T) {
	handler, store := setUp()

	req := httptest.NewRequest("POST", "/my-key", strings.NewReader("my-value"))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, "my-value", store.Get("my-key"))
}

func Test_ReturnsCode200WhenFound(t *testing.T) {
	handler, store := setUp()
	store.Set("my-key", "my-value");

	req := httptest.NewRequest("GET", "/my-key", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func Test_ReturnsValueInResponse(t *testing.T) {
	handler, store := setUp()
	store.Set("my-key", "my-value");

	req := httptest.NewRequest("GET", "/my-key", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, "my-value", w.Body.String())
}

func Test_ReturnsCode404IfNotExists(t *testing.T) {
	handler, _ := setUp()

	req := httptest.NewRequest("GET", "/missing-key", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func Test_RejectsRequestWithUnsupportedMethod(t *testing.T) {
	handler, _ := setUp()

	req := httptest.NewRequest("PATCH", "/hello", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
