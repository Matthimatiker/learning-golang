package key_value_store

import (
	"os"
	"bufio"
	"strings"
)

type Store struct {
	filePath string
}

func NewStore(filePath string) (*Store, error) {
	_, err := os.Stat(filePath);
	if (os.IsNotExist(err)) {
		// File does not exist, try to create it.
		file, err := os.Create(filePath)
		if (err != nil) {
			// Could not create file.
			return nil, err
		}
		file.Close()
		err = nil
	}
	if (err != nil) {
		// Unexpected file error.
		return nil, err
	}
	return &Store{
		filePath: filePath,
	}, nil
}

func (store *Store) Get(key string) string {
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
			value = lineValue
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return value
}

func (store *Store) Set(key string, value string) {
	file := openFile(store.filePath, os.O_APPEND|os.O_WRONLY)
	defer file.Close()

	// We are adding values to the store in O(1), yay!
	_, err := file.WriteString(key + "=" + value + "\n")
	if (err != nil) {
		panic(err)
	}
	err = file.Sync()
	if (err != nil) {
		panic(err)
	}
}

// Opens the file with the given mode. Panics in case of error
// as that should not happen: File is checked during creation of
// store.
func openFile(path string, mode int) *os.File {
	file, err := os.OpenFile(path, mode, os.ModeExclusive)
	if (err != nil) {
		panic(err)
	}
	return file
}
