package tree

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

type Node struct {
	key      interface{}
	children []*Node
	parent   *Node
}

func (n *Node) SetParent(parent *Node) {
	n.parent = parent
	parent.children = append(parent.children, n)
}

func (n *Node) FindRoot() *Node {
	// find root
	var root *Node
	for root = n; root.parent != nil; {
		root = root.parent
	}
	return root
}

func (n *Node) HeightWithoutRecursion() int {

	if n == nil {
		return 0
	}

	height := 0
	q := queue.New()
	root := n.FindRoot()
	q.Enqueue(root)
	q.Enqueue(nil)

	for q.Len() != 0 {
		node := q.Dequeue()

		if node == nil {
			height++
			if q.Len() != 0 {
				q.Enqueue(nil)
			}
		} else {

			for _, child := range node.(*Node).children {
				q.Enqueue(child)
			}
		}
	}

	return height
}

func (n *Node) HeightUsingRecursion() int {

	root := n.FindRoot()
	return heightRecursive(root)
}

func heightRecursive(root *Node) int {
	if root == nil {
		return 0
	}

	maxHeight := 0
	for _, child := range root.children {
		height := heightRecursive(child)
		if height > maxHeight {
			maxHeight = height
		}
	}
	return 1 + maxHeight
}

func CreateNode(key interface{}) *Node {
	node := Node{key: key, children: nil, parent: nil}
	return &node
}

func PrintTree(nodes []*Node) {
	fmt.Println("Tree")
	for _, node := range nodes {
		fmt.Printf("Key: %v\t", node.key)
		if node.parent != nil {
			fmt.Printf("Parent: %v\t", node.parent.key)
		} else {
			fmt.Printf("Parent: ROOT\t")
		}

		fmt.Printf("Children: ")
		for _, child := range node.children {
			fmt.Printf("%v, ", child.key)
		}
		fmt.Println("")
	}
}
