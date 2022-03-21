package sorts

func maxDepth(n int) int {
	var depth int
	// i >>= 1, 右シフト(1/2)
	for i := n; i > 0; i >>= 1 {
		depth++
	}

	return depth * 2
}
