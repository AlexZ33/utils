package arrays

import (
	"fmt"
	"sync"
)

// 并发读写
// he map type in Go doesn't support concurrent reads and writes. concurrent-map provides a high-performance solution to this by sharding the map with minimal time spent waiting for locks.
// Prior to Go 1.9, there was no concurrent map implementation in the stdlib. In Go 1.9, sync.Map was introduced. The new sync.Map has a few key differences from this map. The stdlib sync.Map is designed for append-only scenarios. So if you want to use the map for something more like in-memory db, you might benefit from using our version. You can read more about it in the golang repo, for example here and here

var SHARD_COUNT = 32

type Stringer interface {
	fmt.Stringer
	comparable
}

// A "thread" safe map of type string:Anything
// To avoid lock bottlenecks this map is dived to severel (SHARD_COUNT) map shards
type ConcurrentMap[K comparable, V any] struct {
	shards   []*ConcurrentMapShared[K, V]
	sharding func(key K) uint32
}

// A "thread“ safe string to anything map
type ConcurrentMapShared[K comparable, V any] struct {
	items        map[K]V
	sync.RWMutex // Read Write mutex, guards access to internal map
}

func create[K comparable, V any](sharding func(key K) uint32) ConcurrentMap[K, V] {
	m := ConcurrentMap[K, V]{
		sharding: sharding,
		shards:   make([]*ConcurrentMapShared[K, V], SHARD_COUNT),
	}
	for i := 0; i < SHARD_COUNT; i++ {

	}
}
