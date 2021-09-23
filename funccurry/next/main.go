package main

import (
	"context"
	"fmt"
)

type ctxHandlerFunc = func(ctx *context.Context)

var middleware = func(name string) ctxHandlerFunc {
	return func(ctx *context.Context) {
		fmt.Printf("%#v start\n", name)
		fn(ctx)
		fmt.Printf("%#v end\n", name)
	}
}

var handler = func(ctx *context.Context) {
	fmt.Printf("%#v\n", "do")
}

func main() {
	ms := []ctxHandlerFunc{}
	for _, item := range []string{} {

	}
}
