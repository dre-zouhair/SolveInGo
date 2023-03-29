package cmd

func maximumCrossingSybArray(nums []int, low, mid, high int) (int, int, int) {

	leftSum, rightSum := -10000000000, -10000000000
	sum, maxRight, maxLeft := 0, mid, mid
	for i := mid; i >= low; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += nums[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

func MaximumSubArray(nums []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, nums[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := MaximumSubArray(nums, low, mid)
		// fmt.Printf("left %v %v => %v %v %v\n", low, mid, leftLow, leftHigh, leftSum)
		rightLow, rightHigh, rightSum := MaximumSubArray(nums, mid+1, high)
		// fmt.Printf("right %v %v => %v %v %v\n", mid+1, high, rightLow, rightHigh, rightSum)
		crossLow, crossHigh, crossSum := maximumCrossingSybArray(nums, low, mid, high)
		// fmt.Printf("cross %v %v %v => %v %v %v\n", low, mid, high, crossLow, crossHigh, crossSum)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func BrutMaximumSubArray(nums []int) (int, int, int) {
	if len(nums) == 1 {
		return 0, 0, nums[0]
	}

	sum, maxSum, low, high := 0, nums[0], 0, 0

	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		temp := maxSum
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum > temp {
				temp = sum
			}
			if temp > maxSum {
				maxSum = temp
				low = i
				high = j
			}
		}
	}

	return low, high, maxSum
}

func LinearMaxSubArray(nums []int) int {
	maxSoFar, maxEndingHere := -10000000, 0

	for i := 0; i < len(nums); i++ {
		maxEndingHere = maxEndingHere + nums[i]
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar

}
