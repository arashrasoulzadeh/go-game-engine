package models

type ObjectPool struct {
	pools    map[string][]interface{}
	maxItems int
	Index    int
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pools:    make(map[string][]interface{}),
		maxItems: 8, // Set the maximum number of items per key
		Index:    1,
	}
}

// Append adds an object to the specified pool, rotating if the pool exceeds maxItems
func (op *ObjectPool) Append(poolName string, value interface{}) {
	// Initialize the slice if it doesn't exist
	op.Index++
	if op.pools[poolName] == nil {
		op.pools[poolName] = []interface{}{}
	}

	// If the pool has reached the maxItems limit, remove the oldest item
	if len(op.pools[poolName]) >= op.maxItems {
		op.pools[poolName] = op.pools[poolName][1:] // Remove the first (oldest) item
	}

	// Append the new value to the pool
	op.pools[poolName] = append(op.pools[poolName], value)
}

// Get retrieves all objects from the specified pool
func (op *ObjectPool) Get(poolName string) []interface{} {
	return op.pools[poolName]
}
