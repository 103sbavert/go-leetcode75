package solutions

import (
	"slices"
)

// @leet start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	triplets := make([][]int, 0)

	for i := range nums {
		iNum := nums[i]

		if iNum > 0 {
			break
		}

		if i > 0 && nums[i-1] == iNum {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := iNum + nums[l] + nums[r]

			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})

				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return triplets
}

// @leet end
