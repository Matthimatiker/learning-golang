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

func NewWebClient(url string) (*webClient) {
	return &webClient{
		url: strings.TrimSuffix(url, "/"),
	}
}

func (client *webClient) Get(key string) string {
	response, err := http.Get(client.url + "/" + key)
	assertNoError(err)
	value, err := ioutil.ReadAll(response.Body)
	assertNoError(err)
	return string(value)
}

func (client *webClient) Set(key string, value string) {
	_, err := http.Post(client.url + "/" + key, "text/plain", bytes.NewBufferString(value))
	assertNoError(err)
}
