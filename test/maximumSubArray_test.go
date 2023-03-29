package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func TestMaximumSubArray(t *testing.T) {
	type args struct {
		nums []int
		low  int
		high int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
				low:  0,
				high: 12,
			},
			want:  7,
			want1: 10,
			want2: 43,
		},
		{
			name: "case 1",
			args: args{
				nums: []int{-3, -3, -25, -3, -16, -23, -7, -5, -22, -4},
				low:  0,
				high: 9,
			},
			want:  0,
			want1: 0,
			want2: -3,
		},
	}
	for _, tt := range tests {
		/*		t.Run(tt.name, func(t *testing.T) {
				got, got1, got2 := MaximumSubArray(tt.args.nums, tt.args.low, tt.args.high)
				if got != tt.want {
					t.Errorf("MaximumSubArray() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("MaximumSubArray() got1 = %v, want %v", got1, tt.want1)
				}
				if got2 != tt.want2 {
					t.Errorf("MaximumSubArray() got2 = %v, want %v", got2, tt.want2)
				}
			})*/
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := cmd.BrutMaximumSubArray(tt.args.nums)
			if got != tt.want {
				t.Errorf("MaximumSubArray() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaximumSubArray() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MaximumSubArray() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
