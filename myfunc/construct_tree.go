package myfunc

import (
	"fmt"
)

// Node tree structure type
type Node struct {
	data        int
	left, right *Node
}

func buildUtil(in []int, post []int, inStrt int, inEnd int, pIndex *int) *Node {
	if inStrt > inEnd {
		return nil
	}
	var node *Node = newNode(post[*pIndex])
	(*pIndex)--
	if inStrt == inEnd {
		return node
	}
	var iIndex int = search(in, inStrt, inEnd, node.data)
	node.right = buildUtil(in, post, iIndex+1, inEnd, pIndex)
	node.left = buildUtil(in, post, inStrt, iIndex-1, pIndex)

	return node
}

// CreateTree comment
func CreateTree(in []int, post []int) *Node {
	var pIndex int = len(post) - 1
	return buildUtil(in, post, 0, pIndex, &pIndex)
}

func search(arr []int, strt int, end int, value int) int {
	var i int
	for i = strt; i < end; i++ {
		if arr[i] == value {
			break
		}
	}

	return i
}

func newNode(data int) *Node {
	var node = Node{data: data, left: nil, right: nil}
	return &node
}

// PreOrder comment
func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.data)
	PreOrder(node.left)
	PreOrder(node.right)
}
