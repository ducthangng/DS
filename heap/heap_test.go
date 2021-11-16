package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeap(t *testing.T) {

	type input struct {
		cmd string
		val int
	}

	tests := []struct {
		name  string
		input []input
		want  []Node
	}{
		{
			name: "test1",
			input: []input{
				{
					cmd: "push",
					val: 5,
				},
				{
					cmd: "push",
					val: 6,
				},
				{
					cmd: "push",
					val: 10,
				},
				{
					cmd: "push",
					val: 12,
				},
				{
					cmd: "push",
					val: 19,
				},
				{
					cmd: "push",
					val: 9,
				},
				{
					cmd: "push",
					val: 4,
				},
			},
			want: []Node{
				{
					Value: 4,
				},
				{
					Value: 6,
				},
				{
					Value: 5,
				},
				{
					Value: 12,
				},
				{
					Value: 19,
				},
				{
					Value: 10,
				},
				{
					Value: 9,
				},
			},
		},
		{
			name: "test2",
			input: []input{
				{
					cmd: "push",
					val: 5,
				},
				{
					cmd: "push",
					val: 6,
				},
				{
					cmd: "push",
					val: 10,
				},
				{
					cmd: "push",
					val: 12,
				},
				{
					cmd: "push",
					val: 19,
				},
				{
					cmd: "push",
					val: 9,
				},
				{
					cmd: "push",
					val: 4,
				},
				{
					cmd: "pop",
					val: 0,
				},
			},
			want: []Node{
				{
					Value: 5,
				},
				{
					Value: 6,
				},
				{
					Value: 9,
				},
				{
					Value: 12,
				},
				{
					Value: 19,
				},
				{
					Value: 10,
				},
			},
		},
	}

	for _, v := range tests {
		tree := NewHeapTree()
		for _, tt := range v.input {
			switch tt.cmd {
			case "push":
				tree.Push(tt.val, "")
			case "pop":
				tree.Pop()
			}
		}

		r := tree.Retrieve()
		for i, tt := range v.want {
			require.Equal(t, tt.Value, r[i].Value)
		}
	}
}
