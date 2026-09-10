package solutions

import "slices"

// @leet start
func productExceptSelf(nums []int) []int {
	prodExceptSelf := make([]int, len(nums))

	counter := 1
	for i := range nums {
		if i == 0 {
			prodExceptSelf[i] = counter
			continue
		}
		prev := nums[i-1]
		counter = counter * prev
		prodExceptSelf[i] = counter
	}

	counter = 1
	for i := range slices.Backward(nums) {
		if i+1 == len(nums) {
			prodExceptSelf[i] *= counter
			continue
		}

		next := nums[i+1]
		counter = counter * next
		prodExceptSelf[i] *= counter
	}

	return prodExceptSelf
}

// @leet end

