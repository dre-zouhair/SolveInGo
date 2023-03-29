package cmd

func IsValidParenthesis(s string) bool {

	if s[0] == '}' || s[0] == ']' || s[0] == ')' {
		return false
	}

	n := len(s) - 1
	if s[n] == '{' || s[n] == '[' || s[n] == '(' {
		return false
	}

	stack := make(map[int]int32)

	index := -1
	for _, c := range s {
		if c == ')' {
			if stack[index] == '(' {
				delete(stack, index)
				index--
				continue
			} else {
				return false
			}
		}

		if c == '}' {
			if stack[index] == '{' {
				delete(stack, index)
				index--
				continue
			} else {
				return false
			}
		}

		if c == ']' {
			if stack[index] == '[' {
				delete(stack, index)
				index--
				continue
			} else {
				return false
			}
		}

		index++
		stack[index] = c

	}

	return len(stack) == 0

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
	return list1
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[1]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}
