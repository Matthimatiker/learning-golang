package main

import (
	"fmt"
	"github.com/matthimatiker/learning-golang/key-value-store"
)

func main() {
	store := key_value_store.NewInMemoryStore()
	benchmark := key_value_store.NewBenchmark(store)

	configurations := [2]struct {
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
	for _, configuration := range configurations {
		fmt.Print(configuration.headline + "\n")
		result := benchmark.Run(configuration.config)
		fmt.Print(result.String() + "\n")
	}
}
