package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", 111)

	do([]int{0}, a, b)
}

type fn func([]int)

func a(v []int) {
	v[0] = 1
	fmt.Printf("%#v\n", v)
}

func b(v []int) {
	v[0] = 2
	fmt.Printf("%#v\n", v)
}

func do(v []int, fns ...fn) {
	for _, item := range fns {
		item(v)
	}
	fmt.Printf("%#v\n", v)
}
