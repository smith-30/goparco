package main

import "fmt"

import "encoding/json"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (a *Tree) insert(v int) {
	if a != nil {
		if a.LeftNode == nil {
			a.LeftNode = NewTree(v)
		} else {
			if a.RightNode == nil {
				a.RightNode = NewTree(v)
			} else {
				if a.LeftNode != nil {
					a.LeftNode.insert(v)
				} else {
					a.RightNode.insert(v)
				}
			}
		}
	} else {
		a = NewTree(v)
	}
}

//print method for printing a Tree
func print(tree *Tree) {
	if tree != nil {
		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf(" Nil\n")
	}
}

func NewTree(v int) *Tree {
	return &Tree{
		Value: v,
	}
}

func main() {
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	fmt.Printf("%#v\n", "*****")
	tree.insert(3)
	print(tree)
	fmt.Printf("%#v\n", "*****")
	tree.insert(5)
	print(tree)
	fmt.Printf("%#v\n", "*****")
	tree.LeftNode.insert(7)
	print(tree)
	fmt.Printf("%#v\n", "*****")

	bs, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Printf("%v\n", string(bs))
}
