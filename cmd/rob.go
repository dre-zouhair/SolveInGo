package cmd

func bt(dp *map[int]int, nums []int, index int, max int) int {
	if _, ok := (*dp)[index]; ok {
		return (*dp)[index]
	}
	if index+2 > len(nums) {
		return max
	}

	temp := 0
	for i := index + 2; i < len(nums); i++ {
		local := bt(dp, nums, i, nums[i])
		if local > temp {
			temp = local
		}
	}

	(*dp)[index] = max + temp

	return (*dp)[index]
}

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	rob1 := helper(nums)

	return rob1
}

func helper(nums []int) int {
	rob1, rob2 := 0, 0

	for _, n := range nums {
		rob1, rob2 = rob2, max(n+rob1, rob2)
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
