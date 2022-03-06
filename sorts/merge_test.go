package sorts

import (
	"sort"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				src: []int{8, 2},
			},
		},
		{
			args: args{
				src: []int{2, 8, 5, 3, 9, 4, 1},
			},
		},
		{
			args: args{
				src: []int{2, 8, 5, 3, 9, 4, 1, 15, 19, 17, 7, 40, 22, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := mergeSort(tt.args.src)
			if !sort.IsSorted(sort.IntSlice(re)) {
				t.Errorf("%v", re)
			}
		})
	}
}
