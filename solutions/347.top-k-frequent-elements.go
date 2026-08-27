package solutions

// @leet start

import "slices"

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{}

	for i := range nums {
		c := nums[i]
		cfq := frequencyMap[c] + 1
		frequencyMap[c] = cfq
	}

	frequencies := make([][]int, len(nums)+1)

	for num, freq := range frequencyMap {
		frequencies[freq] = append(frequencies[freq], num)
	}

	result := make([]int, 0, k)

	for _, curr := range slices.Backward(frequencies) {

		for _, each := range curr {
			if len(result) >= k {
				break
			}
			result = append(result, each)
		}

		if len(result) >= k {
			break
		}
	}

	return result

}

// @leet end
