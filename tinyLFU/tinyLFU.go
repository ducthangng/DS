package tinyLFU

const (
	SIZE   = 1024 * 1024
	W_SIZE = 1024
	SKETCH = 2
)

// TinyLFU implementation
// Eviction process: the cache picks a victim, and the TinyLFU decides if replace
// the victim with the new item will increase the hit-raio.
type TinyLFU struct {
	// Frontline of the tiny LFU
	DoorKeeper *DoorKeeper

	// Approximate counting scheme with bloom filter.
	Counter *CBF

	MemSize     int
	CurrentSize int
}

func NewTinyLFU(MemSize int) *TinyLFU {
	return &TinyLFU{
		DoorKeeper:  NewDoorKeeper(),
		Counter:     NewCounter(),
		MemSize:     MemSize,
		CurrentSize: 0,
	}
}

func (t *TinyLFU) Set(key string, val []byte) {
	// Check if DoorKeeper proceed
	if ok := t.DoorKeeper.Add(key, val); ok {
	}
}

func (t *TinyLFU) Get(key string) []byte {
	return nil
}
