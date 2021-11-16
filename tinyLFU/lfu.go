package tinyLFU

import "errors"

const (
	size = 2
)

// This implementation aims not to used pointer to avoid GC overhead performance.
type LFU struct {
	// Memory Size is limited
	Memory map[string]interface{}

	// Frequency for each values of memory.
	Frq map[string]int

	// Queue for storing hash value
	Queue []string
}

func NewLFU() *LFU {
	return &LFU{
		Memory: make(map[string]interface{}),
		Frq:    make(map[string]int),
	}
}

func (l *LFU) IsExisted(hash string) bool {
	return l.Memory[hash] != nil
}

func (l *LFU) IsFull() bool {
	return len(l.Queue) > size-1
}

func (l *LFU) GetVictim() (key string, err error) {
	least := 1 << 20
	rmTarget := -1

	for i := 0; i < len(l.Queue); i++ {
		// Last > equal
		if l.Frq[l.Queue[i]] <= least {
			rmTarget = i
			least = l.Frq[l.Queue[i]]
		}
	}

	//l.remove(l.Queue[rmTarget], rmTarget)
	return l.Queue[rmTarget], nil
}

func (l *LFU) Add(key string, val interface{}) error {
	if l.IsExisted(key) {
		l.Memory[key] = val
		l.Frq[key]++

		return nil
	}

	if l.IsFull() {
		panic("full")
	}

	l.Frq[key] = 1
	l.Memory[key] = val
	l.Queue = append([]string{key}, l.Queue...)

	return nil
}

// queuePos is the position of value in Queue, without specific equals -1
func (l *LFU) Remove(key string, queuePos int) {
	delete(l.Memory, key)
	delete(l.Frq, key)

	switch queuePos {
	case -1:
		state := 0
		for i := 0; i < len(l.Queue); i++ {
			if l.Queue[i] == key {
				state = i
				break
			}
		}

		l.Queue = append(l.Queue[:state], l.Queue[state:]...)

	default:
		l.Queue = append(l.Queue[:queuePos], l.Queue[queuePos:]...)
	}
}

func (l *LFU) Get(key string) (interface{}, error) {
	if l.IsExisted(key) {
		l.Frq[key]++
		return l.Memory[key], nil
	}

	return nil, errors.New("do not exist")
}

func (l *LFU) Hit(key string) {
	l.Frq[key]++
}
