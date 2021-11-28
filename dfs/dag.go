package dfs

import (
	"fmt"
)

func dag() {
	x := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	fmt.Println(allPathsSourceTarget(x))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	// myStack store the next go-to node
	// currentPath store the current path
	// result store the result
	result, myStack, currentPath := [][]int{}, [][]int{}, []int{}

	// start from first node
	node, numPath, destination := 0, 0, len(graph)-1

	for {
		if (len(currentPath) > 0) && (node == currentPath[len(currentPath)-1]) {
			if (len(myStack) > 0) && (node == myStack[len(myStack)-1][0]) {
				node = myStack[len(myStack)-1][1]
				myStack = myStack[:len(myStack)-1]
				continue
			}

			currentPath = currentPath[:len(currentPath)-1]

			if len(currentPath) == 0 {
				break
			}

			node = currentPath[len(currentPath)-1]
			continue
		}

		numPath = len(graph[node])
		currentPath = append(currentPath, node)

		if node == destination {
			tempt := make([]int, len(currentPath))
			copy(tempt, currentPath)

			result = append(result, tempt)
			continue
		}

		// finds node next paths
		for j := 0; j < numPath-1; j++ {
			myStack = append(myStack, []int{node, graph[node][j]})
		}

		if numPath > 0 {
			node = graph[node][numPath-1]
			continue
		}

	}

	return result
}
