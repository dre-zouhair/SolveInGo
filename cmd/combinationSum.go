package cmd

func CombinationSum(candidates []int, target int) (answers [][]int) {

	if len(candidates) == 0 {
		answers = [][]int{{}}
		return answers
	}

	backtrack(&answers, make([]int, 0), candidates, 0, target)

	return answers
}

func backtrack(answers *[][]int, temp, nums []int, currentElementIndex, target int) {
	if target < 0 || currentElementIndex > len(nums) {
		return
	}

	if target == 0 {
		cpyTmp := make([]int, len(temp))
		copy(cpyTmp, temp)
		*answers = append(*answers, cpyTmp)
	}

	for i := currentElementIndex; i < len(nums); i++ {
		if target < 0 {
			return
		}
		temp = append(temp, nums[i])
		backtrack(answers, temp, nums, i, target-nums[i])
		temp = temp[:len(temp)-1]
	}
}
