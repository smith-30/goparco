package sorts

import (
	"fmt"
)

var c int

func mergeSort(src []int) []int {
	l := len(src)
	if l < 2 {
		return src
	}

	mid := l / 2

	// fmt.Printf("%#v, %#v\n", src[:mid], src[mid:])

	return merge(
		mergeSort(src[:mid]),
		mergeSort(src[mid:]),
	)
}

func merge(left, right []int) []int {
	// fmt.Printf("merge ;;; %#v, %#v\n", left, right)
	re := make([]int, 0, len(right)+len(left))

	var lIdx, rIdx int
	llen := len(left)
	rlen := len(right)

	for lIdx != llen || rIdx != rlen {
		switch {
		case lIdx == llen:
			re = append(re, right[rIdx])
			rIdx++
		case rIdx == rlen:
			re = append(re, left[lIdx])
			lIdx++
		case left[lIdx] < right[rIdx]:
			re = append(re, left[lIdx])
			lIdx++
		case left[lIdx] > right[rIdx]:
			re = append(re, right[rIdx])
			rIdx++
		}
	}

	return re
}
