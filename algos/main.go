package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(22)
	intList.PushBack(33)

	for elm := intList.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value.(int))
	}
}
