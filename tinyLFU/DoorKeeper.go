package tinyLFU

// Def: a regular Bloom filter placed in front of the a
// Upon item arrival, we first check if the item is contained in the Doorkeeper.
// If it is not contained in the Doorkeeper (as is expected with first timers and tail items),
// the item is inserted to the Doorkeeper and otherwise, it is inserted to the main structure.
// When querying items, we use both the Doorkeeper and the main structures.
// That is, if the item is included in the Doorkeeper, TinyLFU estimates the frequency of this item as its estimation in the main structure plus 1.
// Otherwise, TinyLFU returns just the estimation from the main structure
type DoorKeeper struct {
	Filter map[int]int
}

func NewDoorKeeper() *DoorKeeper {
	return &DoorKeeper{
		Filter: make(map[int]int),
	}
}

// Check the existence of the key in DoorKeeper.
// Return true if existed.
func (t *DoorKeeper) Add(key string, value []byte) bool {
	hash_val := Hash32(value)
	hash_key := Hash32([]byte(key))

	if val, ok := t.Filter[hash_key]; ok {
		if hash_val == val {
			return true
		}
	}

	t.Filter[hash_key] = hash_val
	return false
}
