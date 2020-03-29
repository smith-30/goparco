package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (a *LinkedList) AddToHead(p int) {
	n := Node{
		property: p,
	}

	if a.headNode != nil {
		n.nextNode = a.headNode
	}
	a.headNode = &n
}

func main() {
	l := LinkedList{}
	l.AddToHead(1)
	l.AddToHead(3)
	fmt.Printf("%#v\n", l.headNode.property)

	fmt.Printf("%#v\n", l.headNode)
	fmt.Printf("%#v\n", l.headNode.nextNode)
}
