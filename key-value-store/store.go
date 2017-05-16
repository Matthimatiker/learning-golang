package key_value_store

import (
	"os"
	"bufio"
	"strings"
	"encoding/base64"
)

type KeyValueStore interface {
	Get(key string) string
	Set(key string, value string)
	All() (map[string]string)
}

type fileKeyValueStore struct {
	filePath string
}

func NewStore(filePath string) (KeyValueStore, error) {
	_, err := os.Stat(filePath);
	if (os.IsNotExist(err)) {
		// File does not exist, try to create it.
		file, createErr := os.Create(filePath)
		if (createErr != nil) {
			// Could not create file.
			return nil, createErr
		}
		file.Close()
		err = nil
	}
	if (err != nil) {
		// Unexpected file error.
		return nil, err
	}
	return &fileKeyValueStore{
		filePath: filePath,
	}, err
}

// Searches for the value that belongs to the given key.
// Returns an empty string if the value is not in the store.
func (store *fileKeyValueStore) Get(key string) string {
	value := ""
	file := openFile(store.filePath, os.O_RDONLY)
	defer file.Close()

	// We are reading in  O(n)... Not that good...
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		lineKey, lineValue := parts[0], parts[1]
		if (lineKey == key) {
			value = decode(lineValue)
		}
	}
	assertNoError(scanner.Err())

	return value
}

// Stores the given value in the store.
func (store *fileKeyValueStore) Set(key string, value string) {
	file := openFile(store.filePath, os.O_APPEND|os.O_WRONLY)
	defer file.Close()

	// We are adding values to the store in O(1), yay!
	_, err := file.WriteString(key + "=" + encode(value) + "\n")
	assertNoError(err)
	err = file.Sync()
	assertNoError(err)
}

// Returns all values in the store.
func (store *fileKeyValueStore) All() (map[string]string) {
	file := openFile(store.filePath, os.O_RDONLY)
	defer file.Close()

	all := map[string]string{}
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		key, value := ToKeyValue(scanner.Text())
		all[key] = decode(value)
	}
	assertNoError(scanner.Err())

	return all
}

// Checks if the given string contains a key/value pair
func IsKeyValuePair(text string) bool {
	return strings.Contains(text, "=")
}

// Splits the given string into key and value (separated by "=").
func ToKeyValue(text string) (key string, value string) {
	parts := strings.SplitN(text, "=", 2)
	if (len(parts) == 1) {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// Encodes a value before it is written to the store.
//
// Encoding ensures that we get rid of newlines, which are problematic when
// reading the store file line by line
func encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// Decodes a value that is read from store.
func decode(encodedValue string) string {
	decoded, err := base64.StdEncoding.DecodeString(encodedValue)
	assertNoError(err)
	return string(decoded)
}

// Opens the file with the given mode. Panics in case of error
// as that should not happen: File is checked during creation of
// store.
func openFile(path string, mode int) *os.File {
	file, err := os.OpenFile(path, mode, os.ModeExclusive)
	assertNoError(err)
	return file
}

// Asserts that there is no error, panics otherwise.
func assertNoError(err error) {
	if (err != nil) {
		panic(err)
	}
}
