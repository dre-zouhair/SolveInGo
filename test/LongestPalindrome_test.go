package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "",
			args: args{
				s: "abbc",
			},
			want: "bb",
		},
		{
			name: "",
			args: args{
				s: "ccc",
			},
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
