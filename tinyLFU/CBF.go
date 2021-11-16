package tinyLFU

import "errors"

var (
	errNotExisted = errors.New("error not exist")
)

// CBF abbrv: Counting Bloom Filter.
// Its job is to mantain the frequency of the operation and select the victim.
type CBF struct {
	Entries map[int]int
	LFU     *LFU
}

func NewCounter() *CBF {
	return &CBF{Entries: make(map[int]int)}
}

func (filter *CBF) Evaluate(key string, val []byte) (bool, error) {
	return false, nil
}

// The Estimate method is performed by calculating k different hash values for
// the key. Each hash value is treated as an index, and the counter at that index is read.
// Returns the victim's hash (smallest value).
func (filter *CBF) Estimate(key string) (hashValue int, err error) {
	hash32 := Hash32([]byte(key))
	hash64 := Hash64([]byte(key))

	val32, _ := filter.RetrieveHashValue(hash32)
	val64, _ := filter.RetrieveHashValue(hash64)

	if (val32 > val64) && (val32 > -1) {
		return val32, nil
	}

	return val64, err
}

// The Add method also calculates k different hash values for the key.
// However, it reads all k counters and only increments the minimal counters
func (filter *CBF) Add(key string, val []byte) {
}

func (filter *CBF) RetrieveHashValue(hash int) (int, error) {
	if val, ok := filter.Entries[hash]; ok {
		return val, nil
	}

	return -1, errNotExisted
}

// Once this counter reaches the sample size (W), we divide it and all other counters in the
// approximation sketch by SKETCH
// Implement using bitwise operation. ~O(n)
// We also clear the Doorkeeper in addition to halving all counters in the main structure.
func (filter *CBF) Refresh() {}
