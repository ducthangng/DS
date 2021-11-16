package lfu

import (
	"errors"
)

const (
	size = 2
)

// This implementation aims not to used pointer to avoid GC overhead performance.
type LFU struct {
	// Memory Size is limited
	Memory map[uint32]interface{}

	// Frequency for each values of memory.
	Frq map[uint32]int

	// Queue for FIFO
	Queue    []uint32
	lenQueue int
}

func NewLFU() *LFU {
	return &LFU{
		Memory:   make(map[uint32]interface{}),
		Frq:      make(map[uint32]int),
		lenQueue: 0,
	}
}

func (l *LFU) isExisted(hash uint32) bool {
	return l.Memory[hash] != nil
}

func (l *LFU) isEvicted() bool {
	return len(l.Queue) > size-1
}

func (l *LFU) removeLeastFrequent() {
	least := 1 << 20
	rmTarget := 0

	for i := 0; i < l.lenQueue; i++ {
		// Last > equal
		if l.Frq[l.Queue[i]] <= least {
			rmTarget = i
			least = l.Frq[l.Queue[i]]
		}
	}

	l.remove(l.Queue[rmTarget], rmTarget)
}

func (l *LFU) add(key string, val interface{}) error {
	passGate := false

	hash := hash(key)

	if l.isExisted(hash) {
		l.remove(hash, -1)
		passGate = true
	}

	if l.isEvicted() && (!passGate) {
		l.removeLeastFrequent()
		passGate = true
	}

	l.lenQueue++
	l.Frq[hash] = 1
	l.Memory[hash] = val
	l.Queue = append([]uint32{hash}, l.Queue...)

	return nil
}

// queuePos is the position of value in Queue, without specific equals -1
func (l *LFU) remove(key uint32, queuePos int) {
	l.Memory[key] = interface{}(nil)
	l.Frq[key] = 0

	switch queuePos {
	case -1:
		state := false
		for i := 0; i < l.lenQueue-1; i++ {
			if l.Queue[i] == key {
				state = true
			}

			if state {
				l.Queue[i] = l.Queue[i+1]
			}
		}

	default:
		for i := queuePos; i < l.lenQueue-1; i++ {
			l.Queue[i] = l.Queue[i+1]
		}
	}

	l.lenQueue--
}

func (l *LFU) get(key string) (interface{}, error) {
	hash := hash(key)

	if l.isExisted(hash) {
		l.Frq[hash]++
		return l.Memory[hash], nil
	}

	return nil, errors.New("do not exist")
}

func (l *LFU) hit(key string) {
	l.Frq[hash(key)]++
}

func main() {
	lfu := NewLFU()

	lfu.add("ducthang", 18)
	lfu.add("vulong", 16)
	lfu.add("randomstring", 21)
	lfu.add("ducthang", 18)
	lfu.add("vulong", 16)

	lfu.hit("ducthang")
	lfu.hit("vulong")
	lfu.hit("vulong")
	lfu.hit("ducthang")
	lfu.hit("ducthang")
	lfu.hit("ducthang")

	lfu.add("randomstring", 10)
	_, err := lfu.get("vulong")
	if err != nil {
		panic(err)
	}

}
