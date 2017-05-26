package key_value_store

type benchmark struct {
	store SimpleKeyValueStore
	operations chan func(SimpleKeyValueStore)
}

// Encapsulates a configuration for a benchmark run.
type benchmarkRunConfiguration struct {
	numberOfOperations int
	parallelOperations int
	writeOperationRatio float32
}

// Creates a new benchmark configuration with default values.
func NewBenchmarkRunConfiguration() benchmarkRunConfiguration {

}

// Sets the number of operations that will be executed and returns the new configuration.
func (config benchmarkRunConfiguration) NumberOfOperations(value int) benchmarkRunConfiguration {

}

// Sets the number of operations that will be executed in parallel and returns the new configuration.
func (config benchmarkRunConfiguration) ParallelOperations(value int) benchmarkRunConfiguration {

}

// Defines the number of write operations among the whole number of operations.
// value must be a float between 0.0 and 1.0.
func (config benchmarkRunConfiguration) WriteOperationRatio(value float32) benchmarkRunConfiguration {

}

func NewBenchmark(store SimpleKeyValueStore) *benchmark {

}

// Runs a benchmark against the store.
//
func (benchmark *benchmark) run(config benchmarkRunConfiguration) float64 {

}
