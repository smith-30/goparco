package consistent_hash

import (
	"testing"
)

func Test_doConsistentHashing(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "teet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doConsistentHashing()
		})
	}
}

func Test_searchExample(t *testing.T) {
	type args struct {
		k    int
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				k:    1,
				data: []int{10, 20, 30, 40},
			},
		},
		{
			args: args{
				k:    11,
				data: []int{10, 20, 30, 40},
			},
		},
		{
			args: args{
				k:    21,
				data: []int{10, 20, 30, 40},
			},
		},
		{
			args: args{
				k:    31,
				data: []int{10, 20, 30, 40},
			},
		},
		{
			args: args{
				k:    41,
				data: []int{10, 20, 30, 40},
			},
		},
		{
			args: args{
				k:    1000,
				data: []int{10, 20, 30, 40},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchExample(tt.args.k, tt.args.data)
		})
	}
}
