

# concurrent_map

- sync.Map
  - 特点：
    1. 适用于大量读、少量写的场景（变化较少的缓存）
    2. 读取、写入和覆盖map不相交的键集合
  - 缺点
    - 不适用与大量读写的场景
- Map 加锁
  - 优点
    - 结构简单，适用于所有读写场景
  - 缺点
    - 采用全局锁，性能较差
- concurrent_map
  - 特点
    - 采用分段锁的概念
    - 内部采用读写锁(sync.RWMutex)
    - 内部结构
    ```Golang
    var SHARD_COUNT =32
    type ConcurrentMapShared struct {
     items map[string]interface{}
     sync.RWMutex
    
    }  
    ```
      