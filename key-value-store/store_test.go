package key_value_store

import (
	"testing"
	"os"
	"io/ioutil"
	"log"
	"github.com/stretchr/testify/assert"
)

var store *Store
var temporaryFile string
func TestMain(m *testing.M) {
	temporaryFile = createTemporaryFile()
	defer os.Remove(temporaryFile)
	var err error
	store, err = NewStore(temporaryFile)
	if (err != nil) {
		log.Fatalf("Cannot create temporary key value store: %v", err)
	}

	returnCode := m.Run()

	store = nil

	os.Exit(returnCode)
}

func Test_ReturnsErrorIfInvalidFilePathIsProvided(t *testing.T) {
	_, err := NewStore("/this/will/not/end.well")

	assert.New(t).NotNil(err, "Expected error if file path is invalid")
}

func Test_CreatesFileIfItDoesNotExist(t *testing.T) {
	temporaryFile := createTemporaryFile()
	os.Remove(temporaryFile)
	defer os.Remove(temporaryFile)
	var err error
	_, err = os.Stat(temporaryFile);
	assert.New(t).True(os.IsNotExist(err), "File was not removed.")

	NewStore(temporaryFile)

	_, err = os.Stat(temporaryFile);
	assert.New(t).Nil(err, "Store did not create file.")
}

func Test_GetReturnsEmptyStringIfValueIsNotInStore(t *testing.T) {
	assert.New(t).Equal("", store.Get("missing"))
}

func Test_GetReturnsValueFromStore(t *testing.T) {
	store.Set("my-key", "my-value")

	assert.New(t).Equal("my-value", store.Get("my-key"))
}

func Test_SetOverwritesPreviousValue(t *testing.T) {
	store.Set("my-key", "my-value")
	store.Set("my-key", "my-new-value")

	assert.New(t).Equal("my-new-value", store.Get("my-key"))
}

func Test_ReadsFromExistingFile(t *testing.T) {
	store.Set("my-key", "my-value")

	// Read with another store instance from the same file.
	newStore, err := NewStore(temporaryFile)
	assert.New(t).Nil(err);

	assert.New(t).Equal("my-value", newStore.Get("my-key"))
}

// Creates a temporary file and returns its path.
func createTemporaryFile() string {
	temporaryFile, err := ioutil.TempFile(os.TempDir(), "key_value_store_")
	if err != nil {
		log.Fatal(err)
	}
	temporaryFile.Close()
	return temporaryFile.Name()
}