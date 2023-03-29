package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func TestMinWindow1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{
				"ADOBECODEBANC",
				"ABC",
			},
			"BANC",
		}, {
			"case 2",
			args{
				"bba",
				"ab",
			},
			"ba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.MinWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("MinWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
