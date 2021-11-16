package lfu

import (
	"errors"
)

// This implementation aims not to used pointer to avoid GC overhead performance.
type LFUS struct {

	// Memory Size is limited
	Memory map[string]interface{}

	// Frequency for each values of memory.
	Frq map[string]int

	// Queue for FIFO
	Queue    []string
	lenQueue int
}

func NewLFUS() *LFU {
	return &LFU{
		Memory:   make(map[uint32]interface{}),
		Frq:      make(map[uint32]int),
		lenQueue: 0,
	}
}

func (l *LFUS) isExisted(hash string) bool {
	return l.Memory[hash] != nil
}

func (l *LFUS) isEvicted() bool {
	return len(l.Queue) > size-1
}

func (l *LFUS) removeLeastFrequent() {
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

func (l *LFUS) add(key string, val interface{}) error {
	passGate := false

	hash := hash_string(key)

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
	l.Queue = append([]string{hash}, l.Queue...)

	return nil
}

// queuePos is the position of value in Queue, without specific equals -1
func (l *LFUS) remove(key string, queuePos int) {
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

func (l *LFUS) get(key string) (interface{}, error) {
	hash := hash_string(key)

	if l.isExisted(hash) {
		l.Frq[hash]++
		return l.Memory[hash], nil
	}

	return nil, errors.New("do not exist")
}

func (l *LFUS) hit(key string) {
	l.Frq[hash_string(key)]++
}
