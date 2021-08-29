package main

import (
	"context"
	"fmt"
)

type ctxHandlerFunc = func(ctx *context.Context)

type namedMiddleware struct {
	name string
}

func (a *namedMiddleware) handle(fn ctxHandlerFunc) ctxHandlerFunc {
	return func(ctx *context.Context) {
		fmt.Printf("%#v start\n", a.name)
		fn(ctx)
		fmt.Printf("%#v end\n", a.name)
	}
}

var handler = func(ctx *context.Context) {
	fmt.Printf("%#v\n", "do")
}

func main() {
	ctx := context.TODO()
	s := []string{"a", "b", "c"}
	ms := []namedMiddleware{}

	for i := len(s) - 1; i >= 0; i-- {
		nm := namedMiddleware{
			name: s[i],
		}
		ms = append(ms, nm)
	}

	var chained func(fn ctxHandlerFunc) ctxHandlerFunc
	for _, item := range ms {
		nextFn := item.handle
		if chained == nil {
			chained = nextFn
			continue
		}
		chained(handler)(&ctx)
		fmt.Printf("%#v\n", "---")
		prev := chained
		chained = func(fn ctxHandlerFunc) ctxHandlerFunc {
			// prev は合成されたfn
			// next に prev を渡して新しい関数を作りつなげていく
			//
			// b -> c(fn) -> bc(fn)
			// a -> bc(fn) -> abc(fn)
			return nextFn(prev(fn))
		}
	}
	chained(handler)(&ctx)
}

// output
//
// "a" start
// "b" start
// "c" start
// "do"
// "c" end
// "b" end
// "a" end
