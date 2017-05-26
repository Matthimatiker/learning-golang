package key_value_store

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

type benchmark struct {
	store SimpleKeyValueStore
}

func NewBenchmark(store SimpleKeyValueStore) *benchmark {
	return &benchmark{
		store: store,
	}
}

// Runs a benchmark against the store.
//
func (benchmark *benchmark) run(config benchmarkRunConfiguration) benchmarkRunResult {
	// Use random number generator with fixed seed to ensure that runs are deterministic.
	random := rand.New(rand.NewSource(0))
	operations := make(chan func (store SimpleKeyValueStore), config.parallelOperations)
	// Prepare several workers, depending on the number of parallel operations.
	workers := &sync.WaitGroup{}
	for i := 0; i < config.parallelOperations; i++ {
		workers.Add(1)
		go func () {
			for operation := range operations {
				operation(benchmark.store)
			}
			workers.Done()
		}()
	}
	start := time.Now()
	// Push the configured number of operations to the channel.
	// At the same time, the workers start to read and execute the operations.
	for i := 0; i < config.numberOfOperations; i++ {
		if (random.Float32() < config.writeOperationRatio) {
			operations <- func (store SimpleKeyValueStore) {
				store.Set("x", "y")
			}
		} else {
			operations <- func (store SimpleKeyValueStore) {
				store.Get("x")
			}
		}
	}
	// Close the channel to ensure that the workers will terminate when the channel is empty.
	close(operations)
	// Wait for all workers to complete.
	workers.Wait()
	return benchmarkRunResult{
		Config: config,
		Runtime: time.Since(start),
	}
}

// Encapsulates a configuration for a benchmark run.
type benchmarkRunConfiguration struct {
	numberOfOperations  int
	parallelOperations  int
	writeOperationRatio float32
}

// Creates a new benchmark configuration with default values.
func NewBenchmarkRunConfiguration() benchmarkRunConfiguration {
	return benchmarkRunConfiguration{
		numberOfOperations:  1000,
		parallelOperations:  1,
		writeOperationRatio: 0.2,
	}
}

// Sets the number of operations that will be executed and returns the new configuration.
func (config benchmarkRunConfiguration) NumberOfOperations(value int) benchmarkRunConfiguration {
	config.numberOfOperations = value
	return config
}

// Sets the number of operations that will be executed in parallel and returns the new configuration.
func (config benchmarkRunConfiguration) ParallelOperations(value int) benchmarkRunConfiguration {
	config.parallelOperations = value
	return config
}

// Defines the number of write operations among the whole number of operations.
// value must be a float between 0.0 and 1.0.
func (config benchmarkRunConfiguration) WriteOperationRatio(value float32) benchmarkRunConfiguration {
	if value < 0.0 || value > 1.0 {
		panic("Write ration value must be in range [0.0..1.0]")
	}
	config.writeOperationRatio = value
	return config
}

// Creates a string representation of the configuration.
func (config benchmarkRunConfiguration) String() string {
	return fmt.Sprintf(`
	Config:
	- Number of operations: %v
	- Parallel operations: %v
	- Write operations: approx. %v%%
	`, config.numberOfOperations, config.parallelOperations, config.writeOperationRatio*100)
}

// Contains the result of a benchmark run.
type benchmarkRunResult struct {
	// The configuration that has been applied.
	Config benchmarkRunConfiguration
	// The number of seconds the benchmark took to run.
	Runtime time.Duration
}

// Creates a string representation of the result.
func (result benchmarkRunResult) String() string {
	return result.Config.String() + "/n-------\nRuntime: " + result.Runtime.String()
}
