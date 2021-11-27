package lfu

import (
	"log"
)

// Doubly linkedlist to manage the list
type List struct {
	Length int
	Root   *Node
	Tail   *Node
}

func NewList() *List {
	return &List{
		Root:   nil,
		Tail:   nil,
		Length: 0,
	}
}

type Node struct {
	Value    int
	Key      int
	NextNode *Node
	PrevNode *Node

	NodeList      *List
	FrequencyNode *Node
}

func NewNode(key int, val int, p *Node, n *Node) *Node {
	return &Node{
		Key:      key,
		Value:    val,
		PrevNode: p,
		NextNode: n,
	}
}

// Push to end.
func (node *Node) Push(key int, val int) *Node {
	if node == nil {
		panic("node.Push: node is nil")
	}
	node.NextNode = NewNode(key, val, node, nil)
	return node.NextNode
}

// Append to first
func (node *Node) Append(key int, val int) *Node {
	if node == nil {
		panic("node.Append: node is nil")
	}
	node.PrevNode = NewNode(key, val, nil, node)
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

// Push the node to the end of the list
func (list *List) Push(key int, val int) *Node {
	if list.Tail == nil {
		node := NewNode(key, val, nil, nil)
		list.Root = node
		list.Tail = node
		list.Length = 1
		return node
	}

	list.Tail = list.Tail.Push(key, val)
	list.Length++

	return list.Tail
}

// Append the front of the list
func (list *List) Append(key, val int) *Node {
	if list.Root == nil {
		node := NewNode(key, val, nil, nil)
		list.Root = node
		list.Tail = node
		list.Length = 1
		return node
	}
	list.Root = list.Root.Append(key, val)
	list.Length++
	return list.Root
}

// Pop the given node out of list
func (list *List) Pop(node *Node) {
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
	frequencyList *List
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

	val := node.Value
	freqNode := node.FrequencyNode

	if freqNode == nil {
		log.Panic("panic key: ", key)
	}

	if freqNode.NextNode == nil {
		freqNode.NextNode = NewNode(0, freqNode.Value+1, freqNode, nil)
		freqNode.NextNode.NodeList = NewList()
	}

	if freqNode.NextNode.Value == freqNode.Value+1 {
		newNode := node.FrequencyNode.NextNode.NodeList.Push(key, val)
		newNode.FrequencyNode = freqNode.NextNode

		this.memory[key] = newNode

	} else {
		nxtNode := freqNode.NextNode
		freqNode.NextNode = NewNode(0, freqNode.Value+1, freqNode, nxtNode)
		nxtNode.PrevNode = freqNode.NextNode

		freqNode.NextNode.NodeList = NewList()
		newNode := freqNode.NextNode.NodeList.Push(key, val)
		newNode.FrequencyNode = freqNode.NextNode

		this.memory[key] = newNode
	}

	freqNode.NodeList.Pop(node)
	if freqNode.NodeList.Length == 0 {
		this.frequencyList.Pop(freqNode)
	}

	return val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	_, ok := this.memory[key]
	// Not existed -> Add
	if !ok {
		if this.len == this.capacity {
			currentFreqList := this.frequencyList.Root.NodeList
			outdatedKey := currentFreqList.Root.Key
			currentFreqList.Pop(currentFreqList.Root)
			this.len--

			delete(this.memory, outdatedKey)

			if currentFreqList.Length == 0 {
				this.frequencyList.Pop(this.frequencyList.Root)
			}
		}

		if this.frequencyList.Root == nil {
			freqNode := this.frequencyList.Push(0, 1)
			freqNode.NodeList = NewList()

			newNode := freqNode.NodeList.Push(key, value)
			newNode.FrequencyNode = freqNode

			this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value != 1 {
			this.frequencyList.Append(0, 1)
			this.frequencyList.Root.NodeList = NewList()

			newNode := this.frequencyList.Root.NodeList.Push(key, value)
			newNode.FrequencyNode = this.frequencyList.Root
			newNode.Key = key

			this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value == 1 {
			newNode := this.frequencyList.Root.NodeList.Push(0, value)
			newNode.FrequencyNode = this.frequencyList.Root
			newNode.Key = key

			this.memory[key] = newNode
			this.len++

			return
		}

	} else {
		// move to the next frequency node
		_ = this.Get(key)

		node := this.memory[key]
		node.Value = value
	}
}

func (this *LFUCache) Print() {
	currentFNode := this.frequencyList.Root

	for currentFNode != nil {
		log.Println(currentFNode.Value)
		currentNodeIL := currentFNode.NodeList.Root

		for currentNodeIL != nil {
			log.Print(" ", currentNodeIL.Value)
			currentNodeIL = currentNodeIL.NextNode
		}
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
