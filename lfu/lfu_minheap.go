package lfu

import (
	"algo/heap"
	"sync"
)

type EvictedItem struct {
	Key   string
	Value interface{}
}

// This implementation aims not to used pointer to avoid GC overhead performance.
type LFUCache struct {
	// The size of the Cache
	limit       int
	currentSize int

	// Number of items each eviction will handle
	evictionQuantity int

	// Memory Size is limited
	memory map[uint64]interface{}

	// Frequency for each values of memory.
	frq *heap.MinHeapTree

	lock *sync.Mutex
	// evictionChannel chan<- EvictedItem
}

func NewLFUCache(NumItem int) *LFUCache {
	return &LFUCache{
		memory:      make(map[uint64]interface{}),
		frq:         heap.NewHeapTree(),
		lock:        new(sync.Mutex),
		limit:       NumItem,
		currentSize: 0,
	}
}

func (cache *LFUCache) Set(Key string, Value interface{}) error {
	hashKey := hash64(Key)
	cache.frq.Push(1, hashKey)
	cache.memory[hashKey] = Value
	cache.currentSize++

	if cache.currentSize == cache.limit {
		cache.Evict(cache.evictionQuantity)
	}

	return nil
}

func (cache *LFUCache) Get(Key string) interface{} {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	hashKey := hash64(Key)
	val, ok := cache.memory[hashKey]
	if !ok {
		return nil
	}

	cache.frq.Increment(hashKey)
	return val
}

func (cache *LFUCache) Evict(quantity int) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for i := 0; i < quantity; i++ {
		hash := cache.frq.Pop()
		delete(cache.memory, hash.(uint64))
	}
}
