package sorts

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 16,
			},
			want: 10, // 2の4乗 + 1 (log(n)+1)
		},
		{
			args: args{
				n: 5844674407370955161,
			},
			want: 10, // 2の4乗 + 1 (log(n)+1)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.n); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})

		fmt.Printf("%#v\n", 2*5844674407370955161)
	}
}
