package main

import "fmt"

// 双方向の関連を持つ
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (a *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var hit, node *Node
	for node = a.headNode; node != nil; node = node.nextNode {
		if node.nextNode != nil && node.previousNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				hit = node
			}
		}
	}
	return hit
}

func (a *LinkedList) AddToHead(p int) {
	n := &Node{
		property: p,
	}

	if a.headNode != nil {
		a.headNode.previousNode = n
		n.nextNode = a.headNode
	}

	a.headNode = n
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

//AddToEnd method of LinkedList
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	var node *Node
	node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
