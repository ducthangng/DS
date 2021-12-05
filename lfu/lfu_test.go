package lfu

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// RandomNumber returns a random number.
func RandomNumber(min int64, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

// RandomString returns a random string
func RandomString(len int) string {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomNumber(97, 122))
	}

	return string(bytes)
}

func TestLFU(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	lfus := NewLFU()
	// lfu := NewLFU()

	keys := []string{}

	for i := 0; i < 100; i++ {
		keys = append(keys, RandomString(8))
	}

	for i := 0; i < 10000; i++ {
		err := lfus.add(keys[i%100], keys[i%100])
		require.NoError(t, err)
	}

	for i := 0; i < 10000; i++ {
		val, err := lfus.get(keys[i%100])
		require.NoError(t, err)
		if val != nil {
			require.Equal(t, val.(string), keys[i%100])
		}
	}

}

func BenchmarkLFU_Int(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	lfus := NewLFU()
	keys := []string{}

	for i := 0; i < 100; i++ {
		keys = append(keys, RandomString(8))
	}

	for i := 0; i <= b.N; i++ {
		err := lfus.add(keys[i%100], "random_value_to_retrieve")
		if err != nil {
			log.Println(err) // return full?
		}

		_, _ = lfus.get(keys[(i+rand.Intn(100))%100])
	}
}

/*
goos: darwin
goarch: amd64
pkg: algo/lfu
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLFU_Int-12    	  341947	     79766 ns/op	  687923 B/op	       2 allocs/op
PASS
ok  	algo/lfu	28.159s
*/

/**

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

	val := node.Value
	freqNode := node.FrequencyNode

	if freqNode.NextNode == nil {
		freqNode.NextNode = NewFrequencyNode(freqNode.Value+1, freqNode, nil)
		freqNode.NextNode.NodeList = NewNodeList()
	}

	if freqNode.NextNode.Value == freqNode.Value+1 {
		newNode := freqNode.NextNode.NodeList.Push(key, val, freqNode.NextNode)

		this.memory[key] = newNode
	} else {
		nxtNode := freqNode.NextNode
		freqNode.NextNode = NewFrequencyNode(freqNode.Value+1, freqNode, nxtNode)
		nxtNode.PrevNode = freqNode.NextNode

		freqNode.NextNode.NodeList = NewNodeList()
		newNode := freqNode.NextNode.NodeList.Push(key, val, freqNode.NextNode)

		this.memory[key] = newNode
	}

	// log.Println("add: ", key, " value: ", node.Value)

	freqNode.NodeList.Pop(node)
	if freqNode.NodeList.Length == 0 {
		// log.Println("pop freq: ", freqNode.Value)
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
		// full capacity -> we choose the first node and remove it.
		if this.len == this.capacity {
			firstNode := this.frequencyList.Root.NodeList.Root
			outdatedKey := firstNode.Key
			this.frequencyList.Root.NodeList.Pop(firstNode)
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

			newNode := freqNode.NodeList.Push(key, value, freqNode)
			// log.Println("add: ", key, " value: ", newNode.Value)
			this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value != 1 {
			freqNode := this.frequencyList.PushFirst(1)
			freqNode.NodeList = NewNodeList()

			newNode := freqNode.NodeList.Push(key, value, freqNode)
			// log.Println("add: ", key, " value: ", newNode.Value)
			this.memory[key] = newNode
			this.len++

			return
		}

		if this.frequencyList.Root.Value == 1 {
			newNode := this.frequencyList.Root.NodeList.Push(key, value, this.frequencyList.Root)
			// log.Println("add: ", key, " value: ", newNode.Value)
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
**/
