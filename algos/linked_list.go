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

// head を入れ替える前に node をセットして数珠つなぎにしている
func (a *LinkedList) AddToHead(p int) {
	n := Node{
		property: p,
	}

	if a.headNode != nil {
		n.nextNode = a.headNode
	}

	a.headNode = &n
}

func (a *LinkedList) IterateList() {
	var node *Node

	// 多分ここ for のみのscope で　node を headNode にして nextNode 呼ぶことで headNode を次のものにセットしていると思われる
	for node = a.headNode; node != nil; node = node.nextNode {
		// fmt.Printf("%v\n", node.nextNode)
		fmt.Println(node.property)
	}
}

// head を取り出して、処理を行ったあとに、次のnode をセットするようにしてみた。振る舞いは同じ。
func (a *LinkedList) _IterateList() {
	var node *Node
	_a := *a
	for {
		if node = _a.headNode; node != nil {
			fmt.Println(node.property)
			_a.headNode = node.nextNode
		} else {
			return
		}
	}
}

func (a *LinkedList) LastNode() *Node {
	var node *Node
	for {
		if node = a.headNode; node != nil {
			if node.nextNode == nil {
				return node
			}
			a.headNode = node.nextNode
		}
	}
}

func (linkedList *LinkedList) _LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (a *LinkedList) AddToEnd(v int) {
	ln := a._LastNode()
	if ln != nil {
		ln.nextNode = &Node{
			property: v,
		}
	}
}

func (a *LinkedList) AddAfter(nodeProperty int, property int) {
	before := a.NodeWithValue(nodeProperty)
	if before != nil {
		n := &Node{
			property: property,
			nextNode: before.nextNode,
		}
		before.nextNode = n
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

func main() {
	l := LinkedList{}
	l.AddToHead(1)
	l.AddToHead(3)
	l.AddToHead(5)
	l.AddToHead(7)
	l.AddToEnd(9)

	// l.AddAfter(5, 6)

	fmt.Println()
	l.IterateList()
}
