package key_value_store

import (
	"testing"
	"os"
	"io/ioutil"
	"log"
	"github.com/stretchr/testify/assert"
)

var store KeyValueStore
var temporaryFile string

func setUpStore() {
	temporaryFile = createTemporaryFile()
	var err error
	store, err = NewStore(temporaryFile)
	if (err != nil) {
		log.Fatalf("Cannot create temporary key value store: %v", err)
	}
}

func tearDownStore() {
	os.Remove(temporaryFile)
	store = nil
	temporaryFile = ""
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

func Test_StoreDoesNotReturnErrorIfItCreatesFile(t *testing.T) {
	temporaryFile := createTemporaryFile()
	os.Remove(temporaryFile)
	defer os.Remove(temporaryFile)

	_, err := NewStore(temporaryFile)

	assert.New(t).Nil(err, "Store returned error.")
}

func Test_GetReturnsEmptyStringIfValueIsNotInStore(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	assert.New(t).Equal("", store.Get("missing"))
}

func Test_GetReturnsValueFromStore(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	store.Set("my-key", "my-value")

	assert.New(t).Equal("my-value", store.Get("my-key"))
}

func Test_SetOverwritesPreviousValue(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	store.Set("my-key", "my-value")
	store.Set("my-key", "my-new-value")

	assert.New(t).Equal("my-new-value", store.Get("my-key"))
}

func Test_ReadsFromExistingFile(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	store.Set("my-key", "my-value")

	// Read with another store instance from the same file.
	newStore, err := NewStore(temporaryFile)
	assert.New(t).Nil(err);

	assert.New(t).Equal("my-value", newStore.Get("my-key"))
}

func Test_AllReturnsEmptyMapIfStoreIsEmpty(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	all := store.All()

	assert.New(t).Len(all, 0)
}

func Test_AllReturnsCorrectValues(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	store.Set("a", "b")
	store.Set("c", "d")

	all := store.All()

	assert.New(t).Equal("b", all["a"])
	assert.New(t).Equal("d", all["c"])
	assert.New(t).Len(all, 2)
}

func Test_AllContainsCorrectValuesIfKeysWereOverwritten(t *testing.T) {
	setUpStore()
	defer tearDownStore()

	store.Set("a", "b")
	store.Set("c", "d")
	store.Set("a", "oha")

	all := store.All()

	assert.New(t).Equal("oha", all["a"])
	assert.New(t).Len(all, 2)
}

func Test_ToKeyValueReturnsCorrectKey(t *testing.T) {
	key, _ := ToKeyValue("abc=def")

	assert.New(t).Equal("abc", key)
}

func Test_ToKeyValueReturnsCorrectValue(t *testing.T) {
	_, value := ToKeyValue("abc=def")

	assert.New(t).Equal("def", value)
}

func Test_ToKeyValueReturnsWholeTextAsKeyIfThereIsNoSeparator(t *testing.T) {
	key, _ := ToKeyValue("abc")

	assert.New(t).Equal("abc", key)
}

func Test_ToKeyValueReturnsEmptyStringAsValueIfThereIsNoSeparator(t *testing.T) {
	_, value := ToKeyValue("abc")

	assert.New(t).Equal("", value)
}

func Test_IsKeyValuePairReturnsFalseIfNoValueIsPresent(t *testing.T) {
	assert.New(t).False(IsKeyValuePair("abc"))
}

func Test_IsKeyValuePairReturnsTrueIfKeyAndValueArePresent(t *testing.T) {
	assert.New(t).True(IsKeyValuePair("abc=def"))
}

func Test_IsKeyValuePairReturnsTrueIfValueIsEmpty(t *testing.T) {
	assert.New(t).True(IsKeyValuePair("abc="))
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
