package key_value_store

import (
	"testing"
	"os"
	"io/ioutil"
	"log"
	"github.com/stretchr/testify/assert"
)

var store *Store
func TestMain(m *testing.M) {
	temporaryFile, err := ioutil.TempFile(os.TempDir(), "key_value_store_")
	if err != nil {
		log.Fatal(err)
	}
	temporaryFile.Close()
	defer func () {
		os.Remove(temporaryFile.Name())
	}()
	store, err = NewStore(temporaryFile.Name())
	if (err != nil) {
		log.Fatal("Cannot create temporary key value store.")
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

}

func Test_GetReturnsEmptyStringIfValueIsNotInStore(t *testing.T) {

}

func Test_GetReturnsValueFromStore(t *testing.T) {

}

func Test_SetOverwritesPreviousValue(t *testing.T) {

}

func Test_ReadsFromExistingFile(t *testing.T) {

}
