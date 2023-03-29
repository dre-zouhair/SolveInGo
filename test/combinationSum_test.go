package test

import (
	"Blind75MustDo/cmd"
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "case 1",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "case 2",
			candidates: []int{3, 5, 8},
			target:     11,
			want:       [][]int{{3, 3, 5}, {3, 8}},
		},
		{
			name:       "case 3",
			candidates: []int{7, 3, 2},
			target:     18,
			want:       [][]int{{7, 7, 2, 2}, {7, 3, 3, 3, 2}, {7, 3, 2, 2, 2, 2}, {3, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 2, 2, 2}, {3, 3, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 2, 2, 2}},
		},
		{
			name:       "case 3",
			candidates: []int{},
			target:     18,
			want:       [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmd.CombinationSum(tt.candidates, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
