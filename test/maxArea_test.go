package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func Test_maxArea(t *testing.T) {

	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Cas de etst",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Cas de etst",
			height: []int{0, 2},
			want:   0,
		}, {
			name:   "Cas de etst",
			height: []int{1, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.MaxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
