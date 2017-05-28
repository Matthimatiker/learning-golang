package main

import (
	"github.com/matthimatiker/learning-golang/key-value-store"
	"fmt"
)

func main() {
	store := key_value_store.NewInMemoryStore()
	benchmark := key_value_store.NewBenchmark(store)
	sequential := key_value_store.NewBenchmarkRunConfiguration().NumberOfOperations(10000).ParallelOperations(1).WriteOperationRatio(0.2)
	result := benchmark.Run(sequential)
	fmt.Print(result.String())
}
