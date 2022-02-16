package sorts

import "fmt"

// トランプの並べ替えに近い
// 既存のカードに対して山札から引いたもので並べ替えていくイメージ
// 0番目はソート済みとして扱う
// 3, 9, 6, 1, 2

// 3
// 3, 9 <--- 9 を上の 3 と比較
// 3, 6, 9 <--- 6 を上の 3, 9 と比較
func insertionSorts(src []int) {
	if len(src) <= 1 {
		return
	}
	for i := 1; i < len(src); i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%v %v\n", src[i], src[j])
			if src[i] < src[j] {
				swapIntSlice(src, i, j)
			}
		}
	}
}
