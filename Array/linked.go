package main

import "fmt"

type Node struct {
	next *Node
	val int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val:0}
	tail = root

	for i :=1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)

}

func AddNode(tail *Node, val int) *Node {
	// var tail *Node
	// tail = root'''
	// for tail.next !=nil{
	// 	tail = tail.next
	// }

	node := &Node{val:val}
	tail.next = node
	// tail.next = &Node{val:val}
	return node 
}
