package double_linkedlist

// Doubly linkedlist to manage the list
type List struct {
	Root   *Node
	Tail   *Node
	Length int
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
	NextNode *Node
	PrevNode *Node
}

func NewNode(val int, p *Node, n *Node) *Node {
	return &Node{
		Value:    val,
		PrevNode: p,
		NextNode: n,
	}
}

// func (node *Node) Empty() {
// 	node.NextNode = nil
// 	node.PrevNode = nil
// 	node.Value = 0
// }

func (node *Node) Push(val int) *Node {
	node.NextNode = NewNode(val, node, nil)
	return node.NextNode
}

func (node *Node) Append(val int) *Node {
	node.PrevNode = NewNode(val, nil, node)
	return node.PrevNode
}

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

func (list *List) Push(val int) {
	if list.Tail == nil {
		node := NewNode(val, nil, nil)

		list.Root = node
		list.Tail = node

		return
	}

	list.Tail.Push(val)
	list.Tail = list.Tail.NextNode
	list.Length++
}

func (list *List) Append(val int) {
	list.Root.Append(val)
	list.Root = list.Root.PrevNode
	list.Length++
}

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

func (list *List) Search(val int) (*Node, bool) {
	node := list.Root
	for {
		if node == nil {
			return nil, false
		}

		if node.Value == val {
			return node, true
		}

		node = node.NextNode
	}
}

func (list *List) GetAll() (result []*Node) {
	node := list.Root
	for {
		if node == nil {
			return result
		}

		result = append(result, node)
		node = node.NextNode
	}
}
