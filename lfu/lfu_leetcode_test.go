package lfu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

// Write test case to input base on information:
// ["LFUCache","put","put","get","put","get","get","put","get","get","get"]
// [[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
func TestLFULeedCode(t *testing.T) {
	lfu := Constructor(2)

	input := []struct {
		act   string // put: 1, get: 2
		key   int
		value int
	}{
		{act: "put", key: 1, value: 1},
		{act: "put", key: 2, value: 2},
		{act: "get", key: 1, value: 1},
		{act: "put", key: 3, value: 3},
		{act: "get", key: 2, value: -1},
		{act: "get", key: 3, value: 3},
		{act: "put", key: 4, value: 4},
		{act: "get", key: 1, value: -1},
		{act: "get", key: 3, value: 3},
		{act: "get", key: 4, value: 4},
	}

	for i := 0; i < len(input); i++ {
		log.Println("test case: ", i, " ----- ")
		switch input[i].act {
		case "put":
			lfu.Put(input[i].key, input[i].value)
		case "get":
			expected := lfu.Get(input[i].key)
			require.Equal(t, input[i].value, expected)
		}
	}
}
