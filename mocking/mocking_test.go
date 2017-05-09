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
		// toMockStore() is a type cast.
		store: toMockStore(mockedStore),
	}

	contacts.SaveName("Matthias")

	assert.True(t, called, "Set() not called")
}

// Define a type alias for a function that is used to mock DataStore.Set().
type toMockStore func (key string, value string)
// Attach a Set() function to the function to ensure that it is compatible to DataStore...
func (mock toMockStore) Set (key string, value string) {
	// ... and delegate all calls to the function itself.
	mock(key, value)
}
