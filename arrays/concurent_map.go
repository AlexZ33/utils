package arrays

import "fmt"

// 并发读写
// he map type in Go doesn't support concurrent reads and writes. concurrent-map provides a high-performance solution to this by sharding the map with minimal time spent waiting for locks.
// Prior to Go 1.9, there was no concurrent map implementation in the stdlib. In Go 1.9, sync.Map was introduced. The new sync.Map has a few key differences from this map. The stdlib sync.Map is designed for append-only scenarios. So if you want to use the map for something more like in-memory db, you might benefit from using our version. You can read more about it in the golang repo, for example here and here

var SHARD_COUNT = 32

type Stringer interface {
	fmt.Stringer
	comparable
}
