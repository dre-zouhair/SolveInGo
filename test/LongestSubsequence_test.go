package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				"bl",
				"kbh",
			},
			1,
		},
		{
			"test case 2",
			args{
				"psnw",
				"vozsh",
			},
			1,
		},
		{
			"test case 3",
			args{
				"abcde",
				"ace",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.LongestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
