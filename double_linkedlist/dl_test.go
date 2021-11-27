package double_linkedlist

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDL(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	list := NewList()
	list.Push(1)

	for i := 0; i < 10000; i++ {
		list.Push(i)
	}

	currentNode := list.Root
	count := 0
	for {
		x := rand.Intn(1)
		myBool := true
		if x == 1 {
			myBool = false
		}

		if myBool {
			list.Pop(currentNode)
			count++
		}

		currentNode = currentNode.NextNode
		if (currentNode.Value == 9999) && (currentNode.NextNode == nil) {
			break
		}
	}

	require.Equal(t, 10000-count, list.Length)

	// for i := 0; i < 1000; i++ {
	// 	require.Equal(t, currentNode.Value, i)
	// 	currentNodeInNodeList := currentNode.NodeList.Root
	// 	for j := 0; j < 100; j++ {
	// 		require.Equal(t, currentNodeInNodeList.Value, j)
	// 		currentNodeInNodeList = currentNodeInNodeList.NextNode
	// 	}

	// 	currentNode = currentNode.NextNode
	// }

	// currentNode := list.Root
	// for i := 0; i < 1000; i++ {
	// 	require.Equal(t, currentNode.Value, i)
	// 	currentNodeInNodeList := currentNode.NodeList.Root

	// 	checker := false
	// 	if i%2 == 0 {
	// 		checker = true
	// 	}

	// 	for j := 0; j < 100; j++ {
	// 		if checker {
	// 			currentNode.NodeList.Pop(currentNodeInNodeList)
	// 		}

	// 		currentNodeInNodeList = currentNode.NodeList.Root
	// 	}

	// 	if checker {
	// 		require.Equal(t, 0, currentNode.NodeList.Length)
	// 	}

	// 	currentNode = currentNode.NextNode
	// }

	// currentNode = list.Root
	// for i := 0; i < 1000; i++ {
	// 	// log.Println(currentNode.Value)
	// 	require.Equal(t, currentNode.Value, i)
	// 	currentNodeInNodeList := currentNode.NodeList.Root

	// 	checker := false
	// 	if i%2 == 0 {
	// 		checker = true
	// 	}

	// 	if checker {
	// 		require.Equal(t, 0, currentNode.NodeList.Length)
	// 		currentNode = currentNode.NextNode
	// 		continue
	// 	}

	// 	for j := 0; j < 100; j++ {
	// 		require.Equal(t, currentNodeInNodeList.Value, j)
	// 		currentNodeInNodeList = currentNodeInNodeList.NextNode
	// 	}

	// 	currentNode = currentNode.NextNode
	// }

}
