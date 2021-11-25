package heap

type Node struct {
	Hash  interface{}
	Value int
}

func NewNode(value int, hash interface{}) Node {
	return Node{
		Value: value,
		Hash:  hash,
	}
}

type MinHeapTree struct {
	Entries      []Node
	ToBeInserted int
}

func NewHeapTree() *MinHeapTree {
	return &MinHeapTree{
		Entries:      make([]Node, 2),
		ToBeInserted: 1,
	}
}

func (t *MinHeapTree) Swap(i, j int) {
	node := t.Entries[i]
	t.Entries[i] = t.Entries[j]
	t.Entries[j] = node
}

func (t *MinHeapTree) Heapify(pos int) {
	if pos == 0 {
		return
	}

	parent := 0
	if pos%2 == 0 {
		parent = (pos) / 2
	} else {
		parent = (pos - 1) / 2
	}

	if parent > -1 {
		if t.Entries[parent].Value > t.Entries[pos].Value {
			t.Swap(parent, pos)
			// fmt.Printf("swap %v and %v\n", t.Entries[parent], t.Entries[pos])
		}

		if parent > 0 {
			t.Heapify(parent)
		}
	}
}

func (t *MinHeapTree) ReverseHeapify(root int) {
	l, r, swap, val := -1, -1, -1, 1<<20

	if root*2 < t.ToBeInserted {
		l = root * 2
	}

	if root*2+1 < t.ToBeInserted {
		r = root*2 + 1
	}

	if l != -1 {
		if t.Entries[root].Value > t.Entries[l].Value {
			swap = l
			val = t.Entries[l].Value
		}
	}

	if r != -1 {
		if (t.Entries[root].Value > t.Entries[r].Value) && (val > t.Entries[r].Value) {
			swap = r
		}
	}

	if swap != -1 {
		t.Swap(root, swap)
		// fmt.Printf("reverse: swap %v and %v\n", t.Entries[root], t.Entries[swap])
		t.ReverseHeapify(swap)
	}
}

func (t *MinHeapTree) Retrieve() []Node {
	return t.Entries[1:]
}

func (t *MinHeapTree) Push(val int, hash interface{}) {
	node := NewNode(val, hash)

	if t.ToBeInserted >= len(t.Entries) {
		t.Entries = append(t.Entries, make([]Node, t.ToBeInserted*2)...)
	}

	t.Entries[t.ToBeInserted] = node
	t.Heapify(t.ToBeInserted)

	t.ToBeInserted++
}

// Standard operation is to delete the root
func (t *MinHeapTree) Pop() (hash interface{}) {
	if len(t.Entries) == 0 {
		return nil
	}

	hash = t.Entries[1]
	t.Entries[1] = t.Entries[t.ToBeInserted-1]
	t.Entries = t.Entries[:t.ToBeInserted-1]

	t.ToBeInserted--
	t.ReverseHeapify(1)

	return hash
}

// GetPopList returns the lists of going-to-pop elements
func (t *MinHeapTree) GetPopList(quantity int) (list []interface{}) {
	if len(t.Entries) == 0 {
		return nil
	}

	if quantity > len(t.Entries) {
		quantity = len(t.Entries)
	}

	for i := 1; i <= quantity; i++ {
		list = append(list, t.Entries[i].Hash)
	}

	return list
}

// Improvisation
func (t *MinHeapTree) Increment(hashVal interface{}) {
	for i := 0; i < len(t.Entries); i++ {
		if t.Entries[i].Hash == hashVal {
			t.Entries[i].Value++
			t.ReverseHeapify(i)

			return
		}
	}
}
