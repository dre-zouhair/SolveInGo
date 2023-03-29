package test

import (
	"Blind75MustDo/cmd"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *cmd.ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case 1",
			args{
				&cmd.ListNode{
					1,
					&cmd.ListNode{
						2,
						&cmd.ListNode{
							3,
							&cmd.ListNode{
								4,
								nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.ReorderList(tt.args.head)
		})
	}
}
