package main

import "fmt"

type A interface {
	Do()
}

type a struct {
	v string
}

func (_a *a) Do() {
	fmt.Printf("%#v\n", "do a")
}

type B interface {
	A
	DoB()
}

type b struct {
	a
}

func (_b *b) DoB() {
	fmt.Printf("%#v\n", _b.a.v)
	fmt.Printf("%#v\n", "do b")
}

func main() {
	aa := a{"*aa*"}
	bb := b{aa}
	bb.Do()
	bb.DoB()
}
