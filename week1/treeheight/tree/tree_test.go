package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"Create new node", args{key: 1}, &Node{key: 1, children: nil, parent: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateNode(tt.args.key)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestNode_SetParent(t *testing.T) {
	type args struct {
		parent *Node
		child  *Node
	}
	tests := []struct {
		name string
		args
	}{
		{"Set parent for child", args{parent: CreateNode(0), child: CreateNode(1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.child.SetParent(tt.args.parent)
			assert.Equal(t, tt.args.parent, tt.args.child.parent)
			assert.Contains(t, tt.args.parent.children, tt.args.child)
		})
	}
}

func TestNode_HeightWithoutRecursion(t *testing.T) {
	type args struct {
		numNodes        int
		childParentLink []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 4 -1 4 1 1", args{numNodes: 5, childParentLink: []int{4, -1, 4, 1, 1}}, 3},
		{"5 -1 0 4 0 3", args{numNodes: 5, childParentLink: []int{-1, 0, 4, 0, 3}}, 4},
		{"10 9 7 5 5 2 9 9 9 2 -1", args{numNodes: 10, childParentLink: []int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := setupNodes(tt.args.numNodes, tt.args.childParentLink)
			height := nodes[0].HeightWithoutRecursion()
			assert.Equal(t, tt.want, height)
		})
	}
}

func TestNode_HeightUsingRecursion(t *testing.T) {
	type args struct {
		numNodes        int
		childParentLink []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 4 -1 4 1 1", args{numNodes: 5, childParentLink: []int{4, -1, 4, 1, 1}}, 3},
		{"5 -1 0 4 0 3", args{numNodes: 5, childParentLink: []int{-1, 0, 4, 0, 3}}, 4},
		{"10 9 7 5 5 2 9 9 9 2 -1", args{numNodes: 10, childParentLink: []int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := setupNodes(tt.args.numNodes, tt.args.childParentLink)
			height := nodes[0].HeightUsingRecursion()
			assert.Equal(t, tt.want, height)
		})
	}
}

func setupNodes(numNodes int, childParentLink []int) []*Node {

	var nodes []*Node
	for i := 0; i < numNodes; i++ {
		node := CreateNode(i)
		nodes = append(nodes, node)
	}
	for child, parent := range childParentLink {
		// not a root node
		if parent >= 0 {
			nodes[child].SetParent(nodes[parent])
		}
	}
	return nodes
}
