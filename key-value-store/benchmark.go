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

// Creates a new benchmark runner for the given store.
func NewBenchmark(store SimpleKeyValueStore) *benchmark {
	return &benchmark{
		store: store,
	}
}

// Runs a benchmark against the store.
//
// The provided configuration defines number of operations, parallelism etc.
func (benchmark *benchmark) Run(config BenchmarkRunConfiguration) benchmarkRunResult {
	operations := make(chan func (store SimpleKeyValueStore), config.parallelOperations)
	workers := newCoordinator()
	workers.Register(func () {
		benchmark.produce(operations, config.numberOfOperations, config.writeOperationRatio)
	})
	// Prepare several workers, depending on the number of parallel operations.
	for i := 0; i < config.parallelOperations; i++ {
		workers.Register(func () {
			benchmark.consume(operations)
		})
	}
	start := time.Now()
	workers.Run()
	// Wait for all workers to complete.
	workers.Wait()
	return benchmarkRunResult{
		Config: config,
		Runtime: time.Since(start),
	}
}

// Produced the requested number of operations and writes them to the channel.
func (benchmark *benchmark) produce(operations chan func (store SimpleKeyValueStore), numberOfOperations int, writeOperationRatio float32) {
	// Use random number generator with fixed seed to ensure that runs are deterministic.
	random := rand.New(rand.NewSource(0))
	// Returns a character depending on the given index:
	// 0 = A
	// 1 = B
	// ...
	toChar := func (i int) string {
		return string('A' + i)
	}
	// Returns a random key in the range A..Z
	getKey := func () string {
		return toChar(random.Intn(26))
	}
	// Push the configured number of operations to the channel.
	// At the same time, the workers start to read and execute the operations.
	for i := 0; i < numberOfOperations; i++ {
		if (random.Float32() < writeOperationRatio) {
			operations <- func (store SimpleKeyValueStore) {
				store.Set(getKey(), "hello world")
			}
		} else {
			operations <- func (store SimpleKeyValueStore) {
				store.Get(getKey())
			}
		}
	}
	// Close the channel to ensure that the consumers will terminate when the channel is empty.
	close(operations)
}

// Reads operations from the given channel and applies them to the store.
func (benchmark *benchmark) consume(operations chan func (store SimpleKeyValueStore)) {
	for operation := range operations {
		operation(benchmark.store)
	}
}

// Coordinates multiple worker functions.
type coordinator struct {
	workers    []func()
	terminated *sync.WaitGroup
}

func newCoordinator() *coordinator {
	return &coordinator{
		workers: make([]func(), 0),
		terminated: &sync.WaitGroup{},
	}
}

// Registers a worker function.
func (coordinator *coordinator) Register(worker func()) {
	coordinator.terminated.Add(1)
	coordinator.workers = append(coordinator.workers, func () {
		worker()
		coordinator.terminated.Done()
	})
}

// Starts all workers.
func (coordinator *coordinator) Run() {
	for _, worker := range coordinator.workers {
		go worker()
	}
}

// Waits until all workers have finished. Must call Run() first.
func (coordinator *coordinator) Wait() {
	coordinator.terminated.Wait()
}

// Encapsulates a configuration for a benchmark run.
type BenchmarkRunConfiguration struct {
	numberOfOperations  int
	parallelOperations  int
	writeOperationRatio float32
}

// Creates a new benchmark configuration with default values.
func NewBenchmarkRunConfiguration() BenchmarkRunConfiguration {
	return BenchmarkRunConfiguration{
		numberOfOperations:  1000,
		parallelOperations:  1,
		writeOperationRatio: 0.2,
	}
}

// Sets the number of operations that will be executed and returns the new configuration.
func (config BenchmarkRunConfiguration) NumberOfOperations(value int) BenchmarkRunConfiguration {
	config.numberOfOperations = value
	return config
}

// Sets the number of operations that will be executed in parallel and returns the new configuration.
func (config BenchmarkRunConfiguration) ParallelOperations(value int) BenchmarkRunConfiguration {
	config.parallelOperations = value
	return config
}

// Defines the number of write operations among the whole number of operations.
// value must be a float between 0.0 and 1.0.
func (config BenchmarkRunConfiguration) WriteOperationRatio(value float32) BenchmarkRunConfiguration {
	if value < 0.0 || value > 1.0 {
		panic("Write ration value must be in range [0.0..1.0]")
	}
	config.writeOperationRatio = value
	return config
}

// Creates a string representation of the configuration.
func (config BenchmarkRunConfiguration) String() string {
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
	Config  BenchmarkRunConfiguration
	// The number of seconds the benchmark took to run.
	Runtime time.Duration
}

// Returns the average runtime per operation.
func (result benchmarkRunResult) AvgRuntime() time.Duration {
	return result.Runtime / time.Duration(result.Config.numberOfOperations)
}

// Creates a string representation of the result.
func (result benchmarkRunResult) String() string {
	return fmt.Sprintf(`%s
	Runtime:
	- All: %s
	- Avg. per operation: %s
	`, result.Config.String(), result.Runtime.String(), result.AvgRuntime())
}
