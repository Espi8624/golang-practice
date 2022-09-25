package main

import "fmt"

// generic 도입전
// type Node struct {
// 	val interface{}
// 	next *Node
// }

// 도입후
type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}

func main() {
	// node1 := &Node[int] {val: 1}
	node1 := NewNode(1) // *Node[int]
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Print(node1.val, " - ")
		node1 = node1.next
	}

	// node2 := &Node[string]{val: "Hi"}
	node2 := NewNode("Hi") // *Node[string]
	node2.Push("Hello").Push("How are you")

	for node2 != nil {
		fmt.Print(node2.val, " - ")
		node2 = node2.next
	}
}
