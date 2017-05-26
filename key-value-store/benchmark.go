package key_value_store

type benchmark struct {
	store SimpleKeyValueStore
	operations chan func(SimpleKeyValueStore)
}


func NewBenchmark(store SimpleKeyValueStore) *benchmark {

}

// Runs a benchmark against the store.
//
func (benchmark *benchmark) run(numberOfOperations int, parallelOperations int, writeOperationRatio float32) float64 {

}
