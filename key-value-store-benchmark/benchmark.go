package main

import (
	"github.com/matthimatiker/learning-golang/key-value-store"
	"fmt"
)

func main() {
	store := key_value_store.NewInMemoryStore()
	benchmark := key_value_store.NewBenchmark(store)
	sequential := key_value_store.NewBenchmarkRunConfiguration().NumberOfOperations(10000).ParallelOperations(1).WriteOperationRatio(0.2)
	parallel := key_value_store.NewBenchmarkRunConfiguration().NumberOfOperations(10000).ParallelOperations(10).WriteOperationRatio(0.2)

	configurations := [2]struct{
		headline string
		config key_value_store.BenchmarkRunConfiguration
	}{
		{
			headline: "## Sequential",
			config: sequential,
		},
		{
			headline: "## Parallel",
			config: parallel,
		},
	}
	for _, configuration := range configurations {
		fmt.Print(configuration.headline + "\n")
		result := benchmark.Run(configuration.config)
		fmt.Print(result.String() + "\n")
	}
}
