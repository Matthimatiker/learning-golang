package mocking

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_SavesName(t *testing.T) {
	called := false
	mockedStore := func (key string, value string) {
		called = true
		assert.Equal(t, "name", key)
		assert.Equal(t, "Matthias", value)
	}
	contacts := Contacts{
		store: toMockStore(mockedStore),
	}

	contacts.SaveName("Matthias")

	assert.True(t, called, "Set() not called")
}

type toMockStore func (key string, value string)
func (mock toMockStore) Set (key string, value string) {
	mock(key, value)
}
