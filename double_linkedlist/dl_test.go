package double_linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDL(t *testing.T) {
	list := NewList()

	for i := 0; i < 10; i++ {
		list.Push(i)
	}

	dlist := list.GetAll()

	require.Equal(t, len(dlist), 10)
	for i, v := range dlist {
		if i < len(dlist)-1 {
			require.True(t, v.Value < v.NextNode.Value)
		}
	}

	for _, v := range dlist {
		list.Pop(v)
	}

	dlist = list.GetAll()
	require.Equal(t, len(dlist), 0)
}
