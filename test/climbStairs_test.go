package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func TestClimbStairs(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			"case 1",
			4,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
