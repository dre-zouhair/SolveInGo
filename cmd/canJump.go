package cmd

func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}

	lastIndex := len(nums) - 1
	gape := 0
	i := 1
	for ; i <= lastIndex; i++ {
		if nums[lastIndex-i] == 0 {
			gape++
			continue
		}
		if nums[lastIndex-i] > gape {
			gape = 0
		} else {
			gape++
		}
	}
	if gape > 0 {
		return false
	}
	return true
}
