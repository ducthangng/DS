package lfu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

// Write test case to input base on information:
// ["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
func TestLFULeedCode(t *testing.T) {
	lfu := Constructor(10)

	input := []struct {
		act   string // put: 1, get: 2
		key   int
		value int
	}{
		{
			act:   "put",
			key:   10,
			value: 13,
		},
		{
			act:   "put",
			key:   3,
			value: 17,
		},
		{
			act:   "put",
			key:   6,
			value: 11,
		},
		{
			act:   "put",
			key:   10,
			value: 5,
		},
		{
			act:   "put",
			key:   9,
			value: 10,
		},
		{
			act:   "get",
			key:   13,
			value: -1,
		},
		{
			act:   "put",
			key:   2,
			value: 19,
		},
		{
			act:   "get",
			key:   2,
			value: 19,
		},
		{
			act:   "get",
			key:   3,
			value: 17,
		},
		{
			act:   "put",
			key:   5,
			value: 25,
		},
		{
			act:   "get",
			key:   8,
			value: -1,
		},
		{
			act:   "put",
			key:   9,
			value: 22,
		},
		{
			act:   "get",
			key:   5,
			value: 25,
		},
		{
			act:   "put",
			key:   1,
			value: 30,
		},
		{
			act:   "get",
			key:   11,
			value: -1,
		},
		{
			act:   "put",
			key:   9,
			value: 12,
		},
		{
			act:   "get",
			key:   7,
			value: -1,
		},
		{
			act:   "get",
			key:   5,
			value: 25,
		},
		{
			act:   "put",
			key:   8,
			value: -1,
		},
		{
			act:   "get",
			key:   9,
			value: 12,
		},
		{
			act:   "put",
			key:   10,
			value: 5,
		},
		{
			act:   "get",
			key:   10,
			value: 5,
		},
		{
			act:   "put",
			key:   6,
			value: 14,
		},
	}

	for i := 0; i < len(input); i++ {
		log.Println("test case: ", i)
		switch input[i].act {
		case "put":
			lfu.Put(input[i].key, input[i].value)
		case "get":
			expected := lfu.Get(input[i].key)
			require.Equal(t, input[i].value, expected)
		}
	}
}
