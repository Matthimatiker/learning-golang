package main

import (
	"os"
	"github.com/matthimatiker/learning-golang/key-value-store"
	"fmt"
)

func main() {
	store, err := key_value_store.NewStore(getStoreFile())
	if (err != nil) {
		panic(err)
	}
	for _, arg :=range os.Args[1:] {
		key, value := key_value_store.ToKeyValue(arg)
		if (key_value_store.IsKeyValuePair(arg)) {
			store.Set(key, value)

		} else {
			value = store.Get(key)
		}
		fmt.Printf("> %s = %s", key, value)
		fmt.Println()
	}
}

func getStoreFile() string {
	return os.TempDir() + "/learning-golang-default.store"
}
