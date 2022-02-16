package sorts

// 区間の一番小さい値をみつけて
// 先頭に追加していく
func selectSort(src []int) {
	nextIdx := 0
	srcLen := len(src)
	for nextIdx != srcLen {
		now := src[nextIdx]
		swapIdx := nextIdx
		for i := nextIdx; i < srcLen-1; i++ {
			if src[i+1] < now {
				now = src[i+1]
				swapIdx = i + 1
			}
		}
		swapIntSlice(src, nextIdx, swapIdx)
		nextIdx++
	}
}
