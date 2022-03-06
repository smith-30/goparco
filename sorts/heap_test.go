package sorts

import (
	"sort"
	"testing"
)

func Test_treeSimulation(t *testing.T) {
	treeSimulation()
}

func Test_heapSort(t *testing.T) {
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
		{
			args: args{
				src: []int{3, 9, 6, 1, 2},
			},
		},
		{
			args: args{
				src: []int{3, 1, 2},
			},
		},
		{
			args: args{
				src: []int{4, 1, 3, 2},
			},
		},
		{
			args: args{
				src: []int{4, 1, 3, 9, 2},
			},
		},
		{
			args: args{
				src: []int{10, 15, 20, 30, 40},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.src)
			if !sort.IsSorted(sort.IntSlice(tt.args.src)) {
				t.Errorf("%v", tt.args.src)
			}
		})
	}
}
