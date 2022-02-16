package sorts

import (
	"sort"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				src: []int{5, 4, 3, 2, 1},
			},
		},
		{
			args: args{
				src: []int{2, 4, 3, 5, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectSort(tt.args.src)
			if !sort.IsSorted(sort.IntSlice(tt.args.src)) {
				t.Errorf("%v", tt.args.src)
			}
		})
	}
}
