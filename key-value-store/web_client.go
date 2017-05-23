package key_value_store

import (
	"strings"
	"bytes"
	"io/ioutil"
	"net/http"
)

type webClient struct {
	url string
}

// Creates a new web client that interacts with the store server at the given URL.
//
// The web client itself is a SimpleKeyValueStore.
// Multiple clients can share a key value server without influencing each other
// by using a  specific URL segment:
//
//    firstClient := NewWebClient("http://my-key-value.server/bucket-one")
//    secondClient := NewWebClient("http://my-key-value.server/bucket-two")
func NewWebClient(url string) (*webClient) {
	return &webClient{
		url: strings.TrimSuffix(url, "/"),
	}
}

func (client *webClient) Get(key string) string {
	// TODO check http status code, err handles only network errors
	response, err := http.Get(client.toUrl(key))
	assertNoError(err)
	value, err := ioutil.ReadAll(response.Body)
	assertNoError(err)
	return string(value)
}

func (client *webClient) Set(key string, value string) {
	_, err := http.Post(client.toUrl(key), "text/plain", bytes.NewBufferString(value))
	assertNoError(err)
}

// Returns the URL that is used to access (read & write) the given key.
func (client *webClient) toUrl(key string) string {
	return client.url + "/" + key
}
