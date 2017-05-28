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

	configs := [2]key_value_store.BenchmarkRunConfiguration{
		sequential,
		parallel,
	}
	for _, config := range configs {
		result := benchmark.Run(config)
		fmt.Print(result.String() + "\n")
	}
}
