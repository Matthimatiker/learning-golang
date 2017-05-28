package main

import (
	"fmt"
	"github.com/matthimatiker/learning-golang/key-value-store"
	"net/http/httptest"
	"os"
)

func main() {
	server := httptest.NewServer(key_value_store.NewStoreHandler(key_value_store.NewInMemoryStore()))
	defer server.Close()
	fileStore, _ := key_value_store.NewStore(getStoreFile())
	defer os.Remove(getStoreFile())

	scenarios := [...]struct{
		headline string
		store key_value_store.SimpleKeyValueStore
	}{
		{
			headline: "# InMemoryStore",
			store: key_value_store.NewInMemoryStore(),
		},
		{
			headline: "# WebClient against InMemoryStore",
			store: key_value_store.NewWebClient(server.URL),
		},
		{
			headline: "# File based store",
			store: fileStore,
		},
	}
	configurations := [...]struct {
		headline string
		config   key_value_store.BenchmarkRunConfiguration
	}{
		{
			headline: "## Sequential",
			config:   key_value_store.NewBenchmarkRunConfiguration().NumberOfOperations(10000).ParallelOperations(1).WriteOperationRatio(0.2),
		},
		{
			headline: "## Parallel",
			config:   key_value_store.NewBenchmarkRunConfiguration().NumberOfOperations(10000).ParallelOperations(10).WriteOperationRatio(0.2),
		},
	}

	for _, scenario := range scenarios {
		fmt.Print(scenario.headline + "\n\n")
		benchmark := key_value_store.NewBenchmark(scenario.store)
		for _, configuration := range configurations {
			fmt.Print(configuration.headline + "\n")
			result := benchmark.Run(configuration.config)
			fmt.Print(result.String() + "\n")
		}
	}
}

func getStoreFile() string {
	return os.TempDir() + "/key-value-store-benchmark.store"
}
