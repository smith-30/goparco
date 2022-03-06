package sorts

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Tree struct {
	Value int
	Nodes []*Tree
}

func NewTree(v int) *Tree {
	return &Tree{
		Value: v,
	}
}

func (a *Tree) AddNode(v int) {
	if len(a.Nodes) < 2 {
		a.Nodes = append(a.Nodes, NewTree(v))
		return
	}
	for _, item := range a.Nodes {
		if len(item.Nodes) < 2 {
			item.AddNode(v)
			return
		}
	}
}

func (a *Tree) GetIntSlice() []int {
	re := []int{}
	a.GetNodeValues(re)
	return re
}

func (a *Tree) GetNodeValues(v []int) {
	for _, item := range a.Nodes {
		if 0 < len(item.Nodes) {
			item.GetNodeValues(v)
		}
		v = append(v, item.Value)
	}
}

func (a *Tree) ShowValues() {
	fmt.Printf("%#v\n", a.Value)
	for _, item := range a.Nodes {
		item.ShowValues()
	}
}

func (a *Tree) ShowValuesFromBottom() {
	for _, item := range a.Nodes {
		item.ShowValuesFromBottom()
	}
	fmt.Printf("%#v\n", a.Value)
}

func (a *Tree) HasFullNodes() bool {
	return len(a.Nodes) == 2
}

func treeSimulation() {
	a := []int{3, 1, 2}
	t := NewTree(4)
	for _, item := range a {
		t.AddNode(item)
	}

	bss, _ := json.MarshalIndent(t, "", "	")
	fmt.Printf("%v\n", string(bss))

	t.ShowValuesFromBottom()
}

type Heap []int

func NewHeap(src sort.IntSlice) Heap {
	h := Heap(src)
	return h
}

func heapSort(src []int) {
	if len(src) <= 1 {
		return
	}

	n := len(src)
	for i := n/2 - 1; i >= 0; i-- {
		downMax(src, i, n-1)
	}

	for i := n - 1; i >= 0; i-- {
		swapIntSlice(src, 0, i)
		downMax(src, 0, i-1)
	}
}

func pop(src Heap) {
	l := len(src)
	fmt.Printf("%#v\n", src)
	fmt.Printf("%#v\n", src[0:l-1])
}

func up(src Heap, idx int) {

}

// k: ヒープの条件が崩れている位置
// r: 現在のヒープ数-1
func down(src Heap, k, r int) {
	j := 2*k + 1
	for j <= r {
		var v int
		// right child exists
		if j+1 <= r {
			if src[j] < src[j+1] {
				v = src[j]
			} else {
				v = src[j+1]
				j++
			}
		} else {
			v = src[j]
		}

		if src[k] <= v {
			return
		}

		swapIntSlice(src, k, j)
		k = j
		j = 2*k + 1
	}
}

func downMax(src Heap, k, r int) {
	j := 2*k + 1
	for j <= r {
		var v int
		// right child exists
		if j+1 <= r {
			if src[j] > src[j+1] {
				v = src[j]
			} else {
				v = src[j+1]
				j++
			}
		} else {
			v = src[j]
		}

		if src[k] >= v {
			return
		}

		swapIntSlice(src, k, j)
		k = j
		j = 2*k + 1
	}
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
