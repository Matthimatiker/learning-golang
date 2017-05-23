package key_value_store

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func setUpWebClient() (func (), *webClient, SimpleKeyValueStore) {
	store := NewInMemoryStore()
	ts := httptest.NewServer(NewStoreHandler(store))
	client := NewWebClient(ts.URL)
	return func () {
		ts.Close()
	}, client, store
}

func Test_CanReadValueFromWebStore(t *testing.T) {
	tearDown, client, store := setUpWebClient()
	defer tearDown()
	store.Set("hello", "world")

	value := client.Get("hello")

	assert.Equal(t, "world", value)
}

func Test_CanWriteValueToWebStore(t *testing.T) {
	tearDown, client, store := setUpWebClient()
	defer tearDown()

	client.Set("hello", "world")

	assert.Equal(t, "world", store.Get("hello"))
}

func Test_CanReadValueThatWasWrittenViaWebClient(t *testing.T) {
	tearDown, client, _ := setUpWebClient()
	defer tearDown()

	client.Set("hello", "world")
	value := client.Get("hello")

	assert.Equal(t, "world", value)
}
