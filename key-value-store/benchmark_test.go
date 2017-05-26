package key_value_store

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func Test_BenchmarkConfigurationProvidesSensibleDefaultValues(t *testing.T) {
	config := NewBenchmarkRunConfiguration()

	assert.Condition(t, greaterThan(0.0, float64(config.numberOfOperations)))
	assert.Condition(t, greaterThan(0.0, float64(config.parallelOperations)))
	assert.Condition(t, greaterThanOrEqual(0.0, float64(config.writeOperationRatio)))
	assert.Condition(t, greaterThanOrEqual(float64(config.writeOperationRatio), 1.0))
}

func Test_BenchmarkConfigurationCanBeConfigured(t *testing.T) {
	config := NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(10).WriteOperationRatio(0.5)

	assert.Equal(t, 100, config.numberOfOperations)
	assert.Equal(t, 10, config.parallelOperations)
	assert.Equal(t, 0.5, config.writeOperationRatio)
}

func Test_BenchmarkConfigurationRejectsInvalidRatioValue(t *testing.T) {
	assert.Panics(t, func () {
		NewBenchmarkRunConfiguration().WriteOperationRatio(2.5)
	})
}

func Test_BenchmarkConfigurationCanBeConvertedToString(t *testing.T) {
	config := NewBenchmarkRunConfiguration()

	assert.NotEmpty(t, string(config))
}

func Test_BenchmarkExecutesCorrectOfOperationsInCaseOfSequentialExecution(t *testing.T) {
	store := newOperationCountingStore()
	benchmark := NewBenchmark(store)

	benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(1))

	assert.Equal(t, 100, store.NumberOfOperations())
}

func Test_BenchmarkExecutesCorrectNumberOfOperationsInCaseOfParallelExecution(t *testing.T) {
	store := newOperationCountingStore()
	benchmark := NewBenchmark(store)

	benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(10))

	assert.Equal(t, 100, store.NumberOfOperations())
}

func Test_BenchmarkReturnsValidResult(t *testing.T) {
	store := newOperationCountingStore()
	benchmark := NewBenchmark(store)

	result := benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(10))

	assert.Condition(t, greaterThan(0.0, result.Runtime.Seconds()))
}

func Test_ParallelExecutionIsFasterThanSequential(t *testing.T) {
	store := newOperationCountingStore()
	store.SetDelay(time.Duration(10) * time.Millisecond)
	benchmark := NewBenchmark(store)

	sequential := benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(1))
	parallel := benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(100).ParallelOperations(10))

	assert.Condition(t, greaterThan(parallel.Runtime.Seconds(), sequential.Runtime.Seconds()))
}

func Test_PerformsProvidedNumberOfWriteOperations(t *testing.T) {
	store := newOperationCountingStore()
	benchmark := NewBenchmark(store)

	benchmark.run(NewBenchmarkRunConfiguration().NumberOfOperations(10000).WriteOperationRatio(0.4))

	// The store might work with random heuristics. Therefore, a bigger
	// number of operations and a tolerance zone is used in this test.
	assert.InDelta(t, 0.4, store.WriteRatio(), 0.1)
}

func Test_BenchmarkResultCanBeConvertedToString(t *testing.T) {
	config := NewBenchmarkRunConfiguration()

	result := benchmarkRunResult{
		Runtime: time.Duration(5) * time.Second,
		Config: config,
	}

	assert.NotEmpty(t, string(result))
}


// Comparison function that is used to assert that actual is greater than expected.
func greaterThan(expected float64, actual float64) assert.Comparison {
	return assert.Comparison(func () bool {
		return actual > expected
	})
}

// Comparison function that is used to assert that actual is greater than or equal to expected.
func greaterThanOrEqual(expected float64, actual float64) assert.Comparison {
	return assert.Comparison(func () bool {
		return actual >= expected
	})
}

type operationCountingStore struct {
	delay time.Duration
	read int
	write int
}

// Creates a key-value store mock that counts the number of read/write operations.
func newOperationCountingStore() (*operationCountingStore) {
	return &operationCountingStore{
		delay: time.Duration(0),
		read: 0,
		write: 0,
	}
}

func (store *operationCountingStore) SetDelay(duration time.Duration) {
	store.delay = duration
}

func (store *operationCountingStore) Get(key string) string {
	time.Sleep(store.delay)
	store.read++
	// Return a dummy value, it is not important here.
	return ""
}

func (store *operationCountingStore) Set(key string, value string) {
	time.Sleep(store.delay)
	store.write++
}

// Returns the whole number of executed operations.
func (store *operationCountingStore) NumberOfOperations() int {
	return store.read + store.write
}

// Ratio of writes compared to reads.
// Returns a value between 0.0 and 1.0.
func (store *operationCountingStore) WriteRatio() float32 {
	return float32(store.write) / float32(store.NumberOfOperations())
}
