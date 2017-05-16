package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func Test_SendsPost(t *testing.T) {
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		assert.Equal(t, http.MethodPost, r.Method)
	}))
	defer ts.Close()

	post(ts.URL)

	assert.Equal(t, true, called, "Server not called.")
}
