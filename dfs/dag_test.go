package dfs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Write test for dag.go
func TestDAG(t *testing.T) {
	// [[2],[],[1]]
	x := [][]int{
		{2},
		{},
		{1},
	}

	result := allPathsSourceTarget(x)

	// [[0,2]]
	require.Equal(t, [][]int{
		{0, 2},
	}, result)
}
