package main

import (
	"github.com/matthimatiker/learning-golang/key-value-store"
	"os"
	"net/http"
	"fmt"
)

// Starts a server on port 8080 that interacts with the default
// key value store (same that is used by CLI tool).
//
//     go run key-value-store-server/server.go
func main() {
	store, err := key_value_store.NewStore(getStoreFile())
	if (err != nil) {
		panic(err)
	}
	handler := key_value_store.NewStoreHandler(store)
	fmt.Println("Listening on:")
	fmt.Println("http://localhost:8080");
	panic(http.ListenAndServe(":8080", handler))
}

func getStoreFile() string {
	return os.TempDir() + "/learning-golang-default.store"
}

