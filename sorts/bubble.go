package sorts

import "fmt"

func bubble(src []int) {
	maxIdx := len(src) - 1
	for i := 0; i < maxIdx; i++ {
		for j := maxIdx; j > i; j-- {
			_i := src[j]
			_j := src[j-1]
			if _i < _j {
				swapIntSlice(src, j, j-1)
			}
		}
		fmt.Printf("i: %v, %v\n", i, src)
	}
}

func swapIntSlice(src []int, i, j int) {
	src[i], src[j] = src[j], src[i]
}
