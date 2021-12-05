package lfu

// FrequencyList is a double_linkedlist with external NodeList for each node
type FrequencyList struct {
	Length int
	Root   *FrequencyNode
	Tail   *FrequencyNode
}

func NewList() *FrequencyList {
	return &FrequencyList{
		Root:   nil,
		Tail:   nil,
		Length: 0,
	}
}

type FrequencyNode struct {
	// Hold the frequency value
	Value    int
	NextNode *FrequencyNode
	PrevNode *FrequencyNode
	NodeList *NodeList
}

func NewFrequencyNode(val int, p *FrequencyNode, n *FrequencyNode) *FrequencyNode {
	return &FrequencyNode{
		Value:    val,
		PrevNode: p,
		NextNode: n,
		NodeList: NewNodeList(),
	}
}

// Push to end.
func (node *FrequencyNode) Push(val int) *FrequencyNode {
	if node == nil {
		panic("frequencyNode.Push: node is nil")
	}

	node.NextNode = NewFrequencyNode(val, node, nil)
	return node.NextNode
}

// Append to first
func (node *FrequencyNode) PushFront(val int) *FrequencyNode {
	if node == nil {
		panic("frequencyNode.PushFront: node is nil")
	}

	node.PrevNode = NewFrequencyNode(val, nil, node)
	return node.PrevNode
}

// Pop is pop the node out of the linkedlist, and empty the node for GC.
func (node *FrequencyNode) Pop() {
	prv := node.PrevNode
	nxt := node.NextNode

	if prv != nil {
		prv.NextNode = nxt
	}

	if nxt != nil {
		nxt.PrevNode = prv
	}

	node = nil
}

func (list *FrequencyList) Push(val int) *FrequencyNode {
	if list.Tail == nil {
		node := NewFrequencyNode(val, nil, nil)
		list.Root = node
		list.Tail = node
		list.Length = 1
		return node
	}

	list.Tail = list.Tail.Push(val)
	list.Length++

	return list.Tail
}

func (list *FrequencyList) PushFirst(val int) *FrequencyNode {
	if list.Root == nil {
		node := NewFrequencyNode(val, nil, nil)
		list.Root = node
		list.Tail = node
		list.Length = 1
		return node
	}

	list.Root = list.Root.PushFront(val)
	list.Length++

	return list.Root
}

// Pop the given node out of list
func (list *FrequencyList) Pop(node *FrequencyNode) {
	if node == list.Tail {
		list.Tail = node.PrevNode
	}

	if node == list.Root {
		list.Root = node.NextNode
	}

	node.Pop()
	list.Length--
}

// NodeList is a double_linkedlist
type NodeList struct {
	Length int
	Root   *Node
	Tail   *Node
}

func NewNodeList() *NodeList {
	return &NodeList{
		Root:   nil,
		Tail:   nil,
		Length: 0,
	}
}

type Node struct {
	Key           int
	Value         int
	NextNode      *Node
	PrevNode      *Node
	FrequencyNode *FrequencyNode
}

func NewNode(key int, val int, p *Node, n *Node, f *FrequencyNode) *Node {
	return &Node{
		Key:           key,
		Value:         val,
		FrequencyNode: f,
		PrevNode:      p,
		NextNode:      n,
	}
}

// Push to end.
func (node *Node) Push(key int, val int, f *FrequencyNode) *Node {
	if node == nil {
		panic("frequencyNode.Push: node is nil")
	}

	node.NextNode = NewNode(key, val, node, nil, f)
	return node.NextNode
}

// Append to first
func (node *Node) PushFront(key int, val int, f *FrequencyNode) *Node {
	if node == nil {
		panic("frequencyNode.PushFront: node is nil")
	}

	node.PrevNode = NewNode(key, val, nil, node, f)
	return node.PrevNode
}

// Pop is pop the node out of the linkedlist, and empty the node for GC.
func (node *Node) Pop() {
	prv := node.PrevNode
	nxt := node.NextNode

	if prv != nil {
		prv.NextNode = nxt
	}

	if nxt != nil {
		nxt.PrevNode = prv
	}

	node = nil
}

func (list *NodeList) Push(key, val int, freqNode *FrequencyNode) *Node {
	if list.Tail == nil {
		node := NewNode(key, val, nil, nil, freqNode)
		list.Root = node
		list.Tail = node
		list.Length = 1

		return node
	}

	list.Tail = list.Tail.Push(key, val, freqNode)
	list.Length++

	return list.Tail
}

func (list *NodeList) PushFront(key, val int, freqNode *FrequencyNode) *Node {
	if list.Root == nil {
		node := NewNode(key, val, nil, nil, freqNode)
		list.Root = node
		list.Tail = node
		list.Length = 1

		node.FrequencyNode = freqNode
		return node
	}

	list.Root = list.Root.PushFront(key, val, freqNode)
	list.Length++

	return list.Root
}

// Pop the given node out of list
func (list *NodeList) Pop(node *Node) {
	if node == list.Tail {
		list.Tail = node.PrevNode
	}

	if node == list.Root {
		list.Root = node.NextNode
	}

	node.Pop()
	list.Length--
}

type LFUCache struct {
	capacity int
	len      int

	memory        map[int]*Node
	frequencyList *FrequencyList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		len:           0,
		capacity:      capacity,
		memory:        make(map[int]*Node),
		frequencyList: NewList(),
	}
}

// Set at bottom
func (this *LFUCache) Get(key int) int {
	node, ok := this.memory[key]
	if !ok {
		return -1
	}

	// val := node.Value
	// freqNode := node.FrequencyNode

	if node.FrequencyNode.NextNode == nil {
		node.FrequencyNode.NextNode = NewFrequencyNode(node.FrequencyNode.Value+1, node.FrequencyNode, nil)
		node.FrequencyNode.NextNode.NodeList = NewNodeList()
	}

	if node.FrequencyNode.NextNode.Value == node.FrequencyNode.Value+1 {
		this.memory[key] = node.FrequencyNode.NextNode.NodeList.Push(key, node.Value, node.FrequencyNode.NextNode)

		// this.memory[key] = newNode
	} else {
		nxtNode := node.FrequencyNode.NextNode
		node.FrequencyNode.NextNode = NewFrequencyNode(node.FrequencyNode.Value+1, node.FrequencyNode, nxtNode)
		nxtNode.PrevNode = node.FrequencyNode.NextNode

		node.FrequencyNode.NextNode.NodeList = NewNodeList()
		this.memory[key] = node.FrequencyNode.NextNode.NodeList.Push(key, node.Value, node.FrequencyNode.NextNode)

		// this.memory[key] = newNode
	}

	// log.Println("add: ", key, " value: ", node.Value)

	node.FrequencyNode.NodeList.Pop(node)
	if node.FrequencyNode.NodeList.Length == 0 {
		// log.Println("pop freq: ", freqNode.Value)
		this.frequencyList.Pop(node.FrequencyNode)
	}

	return node.Value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	node, ok := this.memory[key]

	// Not existed -> Add
	if !ok {
		// full capacity -> we choose the first node and remove it.
		if this.len == this.capacity {
			outdatedKey := this.frequencyList.Root.NodeList.Root.Key
			this.frequencyList.Root.NodeList.Pop(this.frequencyList.Root.NodeList.Root)
			this.len--

			// log.Println("delete: ", outdatedKey, " value: ", firstNode.Value)
			delete(this.memory, outdatedKey)

			if this.frequencyList.Root.NodeList.Length == 0 {
				// log.Println("pop freq: ", this.frequencyList.Root.Value)
				this.frequencyList.Pop(this.frequencyList.Root)
			}
		}

		if this.frequencyList.Root == nil {
			freqNode := this.frequencyList.Push(1)
			freqNode.NodeList = NewNodeList()

			this.memory[key] = freqNode.NodeList.Push(key, value, freqNode)
			// log.Println("add: ", key, " value: ", newNode.Value)
			// this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value != 1 {
			freqNode := this.frequencyList.PushFirst(1)
			freqNode.NodeList = NewNodeList()

			this.memory[key] = freqNode.NodeList.Push(key, value, freqNode)
			// log.Println("add: ", key, " value: ", newNode.Value)
			// this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value == 1 {
			this.memory[key] = this.frequencyList.Root.NodeList.Push(key, value, this.frequencyList.Root)
			// log.Println("add: ", key, " value: ", newNode.Value)
			// this.memory[key] = newNode
			this.len++

			return
		}

	} else {
		node.Value = value
		// move to the next frequency node
		_ = this.Get(key)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
